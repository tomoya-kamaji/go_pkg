# コマンド確認
help:
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'


api-up: # api-server起動
	docker-compose -f tools/docker-compose.yml up -d

.PHONY: goose-create
goose-create: # マイグレーションファイル作成 ex) make goose-create name=new_migration_name
	MIGRATION_FILE=$(name) sh migrations/goose.sh create

goose-up: # マイグレーション実行
	sh migrations/goose.sh up

goose-down: # マイグレーションロールバック
	sh migrations/goose.sh down

goose-status: # マイグレーションステータス
	sh migrations/goose.sh status


build-protoc: # protoファイルからgoファイルを生成
	protoc --go_out=grpc \
	--go-grpc_out=grpc \
	grpc/proto/base_ball_api.proto

wire: # wireで依存関係を解決
	wire ./...