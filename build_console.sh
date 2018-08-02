#!/bin/bash
cd console

echo "npm install..."

npm install

rm env.js
cp env.prod.js env.js


if [ $? -eq 0 ]; then
  echo "build..."
  npm run build
  echo "build complete"
else
  echo "npm install fail"
  exit 1
fi

if [ $? -eq 0 ]; then
  echo "move dist to /www"
  rm -rf /www/peckergo
  mv dist /www/peckergo
  echo "move complete"
else
  echo "move fail"
  exit 1
fi
