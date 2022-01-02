package common

import "time"

type Query struct {
	creationTime time.Time
	queryType    string
}

func NewQuery(name string) *Query {
	return &Query{time.Now(), name}
}

func (c *Query) GetCreationTime() time.Time {
	return c.creationTime
}

func (c *Query) GetQueryType() string {
	return c.queryType
}
