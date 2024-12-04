package handler

import (
	"GaussConsul/pkg/result"
	"net/http"

	"GaussConsul/service/gauss/api/internal/logic"
	"GaussConsul/service/gauss/api/internal/svc"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHealthCheckLogic(r.Context(), svcCtx)
		resp, err := l.HealthCheck()
		result.HttpResult(r, w, resp, err)
	}
}
