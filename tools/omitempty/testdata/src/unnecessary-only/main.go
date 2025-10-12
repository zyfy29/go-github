// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "time"

type CustomType int

type TestStruct struct {
	// Value types with omitempty - should report (unnecessary check enabled)
	ID   int        `json:"id,omitempty"`   // want `omitempty is unnecessary for value types`
	Name string     `json:"name,omitempty"` // want `omitempty is unnecessary for value types`
	Time time.Time  `json:"time,omitempty"` // want `omitempty is unnecessary for value types`
	Age  CustomType `json:"age,omitempty"`  // want `omitempty is unnecessary for value types`
	Any  any        `json:"any,omitempty"`  // want `omitempty is unnecessary for value types`

	// Pointer types without omitempty - should NOT report (missing check disabled)
	Title       *string `json:"title"`
	Description *string `json:"description"`

	// Other types - no restrictions
	Tags []string `json:"tags"`
}

func main() {}
