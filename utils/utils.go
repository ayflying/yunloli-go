package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"strconv"
	"strings"
	"time"
)

var (
	ctx   = gctx.New()
	Utils *Mod
)

type Number interface {
	int | int64 | int32 | int16 | uint64 | uint32 | uint16 | float32 | float64
}

//type Any interface {
//	interface{} | string | int | int64 | int32 | int16 | uint64 | uint32 | uint16 | float32 | float64
//}

type Mod struct {
}

func (m *Mod) Load() {
	g.Log().Debugf(gctx.New(), "初始化工具类")
	Utils = &Mod{}
	return
}

// 两个时间相隔多少天，需要考虑时区
func (m *Mod) GetDay(t1 *gtime.Time, t2 *gtime.Time) int {
	if t2 == nil {
		t2 = gtime.New(time.UnixMilli(0))
	}
	return int(t1.Sub(t2).Hours() / 24)
}

// 字符串转道具类型
func (m *Mod) Spilt2Item(str string) (result [][]int64) {
	parts := strings.Split(str, "|") // 分割字符串

	for i := 0; i < len(parts)-1; i += 2 {
		num1, _ := strconv.ParseInt(parts[i], 10, 64)
		num2, _ := strconv.ParseInt(parts[i+1], 10, 64)

		pair := []int64{num1, num2}
		result = append(result, pair)
	}
	return
}

// 切片换道具类型
func (m *Mod) Slice2Item(slice []int64) (res [][]int64) {
	res = make([][]int64, 0)
	for i := 0; i < len(slice)-1; i += 2 {
		pair := []int64{slice[i], slice[i+1]}
		res = append(res, pair)
	}
	return res
}

// RemoveSlice 从切片中移除指定的值。
// 参数:
//
//	slice: 待操作的切片。
//	value: 需要移除的值。
//
// 返回值:
//
//	移除指定值后的切片。
//
// 该函数通过遍历切片，从后向前检查每个元素，如果找到与指定值相等的元素，则将其从切片中移除。
// 这种从后向前的遍历方法可以避免因移除元素而导致的数组重新排列带来的额外计算。
// RemoveSlice 删除切片中的某个值
func RemoveSlice[t Number](slice []t, value t) []t {
	// 从后向前遍历切片
	for i := len(slice) - 1; i >= 0; i-- {
		// 检查当前元素是否等于需要移除的值
		if slice[i] == value {
			// 如果相等，移除该元素
			// 使用append和切片操作符来实现移除操作，将i之前和i之后的元素合并到一起
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	// 返回处理后的切片
	return slice
}

// InArray 判断当前切片中是否拥有当前值
// InArray[t Number] 支持的类型
//
//	@Description:
//	@param value 需要查找的值
//	@param array 进行查找的切片
//	@return bool 返回是否存在
func InArray[t Number](value t, array []t) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// ExtractPage 根据给定的页码和每页大小，从项目切片中提取指定页的项目。
// 它支持泛型，可以用于任何类型的切片。
// 参数:
//
//	items: 项目切片，代表所有待分页的项目。
//	page: 指定的页码，起始页码为1。
//	size: 每页的项目数量。
//
// 返回值:
//
//	返回一个切片，包含指定页的项目。
//
// 如果每页大小为0，将默认为1。如果项目切片为空，或指定页的项目数量少于每页大小，且页码大于0，则直接返回整个项目切片。
// ExtractPage [t Any]
//
//	@Description:
//	@param items
//	@param page
//	@param size
//	@return []t
func ExtractPage[t interface{}](items []t, page int, size int) []t {
	// 如果每页大小为0，将其默认设置为1。
	// 如果每页大小为0，将其默认设置为1。
	if size == 0 {
		size = 1
	}
	// 如果项目切片为空，直接返回空切片。
	if len(items) == 0 {
		return []t{}
	}
	//// 如果项目切片长度小于每页大小，并且页码大于0，返回整个项目切片。
	//if len(items) < size && page > 0 {
	//	//return items
	//}

	// 计算起始索引和结束索引。
	// 根据页码和每页大小计算起始索引和结束索引。
	// 计算起始索引和结束索引。
	startIndex := (page - 1) * size
	endIndex := startIndex + size
	// 如果结束索引超出项目切片长度，调整结束索引为项目切片的长度。
	// 如果结束索引超出项目切片的长度，将其调整为项目切片的长度。
	if endIndex >= len(items) { // 确保不会超出切片范围
		endIndex = len(items)
	}

	// 如果起始索引超出项目切片长度，返回空切片。
	if len(items) < startIndex || len(items) < endIndex {
		return []t{}
	}

	// 根据起始索引和结束索引从项目切片中提取指定页的项目，并返回。
	// 返回指定页的项目切片。

	return items[startIndex:endIndex]
}

// 这是一个用于反转切片的函数示例
// reverseSlice[T comparable]
//
//	@Description:
//	@param s
//	@return []T
func ReverseSlice[T comparable](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
