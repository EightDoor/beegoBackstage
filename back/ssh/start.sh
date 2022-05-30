#!/bin/bash
chmod +x ./back
export APP_RUN_MODE=prod
nohup ./back > back.log 2>&1 &
