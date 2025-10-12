// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package omitempty

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRun(t *testing.T) {
	t.Parallel()
	testdata := analysistest.TestData()

	t.Run("default config (both enabled)", func(t *testing.T) {
		t.Parallel()
		plugin, _ := New(nil)
		analyzers, _ := plugin.BuildAnalyzers()
		analysistest.Run(t, testdata, analyzers[0], "has-warnings", "no-warnings")
	})

	t.Run("unnecessary only", func(t *testing.T) {
		t.Parallel()
		plugin, _ := New(map[string]any{
			"unnecessary": true,
			"missing":     false,
		})
		analyzers, _ := plugin.BuildAnalyzers()
		analysistest.Run(t, testdata, analyzers[0], "unnecessary-only")
	})

	t.Run("missing only", func(t *testing.T) {
		t.Parallel()
		plugin, _ := New(map[string]any{
			"unnecessary": false,
			"missing":     true,
		})
		analyzers, _ := plugin.BuildAnalyzers()
		analysistest.Run(t, testdata, analyzers[0], "missing-only")
	})

	t.Run("both disabled", func(t *testing.T) {
		t.Parallel()
		plugin, _ := New(map[string]any{
			"unnecessary": false,
			"missing":     false,
		})
		analyzers, _ := plugin.BuildAnalyzers()
		analysistest.Run(t, testdata, analyzers[0], "both-disabled")
	})
}
