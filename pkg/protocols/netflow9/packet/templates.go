// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package packet

import "unsafe"

const (
	// numPreAllocFlowDataRecs is number of elements to pre allocate in DataRecs slice
	numPreAllocFlowDataRecs = 20
)

var (
	sizeOfTemplateRecordHeader        = unsafe.Sizeof(TemplateRecordHeader{})
	sizeOfOptionsTemplateRecordHeader = unsafe.Sizeof(OptionsTemplateRecordHeader{})
	sizeOfOptionScope                 = unsafe.Sizeof(OptionScope{})
)

// TemplateRecordHeader represents the header of a template record
type TemplateRecordHeader struct {
	// Number of fields in this Template Record. Because a Template FlowSet
	// usually contains multiple Template Records, this field allows the
	// Collector to determine the end of the current Template Record and
	// the start of the next.
	FieldCount uint16

	// Each of the newly generated Template Records is given a unique
	// Template ID. This uniqueness is local to the Observation Domain that
	// generated the Template ID. Template IDs of Data FlowSets are numbered
	// from 256 to 65535.
	TemplateID uint16
}

// OptionsTemplateRecordHeader represents the header of an option template record
type OptionsTemplateRecordHeader struct {
	// The length (in bytes) of any options field definitions
	// contained in this Options Template Record.
	OptionLength uint16

	// Number of fields in this Template Record. Because a Template FlowSet
	// usually contains multiple Template Records, this field allows the
	// Collector to determine the end of the current Template Record and
	// the start of the next.
	OptionScopeLength uint16

	// Each of the newly generated Template Records is given a unique
	// Template ID. This uniqueness is local to the Observation Domain that
	// generated the Template ID. Template IDs of Data FlowSets are numbered
	// from 256 to 65535.
	TemplateID uint16
}

// TemplateRecords is a single template that describes structure of a Flow Record
// (actual Netflow data).
type TemplateRecords struct {
	Header *TemplateRecordHeader

	// List of scopes
	OptionScopes []*OptionScope

	// List of fields in this Template Record.
	Records []*TemplateRecord

	Packet *Packet

	Values [][]byte
}

// OptionScope represents an option scope in an options template flowset
type OptionScope struct {
	// The length (in bytes) of the Scope field, as it would appear in
	//an Options Data Record.
	ScopeFieldLength uint16

	//A numeric value that represents the type of field that would
	//appear in the Options Template Record.  Refer to the Field Type
	//Definitions section.
	ScopeFieldType uint16
}

//TemplateRecord represents a Template Record as described in RFC3954
type TemplateRecord struct {
	// The length (in bytes) of the field.
	Length uint16

	// A numeric value that represents the type of field.
	Type uint16
}

// FlowDataRecord is actual NetFlow data. This structure does not contain any
// information about the actual data meaning. It must be combined with
// corresponding TemplateRecord to be decoded to a single NetFlow data row.
type FlowDataRecord struct {
	// List of Flow Data Record values stored in raw format as []byte
	Values [][]byte
}

// sizeOfTemplateRecord is the raw size of a TemplateRecord
var sizeOfTemplateRecord = unsafe.Sizeof(TemplateRecord{})

// DecodeFlowSet uses current TemplateRecord to decode data in Data FlowSet to
// a list of Flow Data Records.
/*func (dtpl *TemplateRecords) DecodeFlowSet(set FlowSet) (list []FlowDataRecord) {
	if set.Header.FlowSetID != dtpl.Header.TemplateID {
		return nil
	}
	var record FlowDataRecord

	// Pre-allocate some room for flows
	list = make([]FlowDataRecord, 0, numPreAllocFlowDataRecs)

	// Assume total record length must be >= 4, otherwise it is impossible
	// to distinguish between padding and new record. Padding MUST be
	// supported.
	n := len(set.Flows)
	count := 0

	for n >= 4 {
		record.Values, count = parseFieldValues(set.Flows[0:n], dtpl.Records)
		if record.Values == nil {
			return
		}
		list = append(list, record)
		n = n - count
	}

	return
}*/

// DecodeFlowSet uses current TemplateRecord to decode data in Data FlowSet to
// a list of Flow Data Records.
func DecodeFlowSet(templateRecords []*TemplateRecord, set FlowSet) (list []FlowDataRecord) {
	var record FlowDataRecord

	// Pre-allocate some room for flows
	list = make([]FlowDataRecord, 0, numPreAllocFlowDataRecs)

	// Assume total record length must be >= 4, otherwise it is impossible
	// to distinguish between padding and new record. Padding MUST be
	// supported.
	n := len(set.Flows)
	count := 0

	for n >= 4 {
		record.Values, count = parseFieldValues(set.Flows[0:n], templateRecords)
		if record.Values == nil {
			return
		}
		list = append(list, record)
		n = n - count
	}

	return
}

// parseFieldValues reads actual fields values from a Data Record utilizing a template
func parseFieldValues(flows []byte, fields []*TemplateRecord) ([][]byte, int) {
	count := 0
	n := len(flows)
	values := make([][]byte, len(fields))
	for i, f := range fields {
		if n < int(f.Length) {
			return nil, 0
		}
		values[i] = flows[n-int(f.Length) : n]
		count += int(f.Length)
		n -= int(f.Length)
	}
	return values, count
}
