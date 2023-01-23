package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	db "github.com/808-not-found/tik_duck/cmd/user/db"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
)

// unused protection
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
func (p *user.UserRegisterRequest) IsValid() error { //注册信息校验
	// 1: string Username //注册用户名，最长32个字符1525
	// 2: string Password //密码，最长32个字符
	if len(p.Username) < 6 || len(p.Username) > 32 {
		return fmt.Errorf("用户名的长度不合法： 用户名应6-32字符 并且 用户名只能由字母、数字和下划线组成，且必须以字母开头 当前长度: %d", len(p.Username))
	}
	for i := 0; i < len(p.Username); i++ {
		var flag int = 0
		if p.Username[i] >= 'a' && p.Username[i] <= 'z' {
			flag = 1
		}
		if p.Username[i] >= 'A' && p.Username[i] <= 'Z' {
			flag = 1
		}
		if p.Username[i] >= '0' && p.Username[i] <= '9' {
			flag = 1
		}
		if p.Username[i] == '_' {
			flag = 1
		}
		if flag == 0 || (!(p.Username[i] >= 'a' && p.Username[i] <= 'z' || p.Username[i] >= 'A' && p.Username[i] <= 'Z')) {
			return fmt.Errorf("用户名只能由字母、数字和下划线组成，且必须以字母开头")
		}
	}
	if len(p.Password) < 6 || len(p.Password) > 32 {
		return fmt.Errorf("密码的长度不合法： 密码应6-32字符 且 只能由字母、数字和下划线以及'.' 当前长度: %d", len(p.Password))
	}
	for i := 0; i < len(p.Password); i++ {
		var flag int = 0
		if p.Password[i] >= 'a' && p.Password[i] <= 'z' {
			flag = 1
		}
		if p.Password[i] >= 'A' && p.Password[i] <= 'Z' {
			flag = 1
		}
		if p.Password[i] >= '0' && p.Password[i] <= '9' {
			flag = 1
		}
		if p.Password[i] == '_' || p.Password[i] == '.' {
			flag = 1
		}
		if flag == 0 || (!(p.Username[i] >= 'a' && p.Username[i] <= 'z' || p.Username[i] >= 'A' && p.Username[i] <= 'Z')) {
			return fmt.Errorf("密码只能由字母、数字和下划线以及'.'组成")
		}
	}
	return nil
}
func (p *user.UserLoginRequest) IsValid() error { //登录请求校验
	if len(p.Username) < 6 || len(p.Username) > 32 {
		return fmt.Errorf("用户名的长度不合法： 用户名应6-32字符 并且 用户名只能由字母、数字和下划线组成，且必须以字母开头 当前长度: %d", len(p.Username))
	}
	for i := 0; i < len(p.Username); i++ {
		var flag int = 0
		if p.Username[i] >= 'a' && p.Username[i] <= 'z' {
			flag = 1
		}
		if p.Username[i] >= 'A' && p.Username[i] <= 'Z' {
			flag = 1
		}
		if p.Username[i] >= '0' && p.Username[i] <= '9' {
			flag = 1
		}
		if p.Username[i] == '_' {
			flag = 1
		}
		if flag == 0 || (!(p.Username[i] >= 'a' && p.Username[i] <= 'z' || p.Username[i] >= 'A' && p.Username[i] <= 'Z')) {
			return fmt.Errorf("用户名只能由字母、数字和下划线组成，且必须以字母开头")
		}
	}
	if len(p.Password) < 6 || len(p.Password) > 32 {
		return fmt.Errorf("密码的长度不合法： 密码应6-32字符 且 只能由字母、数字和下划线以及'.' 当前长度: %d", len(p.Password))
	}
	for i := 0; i < len(p.Password); i++ {
		var flag int = 0
		if p.Password[i] >= 'a' && p.Password[i] <= 'z' {
			flag = 1
		}
		if p.Password[i] >= 'A' && p.Password[i] <= 'Z' {
			flag = 1
		}
		if p.Password[i] >= '0' && p.Password[i] <= '9' {
			flag = 1
		}
		if p.Password[i] == '_' || p.Password[i] == '.' {
			flag = 1
		}
		if flag == 0 || (!(p.Username[i] >= 'a' && p.Username[i] <= 'z' || p.Username[i] >= 'A' && p.Username[i] <= 'Z')) {
			return fmt.Errorf("密码只能由字母、数字和下划线以及'.'组成")
		}
	}
	return nil
}

func (p *user.UserRequest) IsValid() error { //用户信息请求
	// var My_Token string = p.Token //申请者的token
	// var My_Claims jwt.MyClaims = jwt.ParseToken(My_Token)
	// var My_Username string = My_Claims.Username

	var Aim_Id int64 = p.UserId
	Aim_Info, err := db.MGetUsers(Aim_Id)
	if err != nil {
		return fmt.Errorf("错误的用户id号")
	}
	return nil
}
