package forwardauth

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	ssopb "github.com/go-bamboo/pkg/api/sys"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func TokenNull(format string, a ...interface{}) error {
	return errors.Unauthorized("TokenNull", fmt.Sprintf(format, a...))
}

func IsTokenNull(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "TokenNull" && se.Code == 401
}

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if info, ok := transport.FromServerContext(ctx); ok {
				token := extractToken(info)
				if len(token) == 0 {
					err = TokenNull("token is null")
					return
				}
				areq := req.(*ssopb.AuthRequest)
				areq.AccessToken = token
				reply, err = handler(ctx, req)
				if err != nil {
					return
				}
				dp, _ := json.Marshal(reply)
				info.ReplyHeader().Set("x-md-global-dp", string(dp))
				return
			}
			return handler(ctx, req)
		}
	}
}

func extractToken(info transport.Transporter) string {
	tr, ok := info.(*http.Transport)
	if ok {
		return extractHttpToken(tr)
	}
	return ""
}

func extractHttpToken(info *http.Transport) string {
	auth := info.RequestHeader().Get("Authorization")
	access := info.RequestHeader().Get("access_token")
	cookie, cookieErr := info.Request().Cookie("Admin-Token")
	prefix := "Bearer "
	token := ""

	if len(auth) > 0 && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else if len(access) > 0 {
		token = access
	} else if cookieErr == nil {
		token = cookie.Value
	} else if cookieErr != nil {
		log.Error(cookieErr)
	}
	return token
}
