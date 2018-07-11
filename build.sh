#!/bin/bash
# git pull 

if [ $? -eq 0 ]; then
  echo "start build api"
  ./build_api.sh
else
  echo "build fail"
  exit 1
fi

if [ $? -eq 0 ]; then
  echo "start build console"
  sh ./build_console.sh
else
  echo "build api fail"
  exit 1
fi

if [ $? -eq 0 ]; then
  echo "docker-compose"
  docker-compose up -d --build --remove-orphans
else
  echo "build console fail"
  exit 1
fi

if [ $? -eq 0 ]; then
  echo "docker-compose restart"
  docker-compose restart
else
  echo "docker-compose fail"
  exit 1
fi
