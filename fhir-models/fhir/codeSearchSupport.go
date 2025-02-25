// Copyright 2019 - 2022 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// CodeSearchSupport is documented here http://hl7.org/fhir/ValueSet/code-search-support
type CodeSearchSupport int

const (
	CodeSearchSupportInCompose CodeSearchSupport = iota
	CodeSearchSupportInExpansion
	CodeSearchSupportInComposeOrExpansion
)

func (code CodeSearchSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CodeSearchSupport) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "in-compose":
		*code = CodeSearchSupportInCompose
	case "in-expansion":
		*code = CodeSearchSupportInExpansion
	case "in-compose-or-expansion":
		*code = CodeSearchSupportInComposeOrExpansion
	default:
		return fmt.Errorf("unknown CodeSearchSupport code `%s`", s)
	}
	return nil
}
func (code CodeSearchSupport) String() string {
	return code.Code()
}
func (code CodeSearchSupport) Code() string {
	switch code {
	case CodeSearchSupportInCompose:
		return "in-compose"
	case CodeSearchSupportInExpansion:
		return "in-expansion"
	case CodeSearchSupportInComposeOrExpansion:
		return "in-compose-or-expansion"
	}
	return "<unknown>"
}
func (code CodeSearchSupport) Display() string {
	switch code {
	case CodeSearchSupportInCompose:
		return "In Compose"
	case CodeSearchSupportInExpansion:
		return "In Expansion"
	case CodeSearchSupportInComposeOrExpansion:
		return "In Compose Or Expansion"
	}
	return "<unknown>"
}
func (code CodeSearchSupport) Definition() string {
	switch code {
	case CodeSearchSupportInCompose:
		return "The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition of the ValueSet."
	case CodeSearchSupportInExpansion:
		return "The search for code on ValueSet returns ValueSet resources where the code is contained in the  ValueSet expansion."
	case CodeSearchSupportInComposeOrExpansion:
		return "The search for code on ValueSet returns ValueSet resources where the code is included in the extensional definition or contained in the ValueSet expansion."
	}
	return "<unknown>"
}
