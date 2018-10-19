#!/usr/bin/env bash

# 备份数据
xorm dump mysql root:123456@\(127.0.0.1:3306\)/gin_project?charset=utf8 > dbs/data/mysql.sql
