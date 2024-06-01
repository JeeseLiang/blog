package service

import (
	"blog/global"
	"blog/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	return &Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}
