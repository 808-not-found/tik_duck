package allerrors

import (
	"errors"
	"fmt"
)

// 如果存在err可以返回就直接返回err就可以
// 此文件下存放自定义错误以及错误码的定义
// 0   ：一切正常
// 1001: 基础接口-视频流错误
// 1002: 基础接口-创建用户-数据库失败
// 1003: 基础接口-创建用户-鉴权失败
// 1004: 基础接口-登陆用户-数据库用户运行问题的失败
// 1005: 基础接口-登陆用户-登陆用户不存在
// 1006: 基础接口-登陆用户-密码错误
// 1007: 基础接口-登陆用户-鉴权失败
// 1008: 基础接口-获取用户信息-解析鉴权失败
// 1009: 基础接口-获取用户信息-发起者用户不存在
// 1010: 基础接口-获取用户信息-被调查用户不存在
// 1011: 基础接口-获取用户信息-数据库到rpc 转换失败

// 1012: 基础接口-创建用户-用户已经存在

// 1021: base-视频流-鉴权错误
// 1022：base-视频流-数据库查询错误
// 1023：base-视频流-数据封装错误
// 1024：base-视频发布-鉴权错误
// 1025：base-视频发布-数据库写入错误
// 1026：base-视频列表-鉴权错误
// 1027：base-视频列表-数据库查询错误
// 1028：base-视频列表-数据封装错误

// 2031：plat-commentList-鉴权错误
// 2032: plat-commentList-数据库查询错误
// 2033: plat-commentList-数据封装错误
// 2034: plat-commentAction-鉴权错误
// 2035: plat-commentAction-数据库操作错误
// 2036: plat-commentAction-数据封装错误
// 2037: plat-favoriteList-鉴权错误
// 2038: plat-favoriteList-数据库查询错误
// 2039: plat-favoriteList-数据封装错误
// 2040: plat-favoriteAction-鉴权错误
// 2041: plat-favoriteAction-数据库操作错误

// 3001: uu-action-token
// 3002: uu-action-登陆状态
// 3003: uu-action-db-关注
// 3004: uu-action-db-取关
// 3005: uu-关注列表-token
// 3006: uu-关注列表-登录状态
// 3007：uu-关注列表-db
// 3008: uu-关注列表-pack
// 3009: uu-粉丝列表-token
// 3010: uu-粉丝列表-登陆状态
// 3011: uu-粉丝列表-db
// 3012: uu-粉丝列表-pack
// 3013: uu-朋友列表-token
// 3014: uu-朋友列表-登录状态
// 3015: uu-朋友列表-db
// 3016: uu-朋友列表-pack

// jwt错误.
var errJWTParseTokenRun = errors.New("jwt：不合法的token")

func ErrJWTParseTokenRun() error {
	return fmt.Errorf("Err_Token %w", errJWTParseTokenRun)
}

// 基础接口DBUser转换成为RPCUser中查询发生错误.
var errDBUserToRPCUserRun = errors.New("gorm：基础接口pack中DBUser转换成为RPCUser时 查询错误")

func ErrDBUserToRPCUserRun() error {
	return fmt.Errorf("ErrDBUserToRpcUser %w", errDBUserToRPCUserRun)
}

// 基础接口DBUser转换成为RPCUser查询参数为空.
var errDBUserToRPCUserVali = errors.New("gorm：user查询信息时 传入查询对象为空")

func ErrDBUserToRPCUserVali() error {
	return fmt.Errorf("ErrDBUserToRpcUser %w", errDBUserToRPCUserVali)
}

// 基础接口参数校验 注册和登陆的用户名不合法.
var errUserRegisteAndLoginrRequestUsername = errors.New("用户名不合法： 用户名应6-32字符 并且 用户名只能由字母、数字和下划线组成，且必须以字母开头")

func ErrUserRegisterRequestUsername() error {
	return fmt.Errorf("Err_UserRegisterRequest %w", errUserRegisteAndLoginrRequestUsername)
}
func ErrUserLoginRequestUsername() error {
	return fmt.Errorf("Err_UserLoginRequest %w", errUserRegisteAndLoginrRequestUsername)
}

// 基础接口参数校验 注册和登陆的密码不合法.
var errUserRegisterAndLoginRequestPassword = errors.New("密码不合法： 密码应6-32字符 且 只能由字母、数字和下划线以及'.' ")

func ErrUserLoginRequestPassword() error {
	return fmt.Errorf("Err_UserLoginRequest %w", errUserRegisterAndLoginRequestPassword)
}
func ErrUserRegisterRequestPassword() error {
	return fmt.Errorf("Err_UserRegisterRequest %w", errUserRegisterAndLoginRequestPassword)
}

// 基础接口参数校验 获取用户信息id不合法.
var errUserRequestID = errors.New("错误的用户id号")

func ErrUserRequestID() error {
	return fmt.Errorf("Err_UserRequest %w", errUserRequestID)
}

// 测试用 表示非nil错误
// 基础接口参数校验 获取用户信息id不合法.
var errTestnotnil = errors.New("notnil")

func ErrTestnotnil() error {
	return fmt.Errorf("Err_Test %w", errTestnotnil)
}
