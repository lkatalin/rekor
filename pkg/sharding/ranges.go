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

type LogRanges struct {
	Ranges []LogRange
}

type LogRange struct {
	TreeID     uint64
	TreeLength uint64
}

func CreateEmptyLogRanges() LogRanges {
	return LogRanges {
		Ranges: []LogRange{},
	}
}

func CreateLogRangesFromRanges(r []LogRange) LogRanges {
	return LogRanges {
		Ranges: r}
}

func (l *LogRanges) IsEmpty() bool {
	return len(l.Ranges) == 0
}

func (l *LogRanges) IsValid() bool {
	// Test whether struct is empty
	if l.IsEmpty() {
		return false
	}

	// Test whether any TreeIDs are zero or ranges
	// have null data
	for _, r := range l.Ranges {
		if r.TreeID == 0 {
			return false
		}

		if r == (LogRange{}) {
			return false
		}
	}

	// TODO: more invalid cases here?

	return true
}

func (l *LogRanges) ResolveVirtualIndex(index int) (uint64, uint64) {
	indexLeft := index
	for _, l := range l.Ranges {
		if indexLeft < int(l.TreeLength) {
			return l.TreeID, uint64(indexLeft)
		}
		indexLeft -= int(l.TreeLength)
	}

	// Return the last one!
	return l.Ranges[len(l.Ranges)-1].TreeID, uint64(indexLeft)
}

// ActiveIndex returns the active shard index, always the last shard in the range
func (l *LogRanges) ActiveIndex() uint64 {
	if len(l.Ranges) == 0 {
		return 0
	} else {
		return l.Ranges[len(l.Ranges)-1].TreeID
	}
}
