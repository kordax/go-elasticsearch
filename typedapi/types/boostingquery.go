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
// https://github.com/elastic/elasticsearch-specification/tree/1ad7fe36297b3a8e187b2259dedaf68a47bc236e

package types

// BoostingQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/_types/query_dsl/compound.ts#L36-L40
type BoostingQuery struct {
	Boost         *float32 `json:"boost,omitempty"`
	Negative      *Query   `json:"negative,omitempty"`
	NegativeBoost Float64  `json:"negative_boost"`
	Positive      *Query   `json:"positive,omitempty"`
	QueryName_    *string  `json:"_name,omitempty"`
}

// NewBoostingQuery returns a BoostingQuery.
func NewBoostingQuery() *BoostingQuery {
	r := &BoostingQuery{}

	return r
}