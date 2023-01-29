package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	allerrors "github.com/808-not-found/tik_duck/pkg/allerrors"
)

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
		return allerrors.ErrUserRegisterRequestUsername()
	}
	if !RALCheckUsername(p.Username) {
		return allerrors.ErrUserRegisterRequestUsername()
	}
	if !(p.Username[0] >= 'a' && p.Username[0] <= 'z' || p.Username[0] >= 'A' && p.Username[0] <= 'Z') {
		return allerrors.ErrUserRegisterRequestUsername()
	}

	if len(p.Password) < 6 || len(p.Password) > 32 {
		return allerrors.ErrUserRegisterRequestPassword()
	}
	if !RALCheckPassword(p.Password) {
		return allerrors.ErrUserRegisterRequestPassword()
	}

	return nil
}
func (p *UserLoginRequest) IsValid() error { // 登录请求校验
	if len(p.Username) < 6 || len(p.Username) > 32 {
		return allerrors.ErrUserLoginRequestUsername()
	}
	if !RALCheckUsername(p.Username) {
		return allerrors.ErrUserLoginRequestUsername()
	}
	if !(p.Username[0] >= 'a' && p.Username[0] <= 'z' || p.Username[0] >= 'A' && p.Username[0] <= 'Z') {
		return allerrors.ErrUserLoginRequestUsername()
	}

	if len(p.Password) < 6 || len(p.Password) > 32 {
		return allerrors.ErrUserLoginRequestPassword()
	}
	if !RALCheckPassword(p.Password) {
		return allerrors.ErrUserLoginRequestPassword()
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
		return allerrors.ErrUserRequestID()
	}
	return nil
}

func (p *FeedRequest) IsValid() error {
	return nil
}

func (p *PublishListRequest) IsValid() error {
	return nil
}

func (p *PublishActionRequest) IsValid() error {
	return nil
}
