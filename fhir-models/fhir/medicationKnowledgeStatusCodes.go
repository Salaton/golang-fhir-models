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

// MedicationKnowledgeStatusCodes is documented here http://hl7.org/fhir/ValueSet/medicationknowledge-status
type MedicationKnowledgeStatusCodes int

const (
	MedicationKnowledgeStatusCodesActive MedicationKnowledgeStatusCodes = iota
	MedicationKnowledgeStatusCodesEnteredInError
	MedicationKnowledgeStatusCodesInactive
)

func (code MedicationKnowledgeStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationKnowledgeStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = MedicationKnowledgeStatusCodesActive
	case "entered-in-error":
		*code = MedicationKnowledgeStatusCodesEnteredInError
	case "inactive":
		*code = MedicationKnowledgeStatusCodesInactive
	default:
		return fmt.Errorf("unknown MedicationKnowledgeStatusCodes code `%s`", s)
	}
	return nil
}
func (code MedicationKnowledgeStatusCodes) String() string {
	return code.Code()
}
func (code MedicationKnowledgeStatusCodes) Code() string {
	switch code {
	case MedicationKnowledgeStatusCodesActive:
		return "active"
	case MedicationKnowledgeStatusCodesEnteredInError:
		return "entered-in-error"
	case MedicationKnowledgeStatusCodesInactive:
		return "inactive"
	}
	return "<unknown>"
}
func (code MedicationKnowledgeStatusCodes) Display() string {
	switch code {
	case MedicationKnowledgeStatusCodesActive:
		return "Active"
	case MedicationKnowledgeStatusCodesEnteredInError:
		return "Entered in Error"
	case MedicationKnowledgeStatusCodesInactive:
		return "Inactive"
	}
	return "<unknown>"
}
func (code MedicationKnowledgeStatusCodes) Definition() string {
	switch code {
	case MedicationKnowledgeStatusCodesActive:
		return "The medication referred to by this MedicationKnowledge is in active use within the drug database or inventory system."
	case MedicationKnowledgeStatusCodesEnteredInError:
		return "The medication referred to by this MedicationKnowledge was entered in error within the drug database or inventory system."
	case MedicationKnowledgeStatusCodesInactive:
		return "The medication referred to by this MedicationKnowledge is not in active use within the drug database or inventory system."
	}
	return "<unknown>"
}
