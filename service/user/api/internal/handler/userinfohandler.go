package handler

import (
	"GaussConsul/pkg/result"
	"net/http"

	"GaussConsul/service/user/api/internal/logic"
	"GaussConsul/service/user/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		result.HttpResult(r, w, resp, err)
	}
}
