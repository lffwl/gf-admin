package helper

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

var Array = array{}

type array struct {
}

// In 检查元素是否在集合中
func (h *array) In(value string, arr []string) bool {

	for _, item := range arr {
		if item == value {
			return true
		}
	}

	return false
}

// MapKeyToValue 反转Map键值
func (h *array) MapKeyToValue(arr g.Map) g.Map {
	n := g.Map{}
	for k, v := range arr {
		n[strings.ToUpper(gconv.String(v))] = k
	}
	return n
}

// MaxSliceInt 获取int切片中的最大值
func (h *array) MaxSliceInt(values []int) int {
	var max int
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}
