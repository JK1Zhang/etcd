// Copyright 2019 The etcd Authors
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

package quorum

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func BenchmarkMajorityConfig_CommittedIndex(b *testing.B) {
	use_group_commit := false
	// go test -run - -bench . -benchmem ./raft/quorum
	for _, n := range []int{1, 3, 5, 7, 9, 11} {
		b.Run(fmt.Sprintf("voters=%d", n), func(b *testing.B) {
			c := MajorityConfig{}
			l := mapAckIndexer{}
			for i := uint64(0); i < uint64(n); i++ {
				c[i+1] = struct{}{}
				l[i+1] = Index{Index: uint64(rand.Int63n(math.MaxInt64)), Group_id: 0}
			}

			for i := 0; i < b.N; i++ {
				_, _ = c.CommittedIndex(l, use_group_commit)
			}
		})
	}
}
