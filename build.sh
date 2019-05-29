#!/usr/bin/env bash

DOCKER_PASSWORD = $1
DOCKER_USERNAME = $2

docker pull serega753/db_go_proj:builder || true
docker pull serega753/common_db_go_proj:builder || true
docker pull serega753/api_go_proj:builder || true
docker pull serega753/auth_go_proj:builder || true
docker pull serega753/game_go_proj:builder || true
docker pull serega753/chat_go_proj:builder || true

docker build \
  --target builder \
  --cache-from serega753/db_go_proj:builder \
  -t serega753/db_go_proj:builder \
  -f "db.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/auth_go_proj:builder \
  -t serega753/auth_go_proj:builder \
  -f "auth.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/api_go_proj:builder \
  -t serega753/api_go_proj:builder \
  -f "api.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/game_go_proj:builder \
  -t serega753/game_go_proj:builder \
  -f "game.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/common_db_go_proj:builder \
  -t serega753/common_db_go_proj:builder \
  -f "common_db.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/chat_go_proj:builder \
  -t serega753/chat_go_proj:builder \
  -f "chat.Dockerfile" \
  "."


docker pull serega753/db_go_proj:latest || true
docker pull serega753/common_db_go_proj:latest || true
docker pull serega753/api_go_proj:latest || true
docker pull serega753/auth_go_proj:latest || true
docker pull serega753/game_go_proj:latest || true
docker pull serega753/chat_go_proj:latest || true

docker build \
  --cache-from serega753/db_go_proj:builder \
  --cache-from serega753/db_go_proj:latest \
  -t serega753/db_go_proj:latest \
  -f "db.Dockerfile" \
  "."

docker build \
  --cache-from serega753/common_db_go_proj:builder \
  --cache-from serega753/common_db_go_proj:latest \
  -t serega753/common_db_go_proj:latest \
  -f "common_db.Dockerfile" \
  "."

docker build \
  --cache-from serega753/api_go_proj:builder \
  --cache-from serega753/api_go_proj:latest \
  -t serega753/api_go_proj:latest \
  -f "api.Dockerfile" \
  "."

docker build \
  --cache-from serega753/auth_go_proj:builder \
  --cache-from serega753/auth_go_proj:latest \
  -t serega753/auth_go_proj:latest \
  -f "auth.Dockerfile" \
  "."

docker build \
  --cache-from serega753/game_go_proj:builder \
  --cache-from serega753/game_go_proj:latest \
  -t serega753/game_go_proj:latest \
  -f "game.Dockerfile" \
  "."

docker build \
  --cache-from serega753/chat_go_proj:builder \
  --cache-from serega753/chat_go_proj:latest \
  -t serega753/chat_go_proj:latest \
  -f "chat.Dockerfile" \
  "."

docker push serega753/db_go_proj:latest
docker push serega753/common_db_go_proj:latest
docker push serega753/api_go_proj:latest
docker push serega753/auth_go_proj:latest
docker push serega753/game_go_proj:latest
docker push serega753/chat_go_proj:latest

docker push serega753/db_go_proj:builder
docker push serega753/common_db_go_proj:builder
docker push serega753/api_go_proj:builder
docker push serega753/auth_go_proj:builder
docker push serega753/game_go_proj:builder
docker push serega753/chat_go_proj:builder