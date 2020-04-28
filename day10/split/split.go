package split

import (
	"strings"
	"fmt"
)

//切割函数
func Split(s, sep string) (result []string) {

	i := strings.Index(s, sep)
	fmt.Println(i)
	for i > -1 {
		result = append(result, s[:i])

		s = s[i+len(sep):]

		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
