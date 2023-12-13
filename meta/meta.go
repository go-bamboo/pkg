package meta

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-bamboo/pkg/jsonx"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/mssola/user_agent"
)

const (
	KeyToken  = "x-md-global-token"
	KeyDP     = "x-md-global-dp"
	KeyUA     = "x-md-global-ua"
	KeyRealIP = "x-md-global-real-ip"
	KeyLocale = "x-md-global-locale"
	KeyOpID   = "x-md-global-op-id"
)

func GetDataPermissions(ctx context.Context) (permission *DataPermission, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ErrorMdNotFound("不存在md")
		return
	}
	v := md.Get(KeyDP)
	if len(v) <= 0 {
		err = ErrorDpNotFound("不存在dp")
		return
	}
	var dp DataPermission
	if err = jsonx.Unmarshal([]byte(v), &dp); err != nil {
		return
	}
	permission = &dp
	return
}

func GetToken(ctx context.Context) (token string, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ErrorMdNotFound("不存在md")
		return
	}
	v := md.Get(KeyToken)
	if len(v) <= 0 {
		err = ErrorTokenNotFound("不存在token")
		return
	}
	token = v
	return
}

func GetUA(ctx context.Context) (ua *user_agent.UserAgent, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ErrorMdNotFound("不存在md")
		return
	}
	v := md.Get(KeyUA)
	if len(v) <= 0 {
		err = ErrorUaNotFound("ua")
		return
	}
	return user_agent.New(v), nil
}

func GetRealIP(ctx context.Context) (ip string, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return "", ErrorMdNotFound("不存在md")
	}
	v := md.Get(KeyRealIP)
	if len(v) <= 0 {
		err = ErrorRemoteAddrNotFound("remote ip")
		return
	}
	return v, nil
}

func GetLocale(ctx context.Context) (ip string, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return "", ErrorMdNotFound("不存在md")
	}
	v := md.Get(KeyLocale)
	if len(v) <= 0 {
		err = ErrorRemoteAddrNotFound("locale")
		return
	}
	return v, nil
}

func GetOpID(ctx context.Context) (id int64, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0, ErrorMdNotFound("不存在md")
	}
	v := md.Get(KeyOpID)
	if len(v) <= 0 {
		err = ErrorRemoteAddrNotFound("locale")
		return
	}
	id, _ = strconv.ParseInt(v, 10, 64)
	return id, nil
}

func SetOpID(ctx context.Context, id int64) (err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return ErrorMdNotFound("不存在md")
	}
	md.Set(KeyOpID, fmt.Sprint(id))
	return nil
}
