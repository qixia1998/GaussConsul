package handler

import (
	"GaussConsul/pkg/result"
	"net/http"

	"GaussConsul/service/user/api/internal/logic"
	"GaussConsul/service/user/api/internal/svc"
	"GaussConsul/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
