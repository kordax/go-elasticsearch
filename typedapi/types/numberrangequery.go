// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// NumberRangeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_types/query_dsl/term.ts#L83-L90
type NumberRangeQuery struct {
	Boost      *float32                     `json:"boost,omitempty"`
	From       float64                      `json:"from,omitempty"`
	Gt         *float64                     `json:"gt,omitempty"`
	Gte        *float64                     `json:"gte,omitempty"`
	Lt         *float64                     `json:"lt,omitempty"`
	Lte        *float64                     `json:"lte,omitempty"`
	QueryName_ *string                      `json:"_name,omitempty"`
	Relation   *rangerelation.RangeRelation `json:"relation,omitempty"`
	To         float64                      `json:"to,omitempty"`
}

// NewNumberRangeQuery returns a NumberRangeQuery.
func NewNumberRangeQuery() *NumberRangeQuery {
	r := &NumberRangeQuery{}

	return r
}