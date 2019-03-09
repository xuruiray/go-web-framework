package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/xuruiray/gosql"
	"upper.io/db.v3/lib/sqlbuilder"
)

// 初始化 mysql 连接
func Init() {
	var err error
	conn, err = gosql.GetMySQLConn("", "", "", "")
	if err != nil {
		fmt.Println("mysql connect init failed")
	}
}

var conn sqlbuilder.Database

// Execute 封装sql方法 执行sql语句 返回受影响的行数
func Execute(ctx context.Context, tableName string, sqlStr string, params map[string]interface{}) (int64, error) {
	params["table_name"] = tableName
	return gosql.Execute(conn, sqlStr, params)
}

// GetList 封装查询方法 查询多行数据 结果封入res中
func GetList(ctx context.Context, tableName string, sqlStr string, params map[string]interface{}, res interface{}) error {

	return nil
}

// GetOne 封装查询方法 查询单行数据 结果封入res中
func GetOne(ctx context.Context, tableName string, sqlStr string, params map[string]interface{}, res interface{}) error {

	return nil
}

// IsDuplicatedMySQLError 是否是重复键
func IsDuplicatedMySQLError(err error) bool {
	return strings.Contains(err.Error(), "Error 1062: Duplicate")
}
