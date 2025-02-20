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
// https://github.com/elastic/elasticsearch-specification/tree/e16324dcde9297dd1149c1ef3d6d58afe272e646

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sampleraggregationexecutionhint"
)

// DiversifiedSamplerAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e16324dcde9297dd1149c1ef3d6d58afe272e646/specification/_types/aggregations/bucket.ts#L320-L341
type DiversifiedSamplerAggregation struct {
	// ExecutionHint The type of value used for de-duplication.
	ExecutionHint *sampleraggregationexecutionhint.SamplerAggregationExecutionHint `json:"execution_hint,omitempty"`
	// Field The field used to provide values used for de-duplication.
	Field *string `json:"field,omitempty"`
	// MaxDocsPerValue Limits how many documents are permitted per choice of de-duplicating value.
	MaxDocsPerValue *int     `json:"max_docs_per_value,omitempty"`
	Meta            Metadata `json:"meta,omitempty"`
	Name            *string  `json:"name,omitempty"`
	Script          Script   `json:"script,omitempty"`
	// ShardSize Limits how many top-scoring documents are collected in the sample processed
	// on each shard.
	ShardSize *int `json:"shard_size,omitempty"`
}

func (s *DiversifiedSamplerAggregation) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "max_docs_per_value":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxDocsPerValue = &value
			case float64:
				f := int(v)
				s.MaxDocsPerValue = &f
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "script":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return err
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}

				switch t {

				case "lang", "options", "source":
					o := NewInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				}
			}

		case "shard_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		}
	}
	return nil
}

// NewDiversifiedSamplerAggregation returns a DiversifiedSamplerAggregation.
func NewDiversifiedSamplerAggregation() *DiversifiedSamplerAggregation {
	r := &DiversifiedSamplerAggregation{}

	return r
}
