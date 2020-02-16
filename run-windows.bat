@echo off
go run gen.go
Rem use -ldflags "-H windowsgui" to avoid displaying the terminal
go build -o noodle.exe
move noodle.exe out
cd out
noodle.exe