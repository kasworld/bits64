// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bits64

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	var b Bits64
	t.Logf("%v", b)

	(&b).SetBit(5)
	t.Logf("SetBit(5) %v", b)
	(&b).SetBits(0x0f0f0f)
	t.Logf("SetBits(0x0f0f0f) %v", b)

	(&b).ClearBit(2)
	t.Logf("ClearBit(2) %v", b)
	(&b).ClearBits(0x0f)
	t.Logf("ClearBits(0x0f) %v", b)

	b.SetBits(0)
	t.Logf("%v", b)
	(&b).NegBit(5)
	t.Logf("NegBit(5) %v", b)
	(&b).NegBit(5)
	t.Logf("NegBit(5) %v", b)
	(&b).NegBits(0xf0f0f0f)
	t.Logf("NegBits(0xf0f0f0f) %v", b)

	b.SetBits(0)
	t.Logf("GetBit(7) %v", (&b).GetBit(7))
	t.Logf("GetBits(0x704) %v", (&b).GetBits(0x704))
}
