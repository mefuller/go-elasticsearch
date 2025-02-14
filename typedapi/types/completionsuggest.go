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

// CompletionSuggest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_global/search/_types/suggester.ts#L48-L53
type CompletionSuggest struct {
	Length  int                       `json:"length"`
	Offset  int                       `json:"offset"`
	Options []CompletionSuggestOption `json:"options"`
	Text    string                    `json:"text"`
}

// CompletionSuggestBuilder holds CompletionSuggest struct and provides a builder API.
type CompletionSuggestBuilder struct {
	v *CompletionSuggest
}

// NewCompletionSuggest provides a builder for the CompletionSuggest struct.
func NewCompletionSuggestBuilder() *CompletionSuggestBuilder {
	r := CompletionSuggestBuilder{
		&CompletionSuggest{},
	}

	return &r
}

// Build finalize the chain and returns the CompletionSuggest struct
func (rb *CompletionSuggestBuilder) Build() CompletionSuggest {
	return *rb.v
}

func (rb *CompletionSuggestBuilder) Length(length int) *CompletionSuggestBuilder {
	rb.v.Length = length
	return rb
}

func (rb *CompletionSuggestBuilder) Offset(offset int) *CompletionSuggestBuilder {
	rb.v.Offset = offset
	return rb
}

func (rb *CompletionSuggestBuilder) Options(options []CompletionSuggestOptionBuilder) *CompletionSuggestBuilder {
	tmp := make([]CompletionSuggestOption, len(options))
	for _, value := range options {
		tmp = append(tmp, value.Build())
	}
	rb.v.Options = tmp
	return rb
}

func (rb *CompletionSuggestBuilder) Text(text string) *CompletionSuggestBuilder {
	rb.v.Text = text
	return rb
}
