package schema_cleaner

import "gitlab.com/zullpro/core/1cclientgenerator.git/shared"

func ExtractAssociations(source []shared.Association) map[string]map[string]string {
	result := make(map[string]map[string]string)
	for _, assoc := range source {
		name := "StandardODATA." + assoc.Name
		if _, ok := result[name]; !ok {
			result[name] = make(map[string]string, len(assoc.Ends))
		}
		for _, end := range assoc.Ends {
			result[name][end.Role] = end.Type
		}
	}
	return result
}

func ClearSchema(source *shared.Schema, typeMap map[string]string) *shared.Schema {
	result := *source
	associations := ExtractAssociations(source.Association)
	types := make(map[string]bool)
	for t := range typeMap {
		types[t] = true
	}

	prevCnt := -1
	for prevCnt < len(types) {
		prevCnt = len(types)
		for _, e := range source.Entities {
			if _, ok := types[TranslateType(e.Name)]; ok {
				for _, p := range e.Properties {
					types[TranslateType(p.Type)] = true
				}

				for _, n := range e.Navigations {
					types[TranslateType(associations[n.Type][n.FromRole])] = true
					types[TranslateType(associations[n.Type][n.ToRole])] = true
				}
			}
		}

		for _, e := range source.Complexes {
			if _, ok := types[TranslateType(e.Name)]; ok {
				for _, p := range e.Properties {
					types[TranslateType(p.Type)] = true
				}

				for _, n := range e.Navigations {
					types[TranslateType(associations[n.Type][n.FromRole])] = true
					types[TranslateType(associations[n.Type][n.ToRole])] = true
				}
			}
		}
	}

	result.Entities = make([]shared.OneCType, 0)
	result.Complexes = make([]shared.OneCType, 0)

	for _, e := range source.Entities {
		if _, ok := types[e.Name]; ok {
			result.Entities = append(result.Entities, e)
		}
	}

	for _, e := range source.Complexes {
		if _, ok := types[e.Name]; ok {
			result.Complexes = append(result.Complexes, e)
		}
	}
	return &result
}
