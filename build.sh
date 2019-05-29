#!/usr/bin/env bash

docker images

docker pull serega753/common_db_go_proj:builder || true
docker pull serega753/api_go_proj:builder || true
docker pull serega753/auth_go_proj:builder || true
docker pull serega753/game_go_proj:builder || true
docker pull serega753/chat_go_proj:builder || true

docker build \
  --target builder \
  --cache-from serega753/auth_go_proj:builder \
  -t serega753/auth_go_proj:builder \
  -f "./deploy/auth.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/api_go_proj:builder \
  -t serega753/api_go_proj:builder \
  -f "./deploy/api.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/game_go_proj:builder \
  -t serega753/game_go_proj:builder \
  -f "./deploy/game.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/common_db_go_proj:builder \
  -t serega753/common_db_go_proj:builder \
  -f "./deploy/common_db.Dockerfile" \
  "."

docker build \
  --target builder \
  --cache-from serega753/chat_go_proj:builder \
  -t serega753/chat_go_proj:builder \
  -f "./deploychat.Dockerfile" \
  "."


docker pull serega753/db_go_proj:latest || true
docker pull serega753/common_db_go_proj:latest || true
docker pull serega753/api_go_proj:latest || true
docker pull serega753/auth_go_proj:latest || true
docker pull serega753/game_go_proj:latest || true
docker pull serega753/chat_go_proj:latest || true

docker build \
  --cache-from serega753/db_go_proj:latest \
  -t serega753/db_go_proj:latest \
  -f "./deploy/db.Dockerfile" \
  "."

docker build \
  --cache-from serega753/common_db_go_proj:builder \
  --cache-from serega753/common_db_go_proj:latest \
  -t serega753/common_db_go_proj:latest \
  -f "./deploy/common_db.Dockerfile" \
  "."

docker build \
  --cache-from serega753/api_go_proj:builder \
  --cache-from serega753/api_go_proj:latest \
  -t serega753/api_go_proj:latest \
  -f "./deploy/api.Dockerfile" \
  "."

docker build \
  --cache-from serega753/auth_go_proj:builder \
  --cache-from serega753/auth_go_proj:latest \
  -t serega753/auth_go_proj:latest \
  -f "./deploy/auth.Dockerfile" \
  "."

docker build \
  --cache-from serega753/game_go_proj:builder \
  --cache-from serega753/game_go_proj:latest \
  -t serega753/game_go_proj:latest \
  -f "./deploy/game.Dockerfile" \
  "."

docker build \
  --cache-from serega753/chat_go_proj:builder \
  --cache-from serega753/chat_go_proj:latest \
  -t serega753/chat_go_proj:latest \
  -f "./deploy/chat.Dockerfile" \
  "."

docker push serega753/db_go_proj:latest
docker push serega753/common_db_go_proj:latest
docker push serega753/api_go_proj:latest
docker push serega753/auth_go_proj:latest
docker push serega753/game_go_proj:latest
docker push serega753/chat_go_proj:latest

docker push serega753/common_db_go_proj:builder
docker push serega753/api_go_proj:builder
docker push serega753/auth_go_proj:builder
docker push serega753/game_go_proj:builder
docker push serega753/chat_go_proj:builder
