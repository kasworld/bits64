package bits64

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	var b Bits64
	t.Logf("%v", b)

	(&b).SetBit(5)
	t.Logf("%v", b)
	(&b).SetBits(0x0f0f0f)
	t.Logf("%v", b)

	(&b).ClearBit(2)
	t.Logf("%v", b)
	(&b).ClearBits(0x0f)
	t.Logf("%v", b)

	(&b).NegBit(5)
	t.Logf("%v", b)
	(&b).NegBits(0xf0f0f0f)
	t.Logf("%v", b)

	t.Logf("%v", (&b).GetBit(7))
	t.Logf("%v", (&b).GetBits(0x704))
}
