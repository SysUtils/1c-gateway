package subscriptions

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
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

func (g *Generator) genWatchers(createEntities, updateEntities []shared.OneCType) string {

	queries := `func (r *Watcher) watch(offset time.Duration) {
	watchInterval := time.Second * 10
	ticker := time.NewTicker(watchInterval)
	dataVersions := map[string]map[Guid]String {}
	itemkeys  := map[string][]Guid {}
	catalogKeys  := map[string]map[Guid]struct{} {}
`
	for _, val := range createEntities {
		queries += fmt.Sprintf(`	itemkeys["%s"] = []Guid {}
`, g.translateType(val.Name))
	}
	for _, val := range updateEntities {
		queries += fmt.Sprintf(`	dataVersions["%s"] = map[Guid]String {}
`, g.translateType(val.Name))
		queries += fmt.Sprintf(`	catalogKeys["%s"] = map[Guid]struct{} {}
`, g.translateType(val.Name))
		if !FindString(val.Name, createEntities) {
			queries += fmt.Sprintf(`	itemkeys["%s"] = []Guid {}
`, g.translateType(val.Name))
		}

	}
	queries += `
	for range ticker.C {
		r.listenersCountsMutex.RLock()
			for sType, count := range r.listenersCounts {
				switch sType {
`
	for _, val := range createEntities {
		queries += g.genCreateWatcher(val)
	}
	for _, val := range updateEntities {
		queries += g.genUpdateWatcher(val)
		queries += g.genDeleteWatcher(val)
		if !FindString(val.Name, createEntities) {
			queries += g.genCreateHeavyWatcher(val)
		}
	}

	queries += fmt.Sprintf(`
				}
			}
		r.listenersCountsMutex.RUnlock()
	}
}
`)
	return queries
}

func (g *Generator) genCreateWatcher(source shared.OneCType) string {
	t := g.translateType(source.Name)
	f := g.translateName("Period")
	result := fmt.Sprintf(`
				case Create%s:
					if count > 0 {
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
								r.Created <- &p
							}
						}
					} else if count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, count)
					}
`, t, t, f, t)
	return result
}

func (g *Generator) genCreateHeavyWatcher(source shared.OneCType) string {
	t := g.translateType(source.Name)
	k := g.translateName("Ref_Key")
	result := fmt.Sprintf(`
				case Create%s:
					flag := len(catalogKeys["%s"]) == 0
					if count > 0 {
						items, err := r.Client.%ss(Where{Fields:[]string{"Ref_Key"}})
						if err != nil {
							log.Println(err)
						} else  {
							for  _, skeleton := range *items {
								if _, ok := catalogKeys["%s"][*skeleton.%s]; !ok {
									catalogKeys["%s"][*skeleton.%s] = struct{} {}
									if !flag {
										item, err := r.Client.%s(Primary%s{%s:*skeleton.%s}, nil)
										if err != nil {
											log.Println(err)
										}
										r.Created <- item
									}
								}
							}
						}
					} else if count == 0 {
						catalogKeys["%s"] = map[Guid]struct{} {}
					} else if count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, count)
					}
`, t, t, t, t, k, t, k, t, t, k, k, t)
	return result
}

func (g *Generator) genUpdateWatcher(source shared.OneCType) string {
	t := g.translateType(source.Name)
	k := g.translateName("Ref_Key")
	dv := g.translateName("DataVersion")
	result := fmt.Sprintf(`
				case Update%s:
					if count > 0 {
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
										r.Updated <- item
									}
								}
								dataVersions["%s"][*p.%s] = *p.%s
							}
						}
					} else if count == 0 {
						dataVersions["%s"] = map[Guid]String {}
					} else if count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, count)
					}
`, t, t, t, k, dv, t, t, k, k, t, k, dv, t)
	return result
}

func (g *Generator) genDeleteWatcher(source shared.OneCType) string {
	t := g.translateType(source.Name)
	k := g.translateName("Ref_Key")
	result := fmt.Sprintf(`
				case Delete%s:
					if count > 0 {
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
									r.Deleted <- &%s { %s: &v }
								}
							}
							itemkeys["%s"] = newItems
						}
					} else if count == 0 {
						itemkeys["%s"] = []Guid {}
					} else if count < 0 {
						log.Panicf("count of sypscribers type %%s is %%d", sType, count)
					}
`, t, t, k, k, t, t, k, t, t)
	return result
}
