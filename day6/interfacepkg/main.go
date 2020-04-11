package interfacepkg

import "fmt"

type dog struct{}
type cat struct{}

func main() {

}

// 实现了say的方法
func (d dog) say() {
	fmt.Println("fuck")
}
