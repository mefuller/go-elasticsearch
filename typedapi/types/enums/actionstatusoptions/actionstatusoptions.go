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

// Package actionstatusoptions
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/watcher/_types/Action.ts#L99-L104
package actionstatusoptions

import "strings"

type ActionStatusOptions struct {
	name string
}

var (
	Success = ActionStatusOptions{"success"}

	Failure = ActionStatusOptions{"failure"}

	Simulated = ActionStatusOptions{"simulated"}

	Throttled = ActionStatusOptions{"throttled"}
)

func (a ActionStatusOptions) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *ActionStatusOptions) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "success":
		*a = Success
	case "failure":
		*a = Failure
	case "simulated":
		*a = Simulated
	case "throttled":
		*a = Throttled
	default:
		*a = ActionStatusOptions{string(text)}
	}

	return nil
}

func (a ActionStatusOptions) String() string {
	return a.name
}
