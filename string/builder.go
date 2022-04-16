package string

import (
	"fmt"
	"unsafe"
)

type Builder struct {
	buf     []byte
	start   int
}

func NewBuilder() Builder{
	return Builder{}
}

// Reset resets the Builder to be empty.
func (b *Builder) Reset() {
	b.start = len(b.buf)
}

// Len returns the number of accumulated bytes.
func (b *Builder) Len() int {
	return len(b.buf) - b.start
}

func (b *Builder) needToGrow(n int) bool {
	return b.start < n
}

func (b *Builder) grow(n int) {
	newLen := b.Len() + n
	var newCap int
	if len(b.buf) == 0 {
		newCap = 64 // arbitrary
	} else {
		newCap = 2 * len(b.buf)
	}
	for newCap < newLen {
		newCap *= 2
		if newCap == 0 {
			panic(fmt.Sprintf("required length (%d) causes buffer size to overflow", newLen))
		}
	}
	newBuf := make([]byte, newCap)
	copy(newBuf[newCap-b.Len():], b.buf[b.start:])
	b.start += newCap - len(b.buf)
	b.buf = newBuf
}

// PrependString prepends the given string to b's buffer.
func (b *Builder) PrependString(str string) {
	if b.needToGrow(len(str)) {
		b.grow(len(str))
	}
	b.start -= len(str)
	copy(b.buf[b.start:], str)
}

// PrependByte prepends the given byte to b's buffer.
func (b *Builder) PrependByte(c byte) {
	if b.needToGrow(1) {
		b.grow(1)
	}
	b.start--
	b.buf[b.start] = c
}

// AppendString appends the given string to b's buffer.
func (b *Builder) AppendString(str string) {
	if b.needToGrow(len(str)) {
		b.grow(len(str))
	}
	oldStart := b.start
	b.start -= len(str)
	copy(b.buf[b.start:], b.buf[oldStart:])
	copy(b.buf[len(b.buf)-len(str):], str)
}

// AppendString appends the given string to b's buffer.
func (b *Builder) AppendStringIf(condition bool, str string) {
	if condition {
		b.AppendString(str)
	}
}

// String returns the accumulated string. No other methods should be called
// after String.
func (b *Builder) String() string {
	bytes := b.buf[b.start:]
	return *(*string)(unsafe.Pointer(&bytes))
}
