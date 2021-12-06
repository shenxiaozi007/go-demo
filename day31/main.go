package main

import "fmt"

func main() {
    sn1 := struct {
        age int
        name string
    }{
        age : 11,
        name: "qq",
    }

    sn2 := struct {
        age int
        name string
    }{
        age: 11,
        name: "qq",
    }
    if sn1 == sn2 {
        fmt.Println("sn1 == sn2")
    } else {
        fmt.Println("sn1 != sn2")
    }

    sm1 := struct {
        age int
        m map[string]string
    }{
        age: 11,
        m: map[string]string{
            "a" : "1",
        },
    }

    sm2 := struct {
        age int
        m map[string]string
    }{
        age: 11,
        m: map[string]string{
            "a": "1",
        },
    }

    //对比结构体,报错,因为只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
    //但是结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
    //这里就属于不可比较的类型,map
    if sm1 == sm2 {
        fmt.Println("sm1==sm2")
    }
}
