package strx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unsafe"

	json "github.com/bytedance/sonic"
)

//比[]byte() 性能提升100倍

func B2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

//
//func S2b(s string) (b []byte) {
//	/* #nosec G103 */
//	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
//	/* #nosec G103 */
//	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
//	bh.Data = sh.Data
//	bh.Cap = sh.Len
//	bh.Len = sh.Len
//	return b
//}

func S2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnderscoreToUpperCamelCase 下划线单词转为大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func Str2Number[T number](in string) T {
	if in == "" || in == "0" {
		return 0
	}

	return toNumber[T](in)
}

func Str2Numbers[T number](in ...string) []T {
	var out []T
	if len(in) == 0 {
		return out
	}

	for _, i := range in {
		if i == "" {
			continue
		}

		if i == "0" {
			out = append(out, 0)
			continue
		}
		out = append(out, toNumber[T](i))
	}

	return out
}

func toNumber[T number](in string) T {

	a, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0
	}

	return T(a)
}

func S2Map(in string) map[string]interface{} {
	var out = make(map[string]interface{})
	_ = json.Unmarshal(S2b(in), &out)
	return out
}
func B2Map(in []byte) map[string]interface{} {
	var out = make(map[string]interface{})
	_ = json.Unmarshal(in, &out)
	return out
}

func Any2Map(in any) map[string]interface{} {
	var out = make(map[string]interface{})
	_ = json.Unmarshal(Any2Bytes(in), &out)
	return out
}

func Any2Str(v any) string {
	if v == nil {
		return ""
	}
	// 尝试进行类型断言
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return fmt.Sprintf("%f", v)
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return strconv.FormatBool(v)
	default:
		// 对其他类型尝试转换为字符串
		marshal, _ := json.Marshal(v)
		return B2s(marshal)
	}
}

func Any2Bytes(in any) []byte {
	marshal, _ := json.Marshal(in)
	return marshal
}

func MaskStr(s string) string {
	if len(s) <= 3 {
		return s
	}

	firstTwo := s[:1]
	lastTwo := s[len(s)-3:]
	masked := firstTwo + "********" + lastTwo
	return masked
}

// ReplacePlaceholders 替换字符
func ReplacePlaceholders(matchChar, text string, values map[string]string) string {
	re := regexp.MustCompile(matchChar)
	result := re.ReplaceAllStringFunc(text, func(m string) string {
		key := re.FindStringSubmatch(m)[1]
		if val, ok := values[key]; ok {
			return val
		}

		return m // 如果没有匹配，就保留原样
	})

	return result
}
