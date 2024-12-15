#!/bin/bash

# 要检查的端口列表
ports=(12002 12000 12001 8500)    # 8500放最后。后面注册的服务都在8500上，关闭8500都会关

for port in "${ports[@]}"; do
    # 查找端口对应的进程
    pid=$(lsof -t -i:$port)

    if [[ -n "$pid" ]]; then
        echo "端口 $port 被进程 $pid 占用，正在关闭..."
        kill -9 $pid
    else
        echo "端口 $port 没有被占用。"
    fi
    sleep 1
done