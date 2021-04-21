package main

import "fmt"

type People interface {
	Speak(stringStr string) string
}

type People2 interface {
	Speak2(stringStr2 string) string
}
type Student struct {

}

func (stu *Student) Speak(think string) (talk string)  {
	if think == "sb" {
		talk = "fuck"
	} else {
		talk = "222"
	}
	return
}

func (stu Student) Speak2(think2 string) (talk string)  {
	if think2 == "sb" {
		talk = "fuck"
	} else {
		talk = "222"
	}
	return
}

func main() {
	var peo People = &Student{}
	var peo1 People2 = Student{}
	test := Student{}
	fmt.Println(test)
	think := "bitch"
	fmt.Println(peo.Speak(think))
	fmt.Println(peo1.Speak2(think))
}
