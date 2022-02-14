//
// Copyright 2021 The Sigstore Authors.
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

package sharding

import "testing"

func TestLogRanges_ResolveVirtualIndex(t *testing.T) {
	lrs := LogRanges{
		Ranges: []LogRange{
			{TreeID: 1, TreeLength: 17},
			{TreeID: 2, TreeLength: 1},
			{TreeID: 3, TreeLength: 100},
			{TreeID: 4},
		},
	}

	for _, tt := range []struct {
		Index      int
		WantTreeID uint64
		WantIndex  uint64
	}{
		{
			Index:      3,
			WantTreeID: 1, WantIndex: 3,
		},
		// This is the first (0th) entry in the next tree
		{
			Index:      17,
			WantTreeID: 2, WantIndex: 0,
		},
		// Overflow
		{
			Index:      3000,
			WantTreeID: 4, WantIndex: 2882,
		},
	} {
		tree, index := lrs.ResolveVirtualIndex(tt.Index)
		if tree != tt.WantTreeID {
			t.Errorf("LogRanges.ResolveVirtualIndex() tree = %v, want %v", tree, tt.WantTreeID)
		}
		if index != tt.WantIndex {
			t.Errorf("LogRanges.ResolveVirtualIndex() index = %v, want %v", index, tt.WantIndex)
		}
	}
}

func TestLogRanges_IsEmpty(t *testing.T) {
	emptyLogRanges := CreateEmptyLogRanges()
	if !emptyLogRanges.IsEmpty() {
		t.Errorf("Empty logRanges struct tests as non-empty")
	}

	logRange := logRange {TreeID: 1, TreeLength: 2}
	nonEmptyLogRanges := append(emptyLogRanges.Ranges, logRange)

	if nonEmptyLoganges.IsEmpty() {
		t.Errorf("Non-empty logRanges struct tests as empty")
	}
}

func TestLogRanges_IsValid(t *testing.T) {
	emptyLogRanges := CreateEmptyLogRanges()
	if emptyLogRanges.IsValid() {
		t.Errorf("Empty logRanges struct tests as valid")
	}

	logRangeInvalid := logRange {TreeID: 0, TreeLength: 7}
	logRangeValid := logRange {TreeID: 7:, TreeLength: 0}

	validLogRanges := append(emptyLogRanges.Ranges, logRangeValid)
	if !validLogRanges.IsValid() {
		t.Errorf("Valid logRanges struct tests as invalid")
	}

	invalidLogRanges := append(validLogRanges.Ranges, logRangeInvalid)
	if invalidLogRanges.IsValid() {
		t.Errorf("Invalid logRanges struct tests as valid")
	}
}
