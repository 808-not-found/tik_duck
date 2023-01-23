package user

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
	// "github.com/808-not-found/tik_duck/cmd/user/dal/db"
	// user "github.com/808-not-found/tik_duck/kitex_gen/user".
)

var ErrUsernameInfo = errors.New("用户名不合法： 用户名应6-32字符 并且 用户名只能由字母、数字和下划线组成，且必须以字母开头")
var ErrPasswordInfo = errors.New("密码不合法： 密码应6-32字符 且 只能由字母、数字和下划线以及'.' ")
var ErrIDInfo = errors.New("错误的用户id号")

func ErrUserRegisterRequestUsername() error {
	return fmt.Errorf("Err_UserRegisterRequest %w", ErrUsernameInfo)
}
func ErrUserRegisterRequestPassword() error {
	return fmt.Errorf("Err_UserRegisterRequest %w", ErrPasswordInfo)
}
func ErrUserLoginRequestUsername() error {
	return fmt.Errorf("Err_UserLoginRequest %w", ErrUsernameInfo)
}
func ErrUserLoginRequestPassword() error {
	return fmt.Errorf("Err_UserLoginRequest %w", ErrPasswordInfo)
}
func ErrUserRequestID() error {
	return fmt.Errorf("Err_UserLoginRequest %w", ErrIDInfo)
}

// unused protection.
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

// 合法return nil
// 不合法return 错误信息
// 只能由字母、数字和下划线组成，且必须以字母开头。
func RALCheckUsername(s string) bool { // 登陆和注册校验用户名
	for i := 0; i < len(s); i++ {
		var flag int
		switch {
		case s[i] >= 'a' && s[i] <= 'z':
			flag = 1
		case s[i] >= 'A' && s[i] <= 'Z':
			flag = 1
		case s[i] >= '0' && s[i] <= '9':
			flag = 1
		case s[i] == '_':
			flag = 1
		}
		if flag == 0 {
			return false
		}
	}
	return true
}
func RALCheckPassword(s string) bool { // 登陆和注册校验密码
	for i := 0; i < len(s); i++ {
		var flag int
		switch {
		case s[i] >= 'a' && s[i] <= 'z':
			flag = 1
		case s[i] >= 'A' && s[i] <= 'Z':
			flag = 1
		case s[i] >= '0' && s[i] <= '9':
			flag = 1
		case s[i] == '_':
			flag = 1
		case s[i] == '.':
			flag = 1
		}
		if flag == 0 {
			return false
		}
	}
	return true
}
func (p *UserRegisterRequest) IsValid() error { // 注册信息校验
	// 1: string Username //注册用户名，最长32个字符1525
	// 2: string Password //密码，最长32个字符
	if len(p.Username) < 6 || len(p.Username) > 32 {
		return ErrUserRegisterRequestUsername()
	}
	if !RALCheckUsername(p.Username) {
		return ErrUserRegisterRequestUsername()
	}
	if !(p.Username[0] >= 'a' && p.Username[0] <= 'z' || p.Username[0] >= 'A' && p.Username[0] <= 'Z') {
		return ErrUserRegisterRequestUsername()
	}

	if len(p.Password) < 6 || len(p.Password) > 32 {
		return ErrUserRegisterRequestPassword()
	}
	if !RALCheckPassword(p.Password) {
		return ErrUserRegisterRequestPassword()
	}

	return nil
}
func (p *UserLoginRequest) IsValid() error { // 登录请求校验
	if len(p.Username) < 6 || len(p.Username) > 32 {
		return ErrUserLoginRequestUsername()
	}
	if !RALCheckUsername(p.Username) {
		return ErrUserLoginRequestUsername()
	}
	if !(p.Username[0] >= 'a' && p.Username[0] <= 'z' || p.Username[0] >= 'A' && p.Username[0] <= 'Z') {
		return ErrUserLoginRequestUsername()
	}

	if len(p.Password) < 6 || len(p.Password) > 32 {
		return ErrUserLoginRequestPassword()
	}
	if !RALCheckPassword(p.Password) {
		return ErrUserLoginRequestPassword()
	}
	return nil
}

func (p *UserRequest) IsValid() error { // 用户信息请求
	// var My_Token string = p.Token //申请者的token
	// var My_Claims jwt.MyClaims = jwt.ParseToken(My_Token)
	// var My_Username string = My_Claims.Username

	// var Aim_Id int64 = p.UserId
	// contextnil := context.Context{}
	// _, err := db.MGetUsers(contextnil, Aim_Id)
	// if err != nil {
	// 	return fmt.Errorf("错误的用户id号")
	// }
	AimID := p.UserId
	if AimID < 1 {
		return ErrUserRequestID()
	}
	return nil
}
