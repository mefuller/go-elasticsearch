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

package migratetodatatiers

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package migratetodatatiers
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/ilm/migrate_to_data_tiers/Request.ts#L22-L44
type Request struct {
	LegacyTemplateToDelete *string `json:"legacy_template_to_delete,omitempty"`

	NodeAttribute *string `json:"node_attribute,omitempty"`
}

// RequestBuilder is the builder API for the migratetodatatiers.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Migratetodatatiers request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) LegacyTemplateToDelete(legacytemplatetodelete string) *RequestBuilder {
	rb.v.LegacyTemplateToDelete = &legacytemplatetodelete
	return rb
}

func (rb *RequestBuilder) NodeAttribute(nodeattribute string) *RequestBuilder {
	rb.v.NodeAttribute = &nodeattribute
	return rb
}
