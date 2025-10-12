// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"time"
)

type TestStruct struct {
	// Basic value types with omitempty - should report
	ID      int     `json:"id,omitempty"`     // want `omitempty is unnecessary for value types`
	Name    string  `json:"name,omitempty"`   // want `omitempty is unnecessary for value types`
	Active  bool    `json:"active,omitempty"` // want `omitempty is unnecessary for value types`
	Score   float64 `json:"score,omitempty"`  // want `omitempty is unnecessary for value types`
	Count   int64   `json:"count,omitempty"`  // want `omitempty is unnecessary for value types`
	ByteVal byte    `json:"byte,omitempty"`   // want `omitempty is unnecessary for value types`
	RuneVal rune    `json:"rune,omitempty"`   // want `omitempty is unnecessary for value types`
	UintVal uint    `json:"uint,omitempty"`   // want `omitempty is unnecessary for value types`

	// Named types and package-qualified types with omitempty - should report
	CreatedAt time.Time `json:"created_at,omitempty"` // want `omitempty is unnecessary for value types`
	AnyValue  any       `json:"any,omitempty"`        // want `omitempty is unnecessary for value types`

	// Pointer types without omitempty - should report
	Title       *string `json:"title"`       // want `omitempty is missing for pointer types`
	Description *string `json:"description"` // want `omitempty is missing for pointer types`
	Age         *int    `json:"age"`         // want `omitempty is missing for pointer types`

	// Other types - no restrictions, these should NOT report
	Tags     []string          `json:"tags"`
	Metadata map[string]string `json:"metadata"`
	Data     []byte            `json:"data"`

	// json.RawMessage is excluded - should NOT report
	RawData json.RawMessage `json:"raw,omitempty"`
}

func main() {}
