package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"ks_api_service/internal/model"
	"ks_api_service/internal/types"

	"ks_api_service/internal/svc"
	"ks_api_service/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransListByTransFatherIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTransListByTransFatherIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransListByTransFatherIdLogic {
	return &GetTransListByTransFatherIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetTransListByTransFatherId
func (l *GetTransListByTransFatherIdLogic) GetTransListByTransFatherId(in *api.GetTransListByTransFatherIdReq) (*api.GetTransListByTransFatherIdResp, error) {

	var (
		resp  = &api.GetTransListByTransFatherIdResp{}
		param = types.GetTransListByTransFatherIdParams{}
		page  = 1
		size  = 1000
	)

	ts, _, err := model.TransListByTransFatherId(l.ctx, param, page, size)
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
