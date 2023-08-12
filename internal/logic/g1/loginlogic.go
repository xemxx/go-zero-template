package g1

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"

	"github.com/xemxx/go-zero-template/internal/model"
	"github.com/xemxx/go-zero-template/internal/svc"
	"github.com/xemxx/go-zero-template/internal/types"
	"github.com/xemxx/go-zero-template/pkg/mysql"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	con := mysql.GetConn()
	db := model.NewNoteModel(con)
	_, err = db.FindOne(l.ctx, 1)
	if err != nil {
		return nil, errors.New(int(types.Failed), types.FailedMessage)
	}
	return
}
