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

// SubscriptionStatus is documented here http://hl7.org/fhir/StructureDefinition/SubscriptionStatus
type SubscriptionStatus struct {
	Id                           *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                         *Meta                                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules                *string                               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                     *string                               `bson:"language,omitempty" json:"language,omitempty"`
	Text                         *Narrative                            `bson:"text,omitempty" json:"text,omitempty"`
	Extension                    []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension            []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status                       *SubscriptionStatusCodes              `bson:"status,omitempty" json:"status,omitempty"`
	Type                         SubscriptionNotificationType          `bson:"type" json:"type"`
	EventsSinceSubscriptionStart *integer64                            `bson:"eventsSinceSubscriptionStart,omitempty" json:"eventsSinceSubscriptionStart,omitempty"`
	NotificationEvent            []SubscriptionStatusNotificationEvent `bson:"notificationEvent,omitempty" json:"notificationEvent,omitempty"`
	Subscription                 Reference                             `bson:"subscription" json:"subscription"`
	Topic                        *string                               `bson:"topic,omitempty" json:"topic,omitempty"`
	Error                        []CodeableConcept                     `bson:"error,omitempty" json:"error,omitempty"`
}
type SubscriptionStatusNotificationEvent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	EventNumber       integer64   `bson:"eventNumber" json:"eventNumber"`
	Timestamp         *string     `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Focus             *Reference  `bson:"focus,omitempty" json:"focus,omitempty"`
	AdditionalContext []Reference `bson:"additionalContext,omitempty" json:"additionalContext,omitempty"`
}
type OtherSubscriptionStatus SubscriptionStatus

// MarshalJSON marshals the given SubscriptionStatus as JSON into a byte slice
func (r SubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscriptionStatus
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscriptionStatus: OtherSubscriptionStatus(r),
		ResourceType:            "SubscriptionStatus",
	})
}

// UnmarshalSubscriptionStatus unmarshals a SubscriptionStatus.
func UnmarshalSubscriptionStatus(b []byte) (SubscriptionStatus, error) {
	var subscriptionStatus SubscriptionStatus
	if err := json.Unmarshal(b, &subscriptionStatus); err != nil {
		return subscriptionStatus, err
	}
	return subscriptionStatus, nil
}
