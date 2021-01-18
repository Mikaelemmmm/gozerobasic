package handler

import (
	"gozerobasic/lib/xhttp"
	"net/http"

	"gozerobasic/app/gateway/api/v1/usercenter/internal/logic"
	"gozerobasic/app/gateway/api/v1/usercenter/internal/svc"
	"gozerobasic/app/gateway/api/v1/usercenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.ParamErrorResult(r,w,err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		xhttp.HttpResult(r,w,resp,err)
	}
}
