web_path=src/cweb-app
go_path=src/custompbx
user:=$(shell whoami)
os:=$(shell expr substr $(shell uname -s) 1 5)
go_bindata:=$(shell whereis go-bindata)
go_app:=go

ifneq ($(wildcard /usr/local/go/bin),)
    go_app:=/usr/local/go/bin/go
endif

.PHONY: install show install-dep build dep-front dep-back front back front-serve test docker-fmt docker-vet docker-test docker-race docker-frontend-build docker-integration-test docker-release docker-local-recreate docker-local-smoke

install: install-dep build

show:
		@ echo Timestamp: $(shell date)
		@ echo Node Version: $(shell node -v)
		@ echo Npm Version: $(shell npm -v)
		@ echo os: $(os)
		@ echo user: $(user)
		@ echo go_bindata: $(go_bindata)

install-dep: dep-front dep-back
		@ echo install-dep started at: $(shell date)

dep-front:
		@ cd $(web_path) && npm ci

dep-back:
		@ cd $(go_path) && $(go_app) mod download
		@ cd $(go_path) && $(go_app) install github.com/go-bindata/go-bindata/go-bindata@v3.1.2+incompatible

build: front back
		@ echo build started at: $(shell date)

front:
		@ cd $(web_path) && npm run build
		@ cd $(web_path) && go-bindata -pkg cweb -prefix dist/cweb-app/browser -o ../custompbx/cweb/cweb.go dist/cweb-app/browser/...

back:
		@ cd $(go_path) && export CGO_ENABLED=0 && $(go_app) build -ldflags="-s -w" -o ../../bin/cpbx ./
		@ echo build finished at: $(shell date)

front-serve:
		@ [ "${WS_BACKGROUND_OVERRIDE}" ] || ( echo ">> WS_BACKEND_OVERRIDE is not set (syntax wss://HOST:PORT/ws) " )
		@ cd $(web_path) && sed -i "s#WSServ: \"[^\"]*\"#WSServ: \"${WS_BACKEND_OVERRIDE}\"#g" src/environments/environment.ts
		@ cd $(web_path) && npm run start_dev
test:
		@ cd $(go_path) && CUSTOMPBX_CONFIG=../../config.example.json $(go_app) test ./...

docker-fmt:
		@ docker build --target backend-fmt -f docker/Dockerfile .

docker-vet:
		@ docker build --target backend-vet -f docker/Dockerfile .

docker-test:
		@ docker build --target backend-test -f docker/Dockerfile .

docker-race:
		@ docker build --target backend-race -f docker/Dockerfile .

docker-frontend-build:
		@ docker build --target frontend -f docker/Dockerfile .

docker-integration-test:
		@ docker compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from integration-test

docker-release:
		@ docker build -f docker/Dockerfile .

docker-local-recreate:
		@ docker compose -f docker-compose.yml build custompbx
		@ docker rm -f custompbx-host || true
		@ docker compose -f docker-compose.yml up -d --no-deps custompbx
		@ docker ps --filter name=custompbx-host --format 'table {{.Names}}\t{{.Image}}\t{{.Status}}\t{{.Ports}}'

docker-local-smoke:
		@ curl -kfsS --max-time 10 https://127.0.0.1:8081/ >/dev/null || curl -fsS --max-time 10 http://127.0.0.1:8080/ >/dev/null
