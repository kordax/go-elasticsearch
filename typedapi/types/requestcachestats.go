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

// RequestCacheStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/_types/Stats.ts#L177-L183
type RequestCacheStats struct {
	Evictions         int64   `json:"evictions"`
	HitCount          int64   `json:"hit_count"`
	MemorySize        *string `json:"memory_size,omitempty"`
	MemorySizeInBytes int64   `json:"memory_size_in_bytes"`
	MissCount         int64   `json:"miss_count"`
}

// NewRequestCacheStats returns a RequestCacheStats.
func NewRequestCacheStats() *RequestCacheStats {
	r := &RequestCacheStats{}

	return r
}