package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/xuruiray/go-web-framework/model"
	"github.com/xuruiray/go-web-framework/service"
)

func Hello(ctx context.Context, request *model.TestRequest) (*model.TestResponse, error) {
	fmt.Println(request.ID)
	resultList, err := service.TestService(ctx, request.ID)
	byteArray, _ := json.Marshal(resultList)

	result := model.TestResponse{
		Data: string(byteArray),
	}

	return &result, err
}
