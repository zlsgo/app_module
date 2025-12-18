package account

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

const (
	// 密码字符集
	passwordLowercase = "abcdefghijklmnopqrstuvwxyz"
	passwordUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passwordNumbers   = "0123456789"
	passwordSpecial   = "!@#$%^&*"

	// 默认密码长度
	defaultPasswordLength = 12
)

// GenerateSecurePassword 生成安全的随机密码
// length: 密码长度，最小8位
// includeSpecial: 是否包含特殊字符
func GenerateSecurePassword(length int, includeSpecial bool) (string, error) {
	if length < 8 {
		length = defaultPasswordLength
	}

	// 构建字符集
	charset := passwordLowercase + passwordUppercase + passwordNumbers
	if includeSpecial {
		charset += passwordSpecial
	}

	// 确保密码包含各种字符类型
	var password strings.Builder
	password.Grow(length)

	// 至少包含一个小写字母
	if err := appendRandomChar(&password, passwordLowercase); err != nil {
		return "", err
	}

	// 至少包含一个大写字母
	if err := appendRandomChar(&password, passwordUppercase); err != nil {
		return "", err
	}

	// 至少包含一个数字
	if err := appendRandomChar(&password, passwordNumbers); err != nil {
		return "", err
	}

	// 如果需要特殊字符，至少包含一个
	if includeSpecial {
		if err := appendRandomChar(&password, passwordSpecial); err != nil {
			return "", err
		}
	}

	// 填充剩余字符
	remainingLength := length - password.Len()
	for i := 0; i < remainingLength; i++ {
		if err := appendRandomChar(&password, charset); err != nil {
			return "", err
		}
	}

	// 打乱密码顺序
	result := []rune(password.String())
	shuffleRunes(result)

	return string(result), nil
}

// appendRandomChar 从字符集中随机选择一个字符追加到 builder
func appendRandomChar(builder *strings.Builder, charset string) error {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return err
	}
	builder.WriteByte(charset[n.Int64()])
	return nil
}

// shuffleRunes 使用 Fisher-Yates 算法打乱 rune 数组
func shuffleRunes(runes []rune) {
	n := len(runes)
	for i := n - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		jIdx := j.Int64()
		runes[i], runes[jIdx] = runes[jIdx], runes[i]
	}
}

// PasswordConfig 密码配置
type PasswordConfig struct {
	// MinLength 最小长度
	MinLength int
	// MaxLength 最大长度
	MaxLength int
	// RequireUppercase 需要大写字母
	RequireUppercase bool
	// RequireLowercase 需要小写字母
	RequireLowercase bool
	// RequireNumbers 需要数字
	RequireNumbers bool
	// RequireSpecial 需要特殊字符
	RequireSpecial bool
	// ForceChangeOnFirstLogin 首次登录强制修改
	ForceChangeOnFirstLogin bool
}

// DefaultPasswordConfig 默认密码配置
var DefaultPasswordConfig = PasswordConfig{
	MinLength:               8,
	MaxLength:               32,
	RequireUppercase:        true,
	RequireLowercase:        true,
	RequireNumbers:          true,
	RequireSpecial:          false,
	ForceChangeOnFirstLogin: true,
}

// ValidatePassword 验证密码是否符合配置要求
func ValidatePassword(password string, config PasswordConfig) (bool, string) {
	if len(password) < config.MinLength {
		return false, "密码长度不能少于" + strconv.Itoa(config.MinLength) + "位"
	}

	if len(password) > config.MaxLength {
		return false, "密码长度不能超过" + strconv.Itoa(config.MaxLength) + "位"
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, ch := range password {
		switch {
		case 'A' <= ch && ch <= 'Z':
			hasUpper = true
		case 'a' <= ch && ch <= 'z':
			hasLower = true
		case '0' <= ch && ch <= '9':
			hasNumber = true
		case strings.ContainsRune(passwordSpecial, ch):
			hasSpecial = true
		}
	}

	if config.RequireUppercase && !hasUpper {
		return false, "密码必须包含至少一个大写字母"
	}

	if config.RequireLowercase && !hasLower {
		return false, "密码必须包含至少一个小写字母"
	}

	if config.RequireNumbers && !hasNumber {
		return false, "密码必须包含至少一个数字"
	}

	if config.RequireSpecial && !hasSpecial {
		return false, "密码必须包含至少一个特殊字符"
	}

	return true, ""
}
