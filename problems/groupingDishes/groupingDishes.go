package groupingDishes

import (
	"sort"
)

func groupingDishes(dishes [][]string) [][]string {

	dishMap := make(map[string][]string)
	output := [][]string{}

	for _, dish := range dishes {
		dishName := dish[0]
		ingredients := dish[1:]
		for _, ingredient := range ingredients {
			if _, ok := dishMap[ingredient]; !ok {
				dishMap[ingredient] = []string{}
			}
			dishMap[ingredient] = append(dishMap[ingredient], dishName)
		}
	}

	keys := []string{}
	for k, v := range dishMap {
		keys = append(keys, k)
		sort.Strings(v)
	}
	sort.Strings(keys)

	for _, ingredient := range keys {
		row := []string{ingredient}
		row = append(row, dishMap[ingredient]...)
		if len(row) > 2 {
			output = append(output, row)
		}
	}

	return output
}
