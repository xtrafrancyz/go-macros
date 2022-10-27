set GOARCH=386
set CC=C:\mingw32\bin\gcc
set CGO_ENABLED=1
go build -ldflags="-H=windowsgui -linkmode external -extldflags -static" -o go_macros_x32.exe
