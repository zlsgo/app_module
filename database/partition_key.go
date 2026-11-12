package database

import "strconv"

// Key 分区键。空键表示默认分区，非空键只允许 ASCII 字母、数字、下划线和连字符。
// 键不会被裁剪、改写或解释为业务身份。
type Key string

// KeyOf 把整数转成十进制分区键，不赋予整数特殊语义。
// 例如 KeyOf(0) 返回 "0"，不是默认分区的空键。
func KeyOf(value int64) Key {
	return Key(strconv.FormatInt(value, 10))
}

// String 返回键的字符串形式。
func (k Key) String() string {
	return string(k)
}

// Int64 把十进制分区键转回整数，空键或非十进制键返回错误。
//
// 它与 KeyOf 互逆，但不校验是否为规范写法（"01" 也会解析成 1）；
// 需要按键的原始拼写做严格校验时自行比较 FormatInt 结果。
func (k Key) Int64() (int64, error) {
	if k == "" {
		return 0, ErrInvalidKey
	}
	return strconv.ParseInt(string(k), 10, 64)
}

func validPartitionKey(key Key) bool {
	for i := 0; i < len(key); i++ {
		c := key[i]
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_' || c == '-' {
			continue
		}
		return false
	}
	return true
}
