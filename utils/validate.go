package utils

import "github.com/itcloudy/validator"

//校验密码至少8位且必须包含数字，大小写字母，特殊字符
func ValidatePassword(fl validator.FieldLevel) bool {
	var (
		number      int
		bigLetter   int
		smallLetter int
		special     int
	)
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	for i := 0; i < len(password); i++ {
		ch := password[i]
		if ch > 47 && ch < 58 {
			number++
		}
		if ch > 64 && ch < 91 {
			bigLetter++
		}
		if ch > 96 && ch < 123 {
			smallLetter++
		}
		if (ch < 48) || (ch > 57 && ch < 65) || (ch > 90 && ch < 97) || (ch > 122 && ch < 128) {
			special++
		}
	}
	if number == 0 || bigLetter == 0 || smallLetter == 0 || special == 0 {
		return false
	}
	return true

}
