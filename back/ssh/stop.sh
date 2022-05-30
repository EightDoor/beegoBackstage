#!/bin/bash
PID=$(ps -ef|grep back|grep -v grep|awk '{print $2}')
if [ -z $PID ]; then
	echo "back not exist"
	exit
else
	echo "back id: $PID"
	kill -9 ${PID}
	echo "back 停止"
fi
