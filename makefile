node_version:=$(shell node -v)
npm_version:=$(shell npm -v)
web_path=src/cweb-app
go_path=src/custompbx
user:=$(shell whoami)
os:=$(shell expr substr $(shell uname -s) 1 5)
go_bindata:=$(shell whereis go-bindata)

.PHONY: install show install-dep build dep-front dep-back front back

install: install-dep build

show:
		@ echo Timestamp: $(shell date)
		@ echo Node Version: $(node_version)
		@ echo npm_version: $(npm_version)
		@ echo os: $(os)
		@ echo user: $(user)
		@ echo go_bindata: $(go_bindata)

install-dep: dep-front dep-back
		@ echo install-dep started at: $(shell date)

dep-front:
		@ apt update
		@ apt install go-bindata
		@ cd $(web_path) && npm install

dep-back:
		@ cd $(go_path) && go mod tidy

build: front back
		@ echo build started at: $(shell date)

front:
		@ cd $(web_path) && npm run build --omit=dev
		@ cd $(web_path) && go-bindata -pkg cweb -prefix dist/cweb-app -o ../custompbx/cweb/cweb.go dist/cweb-app/...

back:
		@ cd $(go_path) && export CGO_ENABLED=0 && go build -ldflags="-s -w" -o ../../bin/cpbx ./
		@ echo build finished at: $(shell date)



