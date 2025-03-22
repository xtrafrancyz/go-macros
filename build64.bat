set CGO_ENABLED=1
set CC=C:\Dev\mingw64\bin\gcc
go build -ldflags="-H=windowsgui -linkmode external -extldflags -static" -o go_macros.exe
