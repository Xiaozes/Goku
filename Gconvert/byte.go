package Gconvert

import (
	"reflect"
	"unsafe"
)

// Byte 转str
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// str 转 byte
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
