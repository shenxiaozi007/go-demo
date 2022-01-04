package pkg2

// 首字母小写，外部包不可见，只能在当前包内使用
var a = 100

// 首字母大写外部包可见，可在其他包中使用
const Mode = 1

type person struct {
	name string
}

func Add(x, y int) int {
	return x + y
}

//结构体中的字段名和接口中的方法名如果首字母都是大写，
//外部包可以访问这些字段和方法。例如：
type Student struct {
	Name  string
	class string
}
