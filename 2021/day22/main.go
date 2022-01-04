package main

import (
	"encoding/json"
	"fmt"
)

type JsonTest struct {
	Name string
	email string
}
type Weapon int

const (
	test1 Weapon = iota
	test2
	test3
)


func main1() {
	j := JsonTest{
		Name: "test",
		email: "173315279",
	}
	fmt.Printf("转换前的内容= %+v \n", j)
	jsonInfo, _ := json.Marshal(j)
	fmt.Printf("%+v", string(jsonInfo))
}

func main()  {
	fmt.Printf("%d", test1)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
}

func (c Weapon) String() string {
	switch c {
	case test1:
		return "test1"
	case test2:
		return "test11"
	case test3:
		return "test111"
	}
	return "N/A"
}

