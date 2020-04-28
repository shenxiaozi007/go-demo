package split

import (
	"testing"
	"reflect"
)

//测试
func TestSplit(t *testing.T) {// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")

	want := []string{"a", "b", "c"}

	if !reflect.DeepEqual(want, got){ //因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted: %v, got %v", got, want)   // 测试失败输出错误提示
	}
}

//测试多个字符串
func TestMoreSplit(t *testing.T) {
	got := Split("abcdbcfffbcddd", "bc")
	want := []string{"a", "d", "fff", "ddd"}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted: %v, got %v", got, want)   // 测试失败输出错误提示
	}
}

//测试中文
func TestChineseSplit(t *testing.T) {
	//
}