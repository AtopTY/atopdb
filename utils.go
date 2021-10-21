package atopdb

import (
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// CheckKeysExist check require key present
func CheckKeysExist(input bson.M, keys []string) error {
	for _, v := range keys {
		if _, ok := input[v]; !ok {
			return fmt.Errorf("lost field : %s", v)
		}
	}
	return nil
}

// getIntOr get value form query create by parseQuery(),
func getIntOr(q bson.M, key string, def int64) int64 {
	v, ok := q[key]
	if !ok {
		return def
	}
	str, ok := v.(string)
	if !ok {
		return def
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return i
}
