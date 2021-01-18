#!/usr/bin/env bash

#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./genmodel

# 数据库配置
host=127.0.0.1
port=3306
dbname=$1
username=root
passwd=root

goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true