@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

: ok

set OLDGOPATH=%GOPATH%
set GOPATH=%cd%

gofmt -w src

go install apiserver

set GOPATH=OLDGOPATH

:end
echo finished