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