package logic

import (
	"context"
	"github.com/davy66666/rpc_service/internal/model"
	"github.com/davy66666/rpc_service/internal/svc"
	"github.com/davy66666/rpc_service/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransUserCountByUserIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransUserCountByUserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransUserCountByUserIdsLogic {
	return &TransUserCountByUserIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetTransListByTransFatherId
func (l *TransUserCountByUserIdsLogic) TransUserCountByUserIds(in *api.TransUserCountByUserIdsReq) (*api.TransUserCountByUserIdsResp, error) {

	var (
		resp = &api.TransUserCountByUserIdsResp{}
	)

	res, err := model.TransAggByUserAndFatherId(l.ctx, in.UserID, in.TransFatherID, in.StartAt, in.EndAt)
	if err != nil {
		return resp, err
	}

	resp.UserIds = res.UserIDs
	resp.UserCount = res.UniqueUserCount
	resp.SumAmount = res.SumAmount
	return resp, nil
}
