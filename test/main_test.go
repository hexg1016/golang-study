package main

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_func_print(t *testing.T) {
	patches := gomonkey.ApplyFunc(Print, func(args ...string) string {
		return "test"
	})
	defer patches.Reset()
	str1 := Print("a", "b")
	str2 := Print("c")
	assert.Equal(t, str1, str2)
}

func Test_func_print2(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(Print, "test")
	defer patches.Reset()
	str1 := Print("a", "b")
	str2 := Print("c")
	fmt.Println(str1)
	assert.Equal(t, str1, str2)

}

func Test_method(t *testing.T) {
	s := &Server{addr: "127.0.0.1:8080"}
	patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Addr", func(s *Server) string {
		return "s.addr"
	})
	defer patches.Reset()
	str1 := s.Addr()
	fmt.Println(str1)
	assert.Equal(t, str1, "s.addr")
}
