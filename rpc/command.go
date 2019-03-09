package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/xuruiray/go-web-framework/util/config"
	"math/rand"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// ServiceManager 装载了 HTTP Client 的配置与链接
var ServiceManager map[string]*config.ServiceClient

// Get 封装GET方法
func Get(ctx context.Context, serviceName string, params map[string]string) ([]byte, error) {

	request, err := getRequest(ctx, serviceName, "GET")
	if err != nil {
		return nil, err
	}

	for key, value := range params {
		request = request.Param(key, value)
	}

	_, body, errs := request.EndBytes()
	if errs != nil {
		return nil, errs[0]
	}

	return body, nil
}

// PostForm 封装POST方法
func PostForm(ctx context.Context, serviceName string, params map[string]string) ([]byte, error) {

	request, err := getRequest(ctx, serviceName, "POST")
	if err != nil {
		return nil, err
	}

	_, body, errs := request.Type("form").Send(params).EndBytes()
	if errs != nil {
		return nil, errs[0]
	}

	return body, nil
}

// PostStruct 封装POST方法
func PostStruct(ctx context.Context, serviceName string, data interface{}) ([]byte, error) {
	request, err := getRequest(ctx, serviceName, "POST")
	if err != nil {
		return nil, err
	}

	request.Set("Content-Type", "application/json")
	_, body, errs := request.Type("json").Send(data).EndBytes()
	if errs != nil {
		return nil, errs[0]
	}

	return body, nil
}

//PostJSON 封装POST方法
func PostJSON(ctx context.Context, serviceName string, data map[string]string) ([]byte, error) {

	request, err := getRequest(ctx, serviceName, "POST")
	if err != nil {
		return nil, err
	}

	request.Set("Content-Type", "application/json")
	_, body, errs := request.Type("json").Send(data).EndBytes()
	if errs != nil {
		return nil, errs[0]
	}

	return body, nil
}

func getRequest(ctx context.Context, serviceName string, method string) (*gorequest.SuperAgent, error) {
	client, ok := ServiceManager[serviceName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("can`t find client to serviceName serviceName:%v", serviceName))
	}
	randFactory := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := randFactory.Intn(len(client.IP))

	var request *gorequest.SuperAgent
	switch method {
	case "GET":
		request = gorequest.New().Get("http://" + client.IP[index] + client.URI)
	case "POST":
		request = gorequest.New().Post("http://" + client.IP[index] + client.URI)
	case "PUT":
		request = gorequest.New().Put("http://" + client.IP[index] + client.URI)
	}

	request.Transport.MaxIdleConnsPerHost = client.MaxIdleConnsPerHost
	request.Timeout(time.Duration(client.ConnTimeoutMs)*time.Millisecond).
		Retry(client.RetryCount, time.Duration(client.RetryTimeoutMs)*time.Millisecond, http.StatusInternalServerError)

	return request, nil
}
