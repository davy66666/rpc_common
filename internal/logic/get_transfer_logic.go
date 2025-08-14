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

type GetTransferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransferLogic {
	return &GetTransferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetTransfer
func (l *GetTransferLogic) GetTransfer(in *api.GetTransferReq) (*api.GetTransferResp, error) {

	var (
		resp = &api.GetTransferResp{}
		ex   = g.Ex{}
	)

	ex["user_id"] = in.UserID
	ex["trans_types_id"] = []int64{19, 20}
	if in.StartAt != 0 && in.EndAt != 0 {
		ex["created_at"] = g.Op{"between": g.Range(in.StartAt, in.EndAt)}
	}
	ts, err := model.GetTransaction(ex)
	if err != nil {
		return resp, err
	}

	for _, t := range ts {

		var item api.Transaction
		err = copier.Copy(&item, t)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		resp.List = append(resp.List, &item)
	}

	return resp, nil
}
