// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import "encoding/json"

type SupplyRequest struct {
	Id                    string                      `json:"id" bson:"_id"`
	Patient               *Reference                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Source                *Reference                  `bson:"source,omitempty" json:"source,omitempty"`
	Date                  *FHIRDateTime               `bson:"date,omitempty" json:"date,omitempty"`
	Identifier            *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                      `bson:"status,omitempty" json:"status,omitempty"`
	Kind                  *CodeableConcept            `bson:"kind,omitempty" json:"kind,omitempty"`
	OrderedItem           *Reference                  `bson:"orderedItem,omitempty" json:"orderedItem,omitempty"`
	Supplier              []Reference                 `bson:"supplier,omitempty" json:"supplier,omitempty"`
	ReasonCodeableConcept *CodeableConcept            `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                  `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	When                  *SupplyRequestWhenComponent `bson:"when,omitempty" json:"when,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *SupplyRequest) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		SupplyRequest
	}{
		ResourceType:  "SupplyRequest",
		SupplyRequest: *resource,
	}
	return json.Marshal(x)
}

type SupplyRequestWhenComponent struct {
	Code     *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Schedule *Timing          `bson:"schedule,omitempty" json:"schedule,omitempty"`
}
