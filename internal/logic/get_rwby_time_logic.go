package logic

import (
	"context"
	g "github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
	"ks_api_service/internal/model"
	"ks_api_service/internal/svc"
	"ks_api_service/pb/api"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRwbyTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRwbyTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRwbyTimeLogic {
	return &GetRwbyTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRWByTime
func (l *GetRwbyTimeLogic) GetRwbyTime(in *api.GetRwbyTimeReq) (*api.GetRwbyTimeResp, error) {

	var (
		resp = &api.GetRwbyTimeResp{}
		ex   = g.Ex{}
	)

	ex["user_id"] = in.UserID
	ex["trans_father_id"] = []int64{1, 2}
	if in.StartAt != 0 && in.EndAt != 0 {
		ex["created_at"] = g.Op{"between": g.Range(in.StartAt, in.EndAt)}
	}
	ts, err := model.TransactionFindAll(ex)
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
