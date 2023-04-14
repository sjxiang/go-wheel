package orm

import (
	"unicode"
)


// 驼峰转下划线
func underscoreName(name string) string {
	var buf []byte

	for i, v := range name {
		// 大写
		if unicode.IsUpper(v) {
			// 排除首字母
			if i != 0 {
				buf = append(buf, '_')
			} 
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}
	}

	return string(buf)
}
