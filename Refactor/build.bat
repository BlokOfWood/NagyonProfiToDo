@echo off

set ExeName=server
echo Building %ExeName%...

set GOOS=windows

if not exist .\build (
	mkdir .\build
)

if exist .\build\%ExeName%.exe (
	del .\build\%ExeName%.exe
)

@echo on
go build -v -ldflags "-X 'github.com/CodeFoxHu/go-serverlib.SERVER_BUILT_TIME=%date% %time%'" -o .\build\%ExeName%.exe

@echo off
if exist .\build\%ExeName%.exe (
	pushd .\build
	.\%ExeName%.exe
	popd
)