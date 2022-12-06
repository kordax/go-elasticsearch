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

// DataPathStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/nodes/_types/Stats.ts#L223-L240
type DataPathStats struct {
	Available            *string `json:"available,omitempty"`
	AvailableInBytes     *int64  `json:"available_in_bytes,omitempty"`
	DiskQueue            *string `json:"disk_queue,omitempty"`
	DiskReadSize         *string `json:"disk_read_size,omitempty"`
	DiskReadSizeInBytes  *int64  `json:"disk_read_size_in_bytes,omitempty"`
	DiskReads            *int64  `json:"disk_reads,omitempty"`
	DiskWriteSize        *string `json:"disk_write_size,omitempty"`
	DiskWriteSizeInBytes *int64  `json:"disk_write_size_in_bytes,omitempty"`
	DiskWrites           *int64  `json:"disk_writes,omitempty"`
	Free                 *string `json:"free,omitempty"`
	FreeInBytes          *int64  `json:"free_in_bytes,omitempty"`
	Mount                *string `json:"mount,omitempty"`
	Path                 *string `json:"path,omitempty"`
	Total                *string `json:"total,omitempty"`
	TotalInBytes         *int64  `json:"total_in_bytes,omitempty"`
	Type                 *string `json:"type,omitempty"`
}

// NewDataPathStats returns a DataPathStats.
func NewDataPathStats() *DataPathStats {
	r := &DataPathStats{}

	return r
}