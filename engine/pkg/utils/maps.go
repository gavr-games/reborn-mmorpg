package utils

import (
	"math/rand"
)

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
			vm, ok := v.(map[string]interface{})
			if ok {
					cp[k] = CopyMap(vm)
			} else {
					cp[k] = v
			}
	}

	return cp
}

func PickRandomInMap(m map[string]interface{}) interface{} {
	k := rand.Intn(len(m))
	for _, x := range m {
		if k == 0 {
			return x
		}
		k--
	}
	panic("unreachable")
}
