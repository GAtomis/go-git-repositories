package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
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

				userClaims, err := helper.AnalyseToken(authorization)
				if err != nil {
					return nil, err

				}
				fmt.Println(userClaims.Identity)
				if userClaims.Identity == "" {

					return nil, errors.New("no userClaims.Identity")

				}
				ctx = metadata.NewServerContext(ctx, metadata.New(map[string][]string{
					"username": []string{
						userClaims.Name,
					},
					"identity": []string{
						userClaims.Identity,
					},
				}))

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}

func Newa(mds ...map[string][]string) {

}
