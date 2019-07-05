package util

import (
	"reflect"
	"testing"
)

func TestBytes2String(t *testing.T) {
	hello := []byte("hello world!")
	expect := Bytes2String(hello)

	if !reflect.DeepEqual(string(hello),expect) {
		t.Fatalf("expect value is %s, got is %s",string(hello),string(expect))
	}
}

func TestString2Bytes(t *testing.T) {
 	s := "hello world!"
	expect := String2Bytes(s)
	if !reflect.DeepEqual([]byte(s),expect) {
		t.Fatalf("expect value is %s, got is %s",s,string(expect))
	}
}

