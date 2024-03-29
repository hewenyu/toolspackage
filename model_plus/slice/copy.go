package slice

/*
see https://github.com/golang/go/wiki/SliceTricks
append(s[:0:0], s...) 表示在一个空切片基础上插入s
从数组或者切片派生切片（取子切片）有两以下两种方式
s[low : high]       // 双下标形式
s[low : high : max] // 三下标形式
双下标实际是三下标的简写
s[low : high] ==>s[low : high:cap(s)]
以下切片等价
s[0 : len(s)]
s[: len(s)]
s[0 :]
s[:]
s[0 : len(s) : cap(s)]
s[: len(s) : cap(s)]
*/



// Copy 安全的复制切片
func Copy(s []interface{}) []interface{} {
	return append(s[:0:0], s...)
}

// CopyString 安全的复制切片
func CopyString(s []string) []string {
	return append(s[:0:0], s...)
}

// CopyInt 安全的复制切片
func CopyInt(s []int) []int {
	return append(s[:0:0], s...)
}

// CopyInt8 安全的复制切片
func CopyInt8(s []int8) []int8 {
	return append(s[:0:0], s...)
}

// CopyInt16 安全的复制切片
func CopyInt16(s []int16) []int16 {
	return append(s[:0:0], s...)
}

// CopyInt32 安全的复制切片
func CopyInt32(s []int32) []int32 {
	return append(s[:0:0], s...)
}

// CopyInt64 安全的复制切片
func CopyInt64(s []int64) []int64 {
	return append(s[:0:0], s...)
}

// CopyUint 安全的复制切片
func CopyUint(s []uint) []uint {
	return append(s[:0:0], s...)
}

// CopyUint8 安全的复制切片
func CopyUint8(s []uint8) []uint8 {
	return append(s[:0:0], s...)
}

// CopyUint16 安全的复制切片
func CopyUint16(s []uint16) []uint16 {
	return append(s[:0:0], s...)
}

// CopyUint32 安全的复制切片
func CopyUint32(s []uint32) []uint32 {
	return append(s[:0:0], s...)
}

// CopyUint64 安全的复制切片
func CopyUint64(s []uint64) []uint64 {
	return append(s[:0:0], s...)
}

// CopyUintptr 安全的复制切片
func CopyUintptr(s []uintptr) []uintptr {
	return append(s[:0:0], s...)
}

// CopyFloat32 安全的复制切片
func CopyFloat32(s []float32) []float32 {
	return append(s[:0:0], s...)
}

// CopyFloat64 安全的复制切片
func CopyFloat64(s []float64) []float64 {
	return append(s[:0:0], s...)
}

// CopyRune 安全的复制切片
func CopyRune(s []rune) []rune {
	return append(s[:0:0], s...)
}

// CopyByte 安全的复制切片
func CopyByte(s []byte) []byte {
	return append(s[:0:0], s...)
}

// CopyComplex64 安全的复制切片
func CopyComplex64(s []complex64) []complex64 {
	return append(s[:0:0], s...)
}

// CopyComplex128 安全的复制切片
func CopyComplex128(s []complex128) []complex128 {
	return append(s[:0:0], s...)
}

// CopyBool 安全的复制切片
func CopyBool(s []bool) []bool {
	return append(s[:0:0], s...)
}
