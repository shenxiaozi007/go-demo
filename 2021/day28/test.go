package day28

import (
	"fmt"
	"reflect"
)

type Server struct {
	ServerName string `key1:"value1" key11:"value11"`
	ServerIp   string `key2:"value2"`
}

func main() {
	s := Server{}
	st := reflect.TypeOf(s)

	field1 := st.Field(0)

	fmt.Printf("key1:%v\n", field1.Tag.Get("key1"))
	fmt.Printf("key11:%v\n", field1.Tag.Get("key11"))

	field2 := st.Field(1)
	fmt.Printf("key2:%v\n", field2.Tag.Get("key2"))
}