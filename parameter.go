package atopdb

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// PageParameter
type PageParameter struct {
	Limit  int64
	Offset int64
	Header int64
	Last   int64
}

var defaultPageParameter = PageParameter{
	Limit:  20,
	Offset: 0,
	Header: 0,
	Last:   0,
}

// GetPageParameter gate page parameter form http request
func GetPageParameter(query bson.M) PageParameter {

	var parameter PageParameter
	parameter.Limit = getIntOr(query, "limit", defaultPageParameter.Limit)
	parameter.Offset = getIntOr(query, "offset", defaultPageParameter.Offset)
	parameter.Header = getIntOr(query, "header", defaultPageParameter.Header)
	parameter.Last = getIntOr(query, "last", defaultPageParameter.Last)
	return parameter
}

// GetMongoQuery Finding query and pick up mongo query
func GetMongoQuery(query bson.M) bson.M {
	var q = bson.M{}
	_q, ok := query["q"]
	if !ok {
		return q
	}
	if err := json.Unmarshal([]byte(_q.(string)), &q); err != nil {
		log.Println("mongo query must JSON format. ", q)
		return bson.M{}
	}
	return q
}
