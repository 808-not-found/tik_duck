package userservice

import (
	"context"

	"github.com/808-not-found/tik_duck/cmd/user/dal/db"
	"github.com/808-not-found/tik_duck/cmd/user/pack"
	"github.com/808-not-found/tik_duck/kitex_gen/user"
	"github.com/808-not-found/tik_duck/pkg/consts"
	"github.com/808-not-found/tik_duck/pkg/jwt"
)

func UserInfoService(ctx context.Context, req *user.UserRequest) (int32, string, *user.User, error) {
	// 创建返回时的变量
	var statusCode int32

	// 取出请求中的两个信息
	reqfromToken := req.Token
	reqAimID := req.UserId

	// 解析发起者用户鉴权 -- 到发起者用户id
	claims, err := jwt.ParseToken(reqfromToken)
	if err != nil {
		statusCode = 1008
		return statusCode, "", nil, err
	}
	reqfromUsername := claims.Username
	dbfromUserinfo, err := db.QueryUser(ctx, reqfromUsername)
	if err != nil {
		statusCode = 1009
		return statusCode, "", nil, err
	}
	reqfromid := dbfromUserinfo.ID

	// 获取db aim数据
	reqAimInfo, err := db.GetUser(ctx, reqAimID)
	if err != nil {
		statusCode = 1010
		return statusCode, "", nil, err
	}
	// 查询关注关系
	userinfo, err := pack.DBUserToRPCUser(&reqAimInfo, reqfromid)
	if err != nil {
		statusCode = 1011
		return statusCode, "", nil, err
	}

	// 成功返回
	statusMsg := consts.Success
	return statusCode, statusMsg, userinfo, nil
}
