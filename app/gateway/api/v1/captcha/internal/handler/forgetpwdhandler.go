package handler

import (
	"gozerobasic/lib/xhttp"
	"net/http"

	"gozerobasic/app/gateway/api/v1/captcha/internal/logic"
	"gozerobasic/app/gateway/api/v1/captcha/internal/svc"
	"gozerobasic/app/gateway/api/v1/captcha/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ForgetPwdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.ParamErrorResult(r,w, err)
			return
		}

		l := logic.NewForgetPwdLogic(r.Context(), ctx)
		resp, err := l.ForgetPwd(req)
		xhttp.HttpResult(r,w, resp, err)
	}
}
