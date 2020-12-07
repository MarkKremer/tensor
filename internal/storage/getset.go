// Code generated by genlib2. DO NOT EDIT.

package storage

import "unsafe"

/* bool */

func (h *Header) Bools() []bool {
	l2 := len(h.Raw) / int(bType.Size())
	return (*(*[]bool)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetB(i int, x bool) { h.Bools()[i] = x }
func (h *Header) GetB(i int) bool    { return h.Bools()[i] }

/* int */

func (h *Header) Ints() []int {
	l2 := len(h.Raw) / int(iType.Size())
	return (*(*[]int)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetI(i int, x int) { h.Ints()[i] = x }
func (h *Header) GetI(i int) int    { return h.Ints()[i] }

/* int8 */

func (h *Header) Int8s() []int8 {
	l2 := len(h.Raw) / int(i8Type.Size())
	return (*(*[]int8)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetI8(i int, x int8) { h.Int8s()[i] = x }
func (h *Header) GetI8(i int) int8    { return h.Int8s()[i] }

/* int16 */

func (h *Header) Int16s() []int16 {
	l2 := len(h.Raw) / int(i16Type.Size())
	return (*(*[]int16)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetI16(i int, x int16) { h.Int16s()[i] = x }
func (h *Header) GetI16(i int) int16    { return h.Int16s()[i] }

/* int32 */

func (h *Header) Int32s() []int32 {
	l2 := len(h.Raw) / int(i32Type.Size())
	return (*(*[]int32)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetI32(i int, x int32) { h.Int32s()[i] = x }
func (h *Header) GetI32(i int) int32    { return h.Int32s()[i] }

/* int64 */

func (h *Header) Int64s() []int64 {
	l2 := len(h.Raw) / int(i64Type.Size())
	return (*(*[]int64)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetI64(i int, x int64) { h.Int64s()[i] = x }
func (h *Header) GetI64(i int) int64    { return h.Int64s()[i] }

/* uint */

func (h *Header) Uints() []uint {
	l2 := len(h.Raw) / int(uType.Size())
	return (*(*[]uint)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetU(i int, x uint) { h.Uints()[i] = x }
func (h *Header) GetU(i int) uint    { return h.Uints()[i] }

/* uint8 */

func (h *Header) Uint8s() []uint8 {
	l2 := len(h.Raw) / int(u8Type.Size())
	return (*(*[]uint8)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetU8(i int, x uint8) { h.Uint8s()[i] = x }
func (h *Header) GetU8(i int) uint8    { return h.Uint8s()[i] }

/* uint16 */

func (h *Header) Uint16s() []uint16 {
	l2 := len(h.Raw) / int(u16Type.Size())
	return (*(*[]uint16)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetU16(i int, x uint16) { h.Uint16s()[i] = x }
func (h *Header) GetU16(i int) uint16    { return h.Uint16s()[i] }

/* uint32 */

func (h *Header) Uint32s() []uint32 {
	l2 := len(h.Raw) / int(u32Type.Size())
	return (*(*[]uint32)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetU32(i int, x uint32) { h.Uint32s()[i] = x }
func (h *Header) GetU32(i int) uint32    { return h.Uint32s()[i] }

/* uint64 */

func (h *Header) Uint64s() []uint64 {
	l2 := len(h.Raw) / int(u64Type.Size())
	return (*(*[]uint64)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetU64(i int, x uint64) { h.Uint64s()[i] = x }
func (h *Header) GetU64(i int) uint64    { return h.Uint64s()[i] }

/* uintptr */

func (h *Header) Uintptrs() []uintptr {
	l2 := len(h.Raw) / int(uintptrType.Size())
	return (*(*[]uintptr)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetUintptr(i int, x uintptr) { h.Uintptrs()[i] = x }
func (h *Header) GetUintptr(i int) uintptr    { return h.Uintptrs()[i] }

/* float32 */

func (h *Header) Float32s() []float32 {
	l2 := len(h.Raw) / int(f32Type.Size())
	return (*(*[]float32)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetF32(i int, x float32) { h.Float32s()[i] = x }
func (h *Header) GetF32(i int) float32    { return h.Float32s()[i] }

/* float64 */

func (h *Header) Float64s() []float64 {
	l2 := len(h.Raw) / int(f64Type.Size())
	return (*(*[]float64)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetF64(i int, x float64) { h.Float64s()[i] = x }
func (h *Header) GetF64(i int) float64    { return h.Float64s()[i] }

/* complex64 */

func (h *Header) Complex64s() []complex64 {
	l2 := len(h.Raw) / int(c64Type.Size())
	return (*(*[]complex64)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetC64(i int, x complex64) { h.Complex64s()[i] = x }
func (h *Header) GetC64(i int) complex64    { return h.Complex64s()[i] }

/* complex128 */

func (h *Header) Complex128s() []complex128 {
	l2 := len(h.Raw) / int(c128Type.Size())
	return (*(*[]complex128)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetC128(i int, x complex128) { h.Complex128s()[i] = x }
func (h *Header) GetC128(i int) complex128    { return h.Complex128s()[i] }

/* string */

func (h *Header) Strings() []string {
	l2 := len(h.Raw) / int(strType.Size())
	return (*(*[]string)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetStr(i int, x string) { h.Strings()[i] = x }
func (h *Header) GetStr(i int) string    { return h.Strings()[i] }

/* unsafe.Pointer */

func (h *Header) UnsafePointers() []unsafe.Pointer {
	l2 := len(h.Raw) / int(unsafePointerType.Size())
	return (*(*[]unsafe.Pointer)(unsafe.Pointer(&h.Raw)))[:l2:l2]
}
func (h *Header) SetUnsafePointer(i int, x unsafe.Pointer) { h.UnsafePointers()[i] = x }
func (h *Header) GetUnsafePointer(i int) unsafe.Pointer    { return h.UnsafePointers()[i] }
