package middleware

import (
	"context"
	"fmt"

	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/SGNYYYY/gomall/app/frontend/utils"
	rpcauth "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		tokenString, ok := session.Get("token").(string)
		if !ok {
			c.Next(ctx)
			return
		}
		if tokenString == "" {
			c.Next(ctx)
			return
		}
		authResp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &rpcauth.VerifyTokenReq{Token: tokenString})
		if err != nil {
			panic(err)
		}
		if !authResp.Res {
			c.Next(ctx)
			return
		}
		session.Set("token", authResp.Token)
		err = session.Save()
		if err != nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, frontendUtils.UserIdKey, authResp.UserId)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// tokenString := c.Request.Header.Get("Authorization")
		session := sessions.Default(c)
		tokenString, ok := session.Get("token").(string)
		if !ok {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if frontendUtils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		if tokenString == "" {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if frontendUtils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		authResp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &rpcauth.VerifyTokenReq{Token: tokenString})
		if err != nil {
			panic(err)
		}

		if !authResp.Res {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if frontendUtils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, frontendUtils.UserIdKey, authResp.UserId)
		c.Next(ctx)
	}
}
