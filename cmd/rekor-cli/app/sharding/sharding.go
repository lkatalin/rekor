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

import (
	"strconv"
	//"github.com/sigstore/rekor/cmd/rekor-server/app"
)

// A FullID refers to a specific artifact's ID and is made of two components,
// the ShardID and the UUID, separated by a character, the FullIDSeparator.
// The ShardID is in human-readable decimal and refers to the specific log or
// shard where the artifact can be found. The UUID refers to the hex-encoded
// merkle leaf hash from trillian for the specific artifact.

const ShardIDMax = 999999
const FullIDSeparator = "-"

const ShardIDLen = 6
const UUIDLen = 64
const SeparatorLen = len(FullIDSeparator)
const FullIDLen = ShardIDLen + UUIDLen + SeparatorLen

// A ShardID is a number with specific properties, including
// its representation as a six-digit string.
type ShardID struct {
	ShardIDInt    uint64
	ShardIDString string
}

// Create default values
func NewCurrent(current uint64) ShardID {
	return ShardID{
		ShardIDInt:    current,
		ShardIDString: strconv.FormatUint(current, 10)}
}

//var CurrentShardID = app.flags.ActiveIndex()