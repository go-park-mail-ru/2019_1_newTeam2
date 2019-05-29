go build -o api.out ./cmd/api/main.go;
go build -o auth.out ./cmd/authorization/main.go;
go build -o chat.out ./cmd/chat/main.go;
go build -o game.out ./cmd/game/main.go;
go build -o mgr.out ./cmd/mgr/main.go;
./auth.out ./config/config_auth.json &
./mgr.out ./config/config_score.json &
./api.out ./config/config_api.json &
./chat.out ./config/config_chat.json &
./game.out ./config/config_game.json &
echo "Launched";
