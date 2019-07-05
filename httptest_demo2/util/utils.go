package util

import (
	"reflect"
	"unsafe"
)

func String2Bytes(s string) []byte {
	b := (*[]byte)(unsafe.Pointer(&s))
	return *b
}

func Bytes2String(b []byte) string  {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}))
}