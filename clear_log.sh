#!/bin/bash

# 定义要清空的日志文件路径
log_file="/root/notice/build.log"
log_file1="/root/notice/build2.log"
# 清空日志文件内容
echo "" > $log_file
echo "" > $log_file1
# 输出清空完成的信息
echo "Log file $log_file cleared at $(date)"
echo "Log file $log_file1 cleared at $(date)"
