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

// StemmerOverrideTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/_types/analysis/token_filters.ts#L313-L317
type StemmerOverrideTokenFilter struct {
	Rules     []string `json:"rules,omitempty"`
	RulesPath *string  `json:"rules_path,omitempty"`
	Type      string   `json:"type,omitempty"`
	Version   *string  `json:"version,omitempty"`
}

// NewStemmerOverrideTokenFilter returns a StemmerOverrideTokenFilter.
func NewStemmerOverrideTokenFilter() *StemmerOverrideTokenFilter {
	r := &StemmerOverrideTokenFilter{}

	r.Type = "stemmer_override"

	return r
}