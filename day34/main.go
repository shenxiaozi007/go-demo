package main

import "fmt"

func f1() (r int) {
    defer func() {
        r++
    }()
    return 0
}

func f2() (r int)  {
    t := 5
    defer func() {
        t++
    }()
    return t
}

func f3() (r int) {
    defer func(r int) {
       r = r + 5
    }(r)
    return 1
}


func main1() {
    fmt.Println(f1()) //1
    fmt.Println(f2()) //5
    fmt.Println(f3()) //1
}

type test struct {
    name string
}

func (t *test) close() {
    fmt.Println(t.name, " closed")
}
func main()  {
    ts := []test{{"a"}, {"b"}, {"c"}}
    for _, t := range ts {
        defer t.close()  // 会输出 三个c
    }
}