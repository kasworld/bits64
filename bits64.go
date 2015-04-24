//goroutine safe uint64 bit operation
package bits64

import (
	"fmt"
	"sync/atomic"
)

type Bits64 uint64

func (t Bits64) String() string {
	return fmt.Sprintf("0b%0b", t)
}

func (bt *Bits64) SetBit(n uint) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval | (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) ClearBit(n uint) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval &^ (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) NegBit(n uint) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval ^ (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) GetBit(n uint) bool {
	val := atomic.LoadUint64((*uint64)(bt)) & (1 << n)
	return val != 0
}

func (bt *Bits64) SetBits(v uint64) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval | v
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) ClearBits(v uint64) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval &^ v
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) NegBits(v uint64) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval ^ v
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) GetBits(v uint64) bool {
	val := atomic.LoadUint64((*uint64)(bt)) & v
	return val != 0
}

func (bt *Bits64) Set(v Bits64) {
	*bt = v
}

func (bt *Bits64) GetUint64() uint64 {
	return atomic.LoadUint64((*uint64)(bt))
}
