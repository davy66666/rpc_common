package logic

import (
	"context"
	g "github.com/doug-martin/goqu/v9"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"ks_api_service/internal/model"

	"ks_api_service/internal/svc"
	"ks_api_service/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLastWithdrawalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLastWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLastWithdrawalLogic {
	return &GetLastWithdrawalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLastWithdrawal
func (l *GetLastWithdrawalLogic) GetLastWithdrawal(in *api.GetLastWithdrawalReq) (*api.GetLastWithdrawalResp, error) {

	var (
		resp = &api.GetLastWithdrawalResp{}
		ex   = g.Ex{}
	)

	ex["user_id"] = in.UserID
	ex["trans_father_id"] = 1
	t, err := model.TransactionLastOne(ex)
	if err != nil {
		return resp, err
	}

	var item api.Transaction
	err = copier.Copy(&item, t)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp.List = &item
	return resp, nil
}
