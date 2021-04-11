@echo off
setlocal
set GO111MODULE=off
go build -ldflags "-H=windowsgui -linkmode=external -s -w" -buildmode=pie -trimpath
endlocal
