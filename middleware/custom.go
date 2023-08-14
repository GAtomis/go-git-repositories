package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"go-git-repositories/helper"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering

				authorization := tr.RequestHeader().Get("authorization")
				if authorization == "" {
					return nil, errors.New("no Auth")
				}
				token, err := helper.AnalyseToken(authorization)

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
