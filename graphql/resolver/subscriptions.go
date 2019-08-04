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
	result += g.genWatchers(updateEntities, createEntities)
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
	r.subscribers <- &Subscriber{uType: "Create%s", events: c, stop: ctx.Done()}

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
	r.subscribers <- &Subscriber{uType: "Update%s", events: c, stop: ctx.Done()}

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
	r.subscribers <- &Subscriber{uType: "Delete%s", events: c, stop: ctx.Done()}

	return c
}
`, t, t, t, t)
}

func (g *Generator) genWatchers(updateEntities, createEntities []shared.OneCType) string {

	queries := `func (r *GqlResolver) watcher(offset time.Duration, subscribersCounts *map[string]*int32) {
	watchInterval := time.Second * 10
	ticker := time.NewTicker(watchInterval)
	dataVersions := map[string]map[Guid]String {}
	itemkeys  := map[string][]Guid {}
`
	for _, val := range createEntities {
		queries += fmt.Sprintf(`	itemkeys["%s"] = []Guid {}
`, g.translateType(val.Name))
	}
	for _, val := range updateEntities {
		queries += fmt.Sprintf(`	dataVersions["%s"] = map[Guid]String{}
`, g.translateType(val.Name))
		queries += fmt.Sprintf(`	itemkeys["%s"] = []Guid {}
`, g.translateType(val.Name))
	}
	queries += `
	for range ticker.C {
		if subscribersCounts != nil {
			for sType, count := range *subscribersCounts {
				switch sType {
`
	for _, val := range createEntities {
		queries += g.genCreateWatcher(val)
	}
	for _, val := range updateEntities {
		queries += g.genUpdateWatcher(val)
		queries += g.genDeleteWatcher(val)
	}

	queries += fmt.Sprintf(`
				}
			}
		}
	}
}
`)
	return queries
}

func (g *Generator) genCreateWatcher(source shared.OneCType) string {
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
								r.createEvents <- &p
							}
						}
					} else if count != nil && *count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, *count)
					}
`, t, t, f, t)
	return result
}

func (g *Generator) genUpdateWatcher(source shared.OneCType) string {
	t := g.translateType(source.Name)
	k := g.translateName("Ref_Key")
	dv := g.translateName("DataVersion")
	result := fmt.Sprintf(`
				case "Update%s":
					if count != nil && *count > 0 {
						items, err := r.Client.%ss(Where{Fields:[]string{"Ref_Key", "DataVersion"}})
						if err != nil {
							log.Println(err)
						} else  {
							for _, p := range *items {
								if val, ok := dataVersions["%s"][*p.%s]; ok {
									if val != *p.%s {
										item, err := r.Client.%s(Primary%s{%s:*p.%s}, nil)
										if err != nil {
											log.Println(err)
										}
										r.updateEvents <- item
									}
								}
								dataVersions["%s"][*p.%s] = *p.%s
							}
						}
					} else if count != nil && *count == 0 {
						dataVersions["%s"] = map[Guid]String {}
					} else if count != nil && *count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, *count)
					}
`, t, t, t, k, dv, t, t, k, k, t, k, dv, t)
	return result
}

func (g *Generator) genDeleteWatcher(source shared.OneCType) string {
	t := g.translateType(source.Name)
	k := g.translateName("Ref_Key")
	result := fmt.Sprintf(`
				case "Delete%s":
					if count != nil && *count > 0 {
						items, err := r.Client.%ss(Where{Fields:[]string{"Ref_Key"}})
						if err != nil {
							log.Println(err)
						} else  {
							newItems := make([]Guid, len(*items))
							itemMap := map[Guid]bool {}
							for i, p := range *items {
								newItems[i] = *p.%s
								itemMap[*p.%s] = true
							}
							for _, v := range itemkeys["%s"] {
								if _, ok := itemMap[v]; !ok {
									r.deleteEvents <- &%s { %s: &v }
								}
							}
							itemkeys["%s"] = newItems
						}
					} else if count != nil && *count == 0 {
						itemkeys["%s"] = []Guid {}
					} else if count != nil && *count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, *count)
					}
`, t, t, k, k, t, t, k, t, t)
	return result
}

func (g *Generator) getBroadcasters(updateEntities, createEntities []shared.OneCType) string {
	counters := ""
	createEvents := ""
	updateEvents := ""
	deleteEvents := ""

	for _, entity := range createEntities {
		t := g.translateType(entity.Name)
		counters += fmt.Sprintf(`"Create%s": new(int32),`, t)
		counters += "\n"

		createEvents += fmt.Sprintf(`case *%s:
				for id, s := range subscribers {
					if strings.HasPrefix(s.uType, "Create") {
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
		counters += fmt.Sprintf(`"Update%s": new(int32),`, t)
		counters += "\n"

		updateEvents += fmt.Sprintf(`case *%s:
				for id, s := range subscribers {
					if strings.HasPrefix(s.uType, "Update") {
						switch eventChan := s.events.(type) {
						case chan *%s:
							go func(id string, s *Subscriber) {
								eventChan <- ev
							}(id, s)
						}
					}
				}
`, t, t)
		counters += fmt.Sprintf(`"Delete%s": new(int32),`, t)
		counters += "\n"
		deleteEvents += fmt.Sprintf(`case *%s:
				for id, s := range subscribers {
					if strings.HasPrefix(s.uType, "Delete") {
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
		case e := <-r.createEvents:
			switch ev := e.(type) {
			` + createEvents + `
			}
		case e := <-r.updateEvents:
			switch ev := e.(type) {
			` + updateEvents + `
			}
		case e := <-r.deleteEvents:
			switch ev := e.(type) {
			` + deleteEvents + `
			}
		}
	}
}
`)
	return result
}
