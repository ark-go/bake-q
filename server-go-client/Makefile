# source ~/.profile
SHELL := /bin/bash
#arklibgo := ~/ProjectsGo/arkAlias.sh
#version = ~/ProjectsGo/arkAlias.sh getlastversion
PROJECTNAME=$(shell basename `pwd`)
.PHONY: check

.SILENT: build run getlasttag buildzip buildwin buildlinux buildwasm www buildandroid library

test:
	golangci-lint run

buildlinux:
#	@echo $$($(version))
	$(info +Компиляция Linux)
#	go build -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o ./bin/main/$(PROJECTNAME) cmd/main/main.go
	go build -ldflags "-s -w" -o cmd/clientgo ./...
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o cmd/clientgo.exe ./...
	cmd/clientgo

buildzip:
	$(info +Компиляция с сжатием)
	go build -ldflags "-s -w" -o ./bin/main/$(PROJECTNAME) cmd/main/main.go

buildwin:
	$(info +Компиляция windows)
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o ./bin/main/$(PROJECTNAME).exe -tags static -ldflags "-s -w -X 'main.versionProg=$$($(version))'" cmd/main/main.go

#buildwasm:
#	$(info +Компиляция WASM)
#	PROJECTNAME="xxxx" go run gioui.org/src/gogio -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o wasm -target js cmd/main/main.go 

buildandroid:
	$(info +Компиляция Android)
	ANDROID_SDK_ROOT=/home/arkadii/Android/Sdk/ go run gioui.org/src/gogio -ldflags "-s -w -X 'main.versionProg=$$($(version))'" -o ./bin/main/$(PROJECTNAME).apk -target android -icon appicon.png -arch arm64 -appid Go.arkiv cmd/main/main.go

run: buildlinux # buildwin 
	$(info +Запуск)
	./bin/main/$(PROJECTNAME)

    
library: buildlinux
	$(info +Сборка библиотеки)
	go build -o ./bin/main/$(PROJECTNAME).so -buildmode=c-shared cmd/main/main.go

build: buildlinux # library # buildwin #buildwasm buildandroid
	$(info +End)
#www: build
#	$(info +Старт сервера http://172.16.172.10:8080)
#	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("wasm")))'

