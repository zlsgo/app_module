package account

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/sohaha/zlsgo"
)

func TestGenerateSecurePassword(t *testing.T) {
	tt := zlsgo.NewTest(t)

	strict := PasswordConfig{
		MinLength:        8,
		MaxLength:        64,
		RequireUppercase: true,
		RequireLowercase: true,
		RequireNumbers:   true,
		RequireSpecial:   true,
	}

	for _, length := range []int{8, 12, 24, 64} {
		pwd, err := GenerateSecurePassword(length, true)
		tt.NoError(err)
		tt.Equal(length, utf8.RuneCountInString(pwd))

		if ok, msg := ValidatePassword(pwd, strict); !ok {
			t.Fatalf("generated password %q rejected: %s", pwd, msg)
		}
	}

	// 小于最小长度时回退到默认长度，且不生成特殊字符
	pwd, err := GenerateSecurePassword(3, false)
	tt.NoError(err)
	tt.Equal(defaultPasswordLength, utf8.RuneCountInString(pwd))
	for _, r := range pwd {
		if strings.ContainsRune(passwordSpecial, r) {
			t.Fatalf("password %q should not contain special character %q", pwd, r)
		}
	}
}

func TestValidatePasswordUsesRuneLength(t *testing.T) {
	tt := zlsgo.NewTest(t)

	config := PasswordConfig{
		MinLength:        8,
		MaxLength:        12,
		RequireUppercase: false,
		RequireLowercase: false,
		RequireNumbers:   false,
		RequireSpecial:   false,
	}

	// 4 个汉字只有 4 个字符（12 字节），应按字符数判定为过短而不是超长
	ok, msg := ValidatePassword("密码密码", config)
	tt.Equal(false, ok)
	tt.Equal("密码长度不能少于8位", msg)

	// 12 个汉字正好等于上限（36 字节），不应被判为超长
	ok, msg = ValidatePassword("密码密码密码密码密码密码", config)
	tt.Equal(true, ok)
	tt.Equal("", msg)

	ok, msg = ValidatePassword("密码密码密码密码密码密码密", config)
	tt.Equal(false, ok)
	tt.Equal("密码长度不能超过12位", msg)
}

func TestValidatePasswordRequiresCharacterClasses(t *testing.T) {
	tt := zlsgo.NewTest(t)

	config := PasswordConfig{
		MinLength:        8,
		MaxLength:        32,
		RequireUppercase: true,
		RequireLowercase: true,
		RequireNumbers:   true,
		RequireSpecial:   true,
	}

	cases := map[string]string{
		"abcdefg1!": "密码必须包含至少一个大写字母",
		"ABCDEFG1!": "密码必须包含至少一个小写字母",
		"Abcdefgh!": "密码必须包含至少一个数字",
		"Abcdefg1":  "密码必须包含至少一个特殊字符",
	}

	for pwd, want := range cases {
		ok, msg := ValidatePassword(pwd, config)
		if ok {
			t.Fatalf("password %q should be rejected", pwd)
		}
		tt.Equal(want, msg)
	}

	ok, msg := ValidatePassword("Abcdefg1!", config)
	if !ok {
		t.Fatalf("password should be accepted: %s", msg)
	}
}
