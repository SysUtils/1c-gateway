package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
)

var updateSubscriptionValidPrefixes = []string{
	"Catalog_",
	"Document_",
}

var createSubscriptionValidPrefixes = []string{
	"AccumulationRegister_",
	"InformationRegister_",
}
var updateSubscriptionValidSuffixes = []string{
	"_RecordType",
	"Turnover",
}

func (g *Generator) GenSubAll(source []shared.OneCType) string {
	result := ""
	updateEntities := []shared.OneCType{}
	createEntities := []shared.OneCType{}
	for _, entity := range source {
		validPrefix := false
		validSuffix := false
		for _, s := range createSubscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}

		for _, s := range updateSubscriptionValidSuffixes {
			if strings.HasSuffix(entity.Name, s) {
				validSuffix = true
				break
			}
		}
		if validPrefix && validSuffix {
			createEntities = append(createEntities, entity)
		}
		validPrefix = false
		validSuffix = false
		for _, s := range updateSubscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}
		validSuffix = strings.Count(entity.Name, "_") == 1

		if validPrefix && validSuffix {
			updateEntities = append(updateEntities, entity)
		}
	}
	result += g.genSubscriptions(createEntities, updateEntities)
	result += g.getBroadcasters(updateEntities, createEntities)
	return result
}

func (g *Generator) genSubscriptions(createEntities, upateEntities []shared.OneCType) string {
	queries := ""
	for _, entity := range createEntities {
		queries += g.genCreateCallback(entity)
	}
	for _, entity := range upateEntities {
		queries += g.genUpdateCallback(entity)
		queries += g.genDeleteCallback(entity)
	}
	return queries
}

func (g *Generator) genCreateCallback(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`func (r *GqlResolver) OnCreate%s(ctx context.Context) <-chan *%s {
	c := make(chan *%s)
	// NOTE: this could take a while
	r.subscribers <- &Subscriber{uClass: Create, uType: Create%s, events: c, stop: ctx.Done()}

	return c
}
`, t, t, t, t)
}

func (g *Generator) genUpdateCallback(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`func (r *GqlResolver) OnUpdate%s(ctx context.Context) <-chan *%s {
	c := make(chan *%s)
	// NOTE: this could take a while
	r.subscribers <- &Subscriber{uClass: Update, uType: Update%s, events: c, stop: ctx.Done()}

	return c
}
`, t, t, t, t)
}

func (g *Generator) genDeleteCallback(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`func (r *GqlResolver) OnDelete%s(ctx context.Context) <-chan *%s {
	c := make(chan *%s)
	// NOTE: this could take a while
	r.subscribers <- &Subscriber{uClass: Delete, uType: Delete%s, events: c, stop: ctx.Done()}

	return c
}
`, t, t, t, t)
}

func (g *Generator) getBroadcasters(updateEntities, createEntities []shared.OneCType) string {
	counters := ""
	createEvents := ""
	updateEvents := ""
	deleteEvents := ""

	for _, entity := range createEntities {
		t := g.translateType(entity.Name)
		counters += fmt.Sprintf(`Create%s: new(int32),`, t)
		counters += "\n"

		createEvents += fmt.Sprintf(`case *%s:
				for id, s := range subscribers {
					if s.uClass == Create {
						switch eventChan := s.events.(type) {
						case chan *%s:
							go func(id string, s *Subscriber) {
								eventChan <- ev
							}(id, s)
						}
					}
				}
`, t, t)
	}

	for _, entity := range updateEntities {
		t := g.translateType(entity.Name)
		counters += fmt.Sprintf(`Update%s: new(int32),`, t)
		counters += "\n"

		updateEvents += fmt.Sprintf(`case *%s:
				for id, s := range subscribers {
					if s.uClass == Update {
						switch eventChan := s.events.(type) {
						case chan *%s:
							go func(id string, s *Subscriber) {
								eventChan <- ev
							}(id, s)
						}
					}
				}
`, t, t)
		counters += fmt.Sprintf(`Delete%s: new(int32),`, t)
		counters += "\n"
		deleteEvents += fmt.Sprintf(`case *%s:
				for id, s := range subscribers {
					if s.uClass == Delete {
						switch eventChan := s.events.(type) {
						case chan *%s:
							go func(id string, s *Subscriber) {
								eventChan <- ev
							}(id, s)
						}
					}
				}
`, t, t)
	}

	result := fmt.Sprintf(`func (r *GqlResolver) broadcast(offset time.Duration) {
	subscribers := map[string]*Subscriber{}

	subscribersCounts := map[EventType]*int32{
		` + counters + `
	}
	unsubscribe := make(chan unsubscribeEvent)

	for {
		select {
		case e := <-unsubscribe:
			atomic.AddInt32(subscribersCounts[e.uType], -1)
			r.watcher.AddEventListener(e.uType, -1)
			delete(subscribers, e.id)
		case s := <-r.subscribers:
			atomic.AddInt32(subscribersCounts[s.uType], 1)
			r.watcher.AddEventListener(s.uType, 1)
			id := uuid.New().String()
			subscribers[id] = s
			go func() {
				<-s.stop
				unsubscribe <- unsubscribeEvent{uType: s.uType, id: id}
			}()
		case e := <-r.watcher.Created:
			switch ev := e.(type) {
			` + createEvents + `
			}
		case e := <-r.watcher.Updated:
			switch ev := e.(type) {
			` + updateEvents + `
			}
		case e := <-r.watcher.Deleted:
			switch ev := e.(type) {
			` + deleteEvents + `
			}
		}
	}
}
`)
	return result
}
