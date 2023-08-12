package g1

import (
	"net/http"

	"github.com/xemxx/go-zero-template/internal/logic/g1"
	"github.com/xemxx/go-zero-template/internal/svc"
	"github.com/xemxx/go-zero-template/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PathExampleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathExampleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := g1.NewPathExampleLogic(r.Context(), svcCtx)
		resp, err := l.PathExample(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
