package main

import (
	"reflect"
	"fmt"
)

type myInt int64
//
func reflectType(x interface{}) {
	//在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v \n", v)
}

//在反射中关于类型还划分为两种：类型（Type）和种类（Kind）
//因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，
//就会用到种类（Kind）。 举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。
func reflectTypeKind(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v kind : %v \n", v.Name(), v.Kind())
}

func reflectValueOf(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("value : %v \n", v)
}

//通过反射获取值 然后强转
func reflectValueReturn(x interface{}) int64 {
/*	v := reflect.TypeOf(x)
	fmt.Printf("fdfdfdsd: %v \n", v)
*/
	g := reflect.ValueOf(x)
	//获取值
	fmt.Printf("fdfdfd: %v --- %T \n", g, g)
	//强转
 	return int64(g.Int())
}

//结构体
type Person struct {
	name string
	age int
}

func main() {
	var a float32 = 3.4
	b := 3.5
	var c int64 = 100

	//打印類型
	reflectType(a)
	reflectType(b)
	reflectType(c)

	//打印值
	reflectValueOf(a)
	reflectValueOf(b)
	reflectValueOf(c)

	//獲取kind類型
	var e *float32
	var f myInt
	var g rune

	reflectTypeKind(e)
	reflectTypeKind(f)
	reflectTypeKind(g)

	var p = Person{
		name : "test",
		age : 1,
	}

	reflectTypeKind(p)

	//通过反射来获取值
 	var test int32 = 323232
	t := reflectValueReturn(test)
	fmt.Printf("日日日%v --- %T \n", t, t)

	//通过反射设置变量的值

	
}