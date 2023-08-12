package g1

import (
	"net/http"

	"github.com/xemxx/go-zero-template/internal/logic/g1"
	"github.com/xemxx/go-zero-template/internal/svc"
	"github.com/xemxx/go-zero-template/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FormExampleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FormExampleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := g1.NewFormExampleLogic(r.Context(), svcCtx)
		err := l.FormExample(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
