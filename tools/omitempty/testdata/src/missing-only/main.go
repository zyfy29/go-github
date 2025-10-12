// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type TestStruct struct {
	// Value types with omitempty - should NOT report (unnecessary check disabled)
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Pointer types without omitempty - should report (missing check enabled)
	Title       *string `json:"title"`       // want `omitempty is missing for pointer types`
	Description *string `json:"description"` // want `omitempty is missing for pointer types`

	// Other types - no restrictions
	Tags []string `json:"tags"`
}

func main() {}
