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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MedicationRequest is documented here http://hl7.org/fhir/StructureDefinition/MedicationRequest
type MedicationRequest struct {
	Id                        *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	Extension                 []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn                   []Reference                       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PriorPrescription         *Reference                        `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
	GroupIdentifier           *Identifier                       `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                    MedicationrequestStatus           `bson:"status" json:"status"`
	StatusReason              *CodeableConcept                  `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	StatusChanged             *string                           `bson:"statusChanged,omitempty" json:"statusChanged,omitempty"`
	Intent                    MedicationRequestIntent           `bson:"intent" json:"intent"`
	Category                  []CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Priority                  *RequestPriority                  `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform              *bool                             `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Medication                CodeableReference                 `bson:"medication" json:"medication"`
	Subject                   Reference                         `bson:"subject" json:"subject"`
	InformationSource         []Reference                       `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	Encounter                 *Reference                        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	SupportingInformation     []Reference                       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	AuthoredOn                *string                           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester                 *Reference                        `bson:"requester,omitempty" json:"requester,omitempty"`
	Reported                  *bool                             `bson:"reported,omitempty" json:"reported,omitempty"`
	PerformerType             *CodeableConcept                  `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Performer                 []Reference                       `bson:"performer,omitempty" json:"performer,omitempty"`
	Device                    []CodeableReference               `bson:"device,omitempty" json:"device,omitempty"`
	Recorder                  *Reference                        `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Reason                    []CodeableReference               `bson:"reason,omitempty" json:"reason,omitempty"`
	CourseOfTherapyType       *CodeableConcept                  `bson:"courseOfTherapyType,omitempty" json:"courseOfTherapyType,omitempty"`
	Insurance                 []Reference                       `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Note                      []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
	RenderedDosageInstruction *string                           `bson:"renderedDosageInstruction,omitempty" json:"renderedDosageInstruction,omitempty"`
	EffectiveDosePeriod       *Period                           `bson:"effectiveDosePeriod,omitempty" json:"effectiveDosePeriod,omitempty"`
	DosageInstruction         []Dosage                          `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationRequestDispenseRequest `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	Substitution              *MedicationRequestSubstitution    `bson:"substitution,omitempty" json:"substitution,omitempty"`
	EventHistory              []Reference                       `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationRequestDispenseRequest struct {
	Id                     *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	InitialFill            *MedicationRequestDispenseRequestInitialFill `bson:"initialFill,omitempty" json:"initialFill,omitempty"`
	DispenseInterval       *Duration                                    `bson:"dispenseInterval,omitempty" json:"dispenseInterval,omitempty"`
	ValidityPeriod         *Period                                      `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed *int                                         `bson:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed,omitempty"`
	Quantity               *Quantity                                    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ExpectedSupplyDuration *Duration                                    `bson:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration,omitempty"`
	Dispenser              *Reference                                   `bson:"dispenser,omitempty" json:"dispenser,omitempty"`
	DispenserInstruction   []Annotation                                 `bson:"dispenserInstruction,omitempty" json:"dispenserInstruction,omitempty"`
	DoseAdministrationAid  *CodeableConcept                             `bson:"doseAdministrationAid,omitempty" json:"doseAdministrationAid,omitempty"`
}
type MedicationRequestDispenseRequestInitialFill struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Duration          *Duration   `bson:"duration,omitempty" json:"duration,omitempty"`
}
type MedicationRequestSubstitution struct {
	Id                     *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	AllowedBoolean         bool             `bson:"allowedBoolean" json:"allowedBoolean"`
	AllowedCodeableConcept CodeableConcept  `bson:"allowedCodeableConcept" json:"allowedCodeableConcept"`
	Reason                 *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}
type OtherMedicationRequest MedicationRequest

// MarshalJSON marshals the given MedicationRequest as JSON into a byte slice
func (r MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationRequest: OtherMedicationRequest(r),
		ResourceType:           "MedicationRequest",
	})
}

// UnmarshalMedicationRequest unmarshals a MedicationRequest.
func UnmarshalMedicationRequest(b []byte) (MedicationRequest, error) {
	var medicationRequest MedicationRequest
	if err := json.Unmarshal(b, &medicationRequest); err != nil {
		return medicationRequest, err
	}
	return medicationRequest, nil
}
