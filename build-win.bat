echo off

cd api/serve
    set GOARCH=amd64
    set GOOS=windows

    rice embed-go
    go build -o app.exe
    del /f /a /q rice-box.go
cd ../..