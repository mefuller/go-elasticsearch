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

// HttpInputBasicAuthentication type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/watcher/_types/Input.ts#L56-L59
type HttpInputBasicAuthentication struct {
	Password Password `json:"password"`
	Username Username `json:"username"`
}

// HttpInputBasicAuthenticationBuilder holds HttpInputBasicAuthentication struct and provides a builder API.
type HttpInputBasicAuthenticationBuilder struct {
	v *HttpInputBasicAuthentication
}

// NewHttpInputBasicAuthentication provides a builder for the HttpInputBasicAuthentication struct.
func NewHttpInputBasicAuthenticationBuilder() *HttpInputBasicAuthenticationBuilder {
	r := HttpInputBasicAuthenticationBuilder{
		&HttpInputBasicAuthentication{},
	}

	return &r
}

// Build finalize the chain and returns the HttpInputBasicAuthentication struct
func (rb *HttpInputBasicAuthenticationBuilder) Build() HttpInputBasicAuthentication {
	return *rb.v
}

func (rb *HttpInputBasicAuthenticationBuilder) Password(password Password) *HttpInputBasicAuthenticationBuilder {
	rb.v.Password = password
	return rb
}

func (rb *HttpInputBasicAuthenticationBuilder) Username(username Username) *HttpInputBasicAuthenticationBuilder {
	rb.v.Username = username
	return rb
}
