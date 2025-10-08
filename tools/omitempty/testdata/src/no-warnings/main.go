// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "time"

type TestStruct struct {
	// Value types without omitempty - correct
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	Score     float64   `json:"score"`
	Count     int64     `json:"count"`
	CreatedAt time.Time `json:"created_at"`

	// Pointer types with omitempty - correct
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Age         *int    `json:"age,omitempty"`

	// Composite types - both with and without omitempty are allowed
	Tags         []string          `json:"tags"`
	TagsOmit     []string          `json:"tags_omit,omitempty"`
	Metadata     map[string]string `json:"metadata"`
	MetadataOmit map[string]string `json:"metadata_omit,omitempty"`
	Data         []byte            `json:"data"`
	DataOmit     []byte            `json:"data_omit,omitempty"`

	// Fields without JSON tags - should be ignored
	Internal string

	// JSON tag with "-" - should be ignored
	Ignored string `json:"-"`
}

func main() {}
