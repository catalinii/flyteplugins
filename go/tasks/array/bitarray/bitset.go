/*
 * Copyright (c) 2018 Lyft. All rights reserved.
 */

package bitarray

import (
	"unsafe"
)

/* #nosec */
const blockSize = uint(unsafe.Sizeof(Block(0))) * 8

type Block uint32

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []Block

// Ensures that the given bit is set in the BitSet.
func (s *BitSet) Set(i uint) {
	if len(*s) < int(i/blockSize+1) {
		*s = append(*s, make([]Block, i/blockSize+1)...)
	}

	(*s)[i/blockSize] |= 1 << (i % blockSize)
}

// Ensures that the given bit is cleared (unset) in the BitSet.
func (s *BitSet) Clear(i uint) {
	if len(*s) >= int(i/blockSize+1) {
		(*s)[i/blockSize] &^= 1 << (i % blockSize)
	}
}

// Returns true if the given bit is set, false if it is cleared.
func (s *BitSet) IsSet(i uint) bool {
	if len(*s) < int(i/blockSize+1) {
		return false
	}

	return (*s)[i/blockSize]&(1<<(i%blockSize)) != 0
}

// Returns the length of the BitSet.
func (s *BitSet) Len() int {
	return len(*s)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (s BitSet) DeepCopyInto(out *BitSet) {
	in := &s
	*out = make(BitSet, len(*in))
	copy(*out, *in)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BitSet.
func (s BitSet) DeepCopy() BitSet {
	if s == nil {
		return nil
	}

	out := new(BitSet)
	s.DeepCopyInto(out)
	return *out
}

// Initializes a new BitSet of the specified size.
func NewBitSet(size uint) *BitSet {
	a := make(BitSet, size/blockSize+1)
	return &a
}
