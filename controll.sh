#!/bin/bash

workspace=$(cd $(dirname $0) && pwd -P)
deploy_path=${workspace}
cd ${workspace}

module=bin/ufs
app=${module}

## function
function start() {

    #使用supervisor启动
    exec -c ./$app -deploy=${deploy_path}

    # 启动成功, 退出码为 0
    exit 0
}

action=$1
case $action in
    "start" )
        # 启动服务
        start
        ;;
    * )
        echo "unknown command"
        exit 1
        ;;
esac