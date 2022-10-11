@echo off

set ExeName=debug
echo Building %ExeName%...

set GOOS=windows

if not exist .\build (
	mkdir .\build
)

if exist .\build\%ExeName%.exe (
	del .\build\%ExeName%.exe
)

@echo on
go build -v -ldflags "-X 'codefox/serverlib.SERVER_BUILT_TIME=%date% %time%'" -gcflags=all="-N -l" -o .\build\%ExeName%.exe

@echo off
if exist .\build\%ExeName%.exe (
	echo "Build done. Press F5 to start debugging!"
)