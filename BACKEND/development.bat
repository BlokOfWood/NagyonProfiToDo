@echo off

set ExeName=development.exe
echo Building %ExeName%...

set GOOS=windows

if not exist .\build (
	mkdir .\build
)

if exist .\build\%ExeName% (
	del .\build\%ExeName%
)

@echo on
go build -v -o .\build\%ExeName%

@echo off
if exist .\build\%ExeName% (
	pushd .\build
	.\%ExeName% --address 0.0.0.0 --port 4001
	popd
)
pause