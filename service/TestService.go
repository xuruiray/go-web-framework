package service

import (
	"context"
	"github.com/xuruiray/go-web-framework/dao"
	"github.com/xuruiray/go-web-framework/model"
)

func TestService(ctx context.Context, id int) ([]model.TestPo, error) {
	return dao.TestDao(ctx, id)
}
