// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type TestStruct struct {
	// Value types with omitempty - should NOT report (value check disabled)
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	// Pointer types without omitempty - should NOT report (pointer check disabled)
	Title       *string `json:"title"`
	Description *string `json:"description"`

	// Composite types - no restrictions
	Tags []string `json:"tags"`
}

func main() {}
