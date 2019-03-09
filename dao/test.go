package dao

import (
	"context"
	"github.com/xuruiray/go-web-framework/model"
	"github.com/xuruiray/go-web-framework/storage/mysql"
	"github.com/xuruiray/go-web-framework/util/logger"
)

const (
	selectTestPoByID = "select * from $ where id>#id;"
)

func TestDao(ctx context.Context, id int) ([]model.TestPo, error) {

	var resultList []model.TestPo

	params := map[string]interface{}{
		"id": id,
	}

	err := mysql.GetList(ctx, testTableName, selectTestPoByID, params, resultList)
	if err != nil {
		logger.Error(ctx, "test get list failed|sql=%v|params=%v", selectTestPoByID, params)
		return nil, err
	}

	return resultList, nil
}
