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

// UserProfileHitMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/security/_types/UserProfile.ts#L29-L32
type UserProfileHitMetadata struct {
	PrimaryTerm_ int64          `json:"_primary_term"`
	SeqNo_       SequenceNumber `json:"_seq_no"`
}

// UserProfileHitMetadataBuilder holds UserProfileHitMetadata struct and provides a builder API.
type UserProfileHitMetadataBuilder struct {
	v *UserProfileHitMetadata
}

// NewUserProfileHitMetadata provides a builder for the UserProfileHitMetadata struct.
func NewUserProfileHitMetadataBuilder() *UserProfileHitMetadataBuilder {
	r := UserProfileHitMetadataBuilder{
		&UserProfileHitMetadata{},
	}

	return &r
}

// Build finalize the chain and returns the UserProfileHitMetadata struct
func (rb *UserProfileHitMetadataBuilder) Build() UserProfileHitMetadata {
	return *rb.v
}

func (rb *UserProfileHitMetadataBuilder) PrimaryTerm_(primaryterm_ int64) *UserProfileHitMetadataBuilder {
	rb.v.PrimaryTerm_ = primaryterm_
	return rb
}

func (rb *UserProfileHitMetadataBuilder) SeqNo_(seqno_ SequenceNumber) *UserProfileHitMetadataBuilder {
	rb.v.SeqNo_ = seqno_
	return rb
}
