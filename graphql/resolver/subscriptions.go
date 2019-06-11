package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
)

var subscriptionValidPrefixes = []string{
	"AccumulationRegister_",
	"InformationRegister_",
}
var subscriptionValidSuffixes = []string{
	"_RecordType",
	"Turnover",
}

func (g *Generator) genSubscriptions(source []shared.OneCType) string {
	queries := ""
	for _, entity := range source {
		validPrefix := false
		validSuffix := false
		for _, s := range subscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}

		for _, s := range subscriptionValidSuffixes {
			if strings.HasSuffix(entity.Name, s) {
				validSuffix = true
				break
			}
		}
		if validPrefix && validSuffix {
			queries += g.genSubscriptionCallback(entity)
			queries += "\n"
		}
	}
	return queries
}

func (g *Generator) genSubscriptionTypes(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`func (r *GqlResolver) OnCreate%s(ctx context.Context) <-chan *%s {
	c := make(chan *%s)
	// NOTE: this could take a while
	r.create <- &create%sSubscriber{events: c, stop: ctx.Done()}

	return c
}
`, t, t, t, t)
}

func (g *Generator) genSubscriptionCallback(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`func (r *GqlResolver) OnCreate%s(ctx context.Context) <-chan *%s {
	c := make(chan *%s)
	// NOTE: this could take a while
	r.create <- &create%sSubscriber{uType: "Create%s", events: c, stop: ctx.Done()}

	return c
}
`, t, t, t, t, t)
}

func (g *Generator) genWatchers(source []shared.OneCType) string {
	queries := `func (r *GqlResolver) watcher(offset time.Duration, subscribersCounts *map[string]*int32) {
	watchInterval := time.Second * 10
	ticker := time.NewTicker(watchInterval)

	for range ticker.C {
		if subscribersCounts != nil {
			for sType, count := range *subscribersCounts {
				switch sType {
`
	for _, entity := range source {
		validPrefix := false
		validSuffix := false
		for _, s := range subscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}

		for _, s := range subscriptionValidSuffixes {
			if strings.HasSuffix(entity.Name, s) {
				validSuffix = true
				break
			}
		}
		if validPrefix && validSuffix {
			queries += g.genWatcher(entity)
			queries += "\n"
		}
	}
	queries += fmt.Sprintf(`
				}
			}
		}
	}
}`)
	return queries
}

func (g *Generator) genWatcher(source shared.OneCType) string {
	t := g.translateType(source.Name)
	f := g.translateName("Period")
	result := fmt.Sprintf(`
				case "Create%s":
					if count != nil && *count > 0 {
						now := time.Now().Add(-watchInterval - time.Hour*offset)
						datetime := (DateTime)(now)
						filter := %sFilter{
							%sGt: &datetime,
						}

						items, err := r.Client.%ss(Where{Filter: filter})
						if items != nil {
							log.Printf("registries requested from %%v to %%v, count of rows is %%d", now, time.Now().Add(-time.Hour*offset), len(*items))
						} else {
							log.Printf("registries requested from %%v to %%v, count of rows is %%d", now, time.Now().Add(-time.Hour*offset), 0)
						}
						if err != nil {
							log.Println(err)
						}
						if items != nil {
							for _, p := range *items {
								r.events <- &p
							}
						}
					} else if count != nil && *count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, *count)
					}`, t, t, f, t)
	return result
}

func (g *Generator) getBroadcasters(source []shared.OneCType) string {
	counters := ""
	newSubs := ""
	events := ""

	for _, entity := range source {
		validPrefix := false
		validSuffix := false
		for _, s := range subscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}

		for _, s := range subscriptionValidSuffixes {
			if strings.HasSuffix(entity.Name, s) {
				validSuffix = true
				break
			}
		}
		if validPrefix && validSuffix {
			t := g.translateType(entity.Name)
			counters += fmt.Sprintf(`"Create%s": &(zero),`, t)
			counters += "\n"

			newSubs += fmt.Sprintf(`
`, t, t)
			events += fmt.Sprintf(`case *%s:
				for id, s := range subscribers {
					switch eventChan := s.events.(type) {
					case chan *%s:
						go func(id string, s *Subscriber) {
							eventChan <- ev
						}(id, s)
					default:
						log.Panicf("Wrong type %%s in circle of createProduction subscribers", eventChan)
					}
				}
`, t, t)
		}
	}

	result := fmt.Sprintf(`func (r *GqlResolver) broadcast(offset time.Duration) {
	subscribers := map[string]*Subscriber{}

	zero := int32(0)

	subscribersCounts := map[string]*int32{
		` + counters + `
	}
	go r.watcher(offset, &subscribersCounts)
	unsubscribe := make(chan unsubscribeEvent)

	for {
		select {
		case e := <-unsubscribe:
			atomic.AddInt32(subscribersCounts[e.uType], -1)
			delete(subscribers, e.id)
		case s := <-r.subscribers:
			atomic.AddInt32(subscribersCounts[s.uType], 1)
			id := uuid.New().String()
			subscribers[id] = s
			go func() {
				<-s.stop
				unsubscribe <- unsubscribeEvent{uType: s.uType, id: id}
			}()
		case e := <-r.events:
			switch ev := e.(type) {
			` + events + `
			}
		}
	}
}`)
	return result
}

func (g *Generator) getBroadcasterUnsub(source shared.OneCType) string {
	result := `
				case "%s":
					atomic.AddInt32(subscribersCounts[uType], -1)
					delete(%s, e.id)`
	return result
}
