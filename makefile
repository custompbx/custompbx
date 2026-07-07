web_path=src/cweb-app
go_path=src/custompbx
user:=$(shell whoami)
os:=$(shell expr substr $(shell uname -s) 1 5)
go_bindata:=$(shell whereis go-bindata)
go_app:=go

ifneq ($(wildcard /usr/local/go/bin),)
    go_app:=/usr/local/go/bin/go
endif

.PHONY: install show install-dep build dep-front dep-back front back front-serve test

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
