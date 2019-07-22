// Package schema_cleaner provides functions to clean the schema
package schema_cleaner

import "github.com/SysUtils/1c-gateway/shared"

func extractAssociations(source []shared.Association) map[string]map[string]string {
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

// Removes all types that are not in the TypeMap and returns the clean schema
func Clean(source *shared.Schema, typeMap map[string]string) *shared.Schema {
	result := *source
	associations := extractAssociations(source.Association)
	types := make(map[string]bool)
	for t := range typeMap {
		types[t] = true
	}

	prevCnt := -1
	for prevCnt < len(types) {
		prevCnt = len(types)
		for _, e := range source.Entities {
			if _, ok := types[translateType(e.Name)]; ok {
				for _, p := range e.Properties {
					types[translateType(p.Type)] = true
				}

				for _, n := range e.Navigations {
					types[translateType(associations[n.Type][n.FromRole])] = true
					types[translateType(associations[n.Type][n.ToRole])] = true
				}
			}
		}

		for _, e := range source.Complexes {
			if _, ok := types[translateType(e.Name)]; ok {
				for _, p := range e.Properties {
					types[translateType(p.Type)] = true
				}

				for _, n := range e.Navigations {
					types[translateType(associations[n.Type][n.FromRole])] = true
					types[translateType(associations[n.Type][n.ToRole])] = true
				}
			}
		}
	}

	result.Entities = make([]shared.OneCType, 0)
	result.Complexes = make([]shared.OneCType, 0)
	result.Functions = make([]shared.Function, 0)

	for _, e := range source.Entities {
		if _, ok := types[e.Name]; ok {
			result.Entities = append(result.Entities, e)
		}
	}

	for _, e := range source.Functions {
		if !e.IsBindable {
			result.Functions = append(result.Functions, e)
		} else if _, ok := types[translateType(e.Parameters[0].Type)]; ok {
			result.Functions = append(result.Functions, e)
		}
	}

	for _, e := range source.Complexes {
		if _, ok := types[e.Name]; ok {
			result.Complexes = append(result.Complexes, e)
		}
	}
	return &result
}
