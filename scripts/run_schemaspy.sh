#!/bin/bash

# データベース接続情報を設定
DB_TYPE=pgsql
DB_HOST=db
DB_PORT=5432
DB_NAME=golang_nextjs
DB_USER=postgres
DB_PASS=password
DB_TIMEZONE=Asia/Tokyo
NETWORK_NAME=golang-nextjs-app_app-network

# SchemaSpyのDockerイメージを取得
docker pull schemaspy/schemaspy

# SchemaSpyを実行
docker run -v "$(pwd)/schema/er_diagram":/output --rm \
  --network ${NETWORK_NAME} \
  schemaspy/schemaspy \
  -t ${DB_TYPE} \
  -host ${DB_HOST}:${DB_PORT} \
  -db ${DB_NAME} \
  -u ${DB_USER} \
  -p ${DB_PASS} \
  -url "jdbc:${DB_TYPE}://${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&serverTimezone=${DB_TIMEZONE}"
