#!/usr/bin/env bash

# 创建模型
xorm reverse mysql root:123456@\(127.0.0.1:3306\)/gin_project?charset=utf8 templates/goxorm
