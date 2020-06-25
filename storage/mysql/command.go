package mysql

import (
	"context"
	"github.com/xuruiray/go-web-framework/util/config"
	"strings"
	"time"

	"github.com/xuruiray/gosql"
	"upper.io/db.v3/lib/sqlbuilder"
)

// 初始化 mysql 连接
func Init(file string) (err error) {

	var mysqlConfig config.MySQLConfig

	err = config.LoadConfig(file, &mysqlConfig)
	if err != nil {
		return err
	}

	conn, err = gosql.GetMySQLConn(mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.IP, mysqlConfig.DB)
	if err != nil {
		return err
	}

	conn.SetConnMaxLifetime(time.Duration(mysqlConfig.ConnMaxLifeTimeMs) * time.Millisecond)
	conn.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	conn.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	return nil
}

var conn sqlbuilder.Database

// Execute 封装sql方法 执行sql语句 返回受影响的行数
func Execute(ctx context.Context, tableName string, sqlStr string, params map[string]interface{}) (int64, error) {
	params["table_name"] = tableName
	return gosql.Execute(conn, sqlStr, params)
}

// GetList 封装查询方法 查询多行数据 结果封入res中
func GetList(ctx context.Context, tableName string, sqlStr string, params map[string]interface{}, res interface{}) error {
	params["table_name"] = tableName
	return gosql.QueryList(conn, sqlStr, params, res)
}

// GetOne 封装查询方法 查询单行数据 结果封入res中
func GetOne(ctx context.Context, tableName string, sqlStr string, params map[string]interface{}, res interface{}) error {
	params["table_name"] = tableName
	return gosql.QueryList(conn, sqlStr, params, res)
}

// IsDuplicatedMySQLError 是否是重复键
func IsDuplicatedMySQLError(err error) bool {
	return strings.Contains(err.Error(), "Error 1062: Duplicate")
}
