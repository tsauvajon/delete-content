export GOARCH=amd64 CGO_ENABLED=0

GOOS=linux go build -o delete-content
zip delete-content-linux64.zip delete-content

GOOS=darwin go build -o delete-content
zip delete-content-macos64.zip delete-content

rm delete-content

GOOS=windows go build -o delete-content.exe
zip delete-content-windows64.zip delete-content.exe

rm delete-content.exe
