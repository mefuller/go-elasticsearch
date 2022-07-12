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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// ParentAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/bucket.ts#L281-L283
type ParentAggregation struct {
	Meta *Metadata     `json:"meta,omitempty"`
	Name *string       `json:"name,omitempty"`
	Type *RelationName `json:"type,omitempty"`
}

// ParentAggregationBuilder holds ParentAggregation struct and provides a builder API.
type ParentAggregationBuilder struct {
	v *ParentAggregation
}

// NewParentAggregation provides a builder for the ParentAggregation struct.
func NewParentAggregationBuilder() *ParentAggregationBuilder {
	r := ParentAggregationBuilder{
		&ParentAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the ParentAggregation struct
func (rb *ParentAggregationBuilder) Build() ParentAggregation {
	return *rb.v
}

func (rb *ParentAggregationBuilder) Meta(meta *MetadataBuilder) *ParentAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *ParentAggregationBuilder) Name(name string) *ParentAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *ParentAggregationBuilder) Type_(type_ RelationName) *ParentAggregationBuilder {
	rb.v.Type = &type_
	return rb
}