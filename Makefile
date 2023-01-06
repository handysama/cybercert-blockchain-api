.PHONY: build clean generate init run validate vendor

APP_NAME=cybercert-blockchain-api
PORT=10788
VERSION=`grep -o 'version:[^\n]*' -m 1 swagger.yaml|sed 's/version: //g'`
VERSION_HASH=`git rev-parse HEAD`

build:
	go build -o ${APP_NAME} -ldflags "-X ${APP_NAME}/config.version=${VERSION} -X ${APP_NAME}/config.versionHash=${VERSION_HASH}" cmd/${APP_NAME}-server/main.go

build-worker:
	go build -o "${APP_NAME}-worker" -ldflags "-X ${APP_NAME}/config.version=${VERSION} -X ${APP_NAME}/config.versionHash=${VERSION_HASH}" worker/main.go

clean:
	rm -rf cmd
	rm -rf models
	rm -rf restapi/operations

generate: validate clean
	swagger generate server -f swagger.yaml -A ${APP_NAME}

init:
	go env -w GO111MODULE=on

run:
	PORT=${PORT} APP_ENVIRONMENT=${APP_ENVIRONMENT} go run cmd/${APP_NAME}-server/main.go

validate:
	swagger validate swagger.yaml

vendor:
	go mod vendor && go mod tidy
