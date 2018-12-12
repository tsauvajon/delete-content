export GOARCH=amd64 CGO_ENABLED=0

rm delete-content-*.zip 2> /dev/null
[ -e delete-content ] && rm delete-content

echo "building for linux64"
GOOS=linux go build -o delete-content
zip delete-content-linux64.zip delete-content

rm delete-content

echo "building for macOS64"
GOOS=darwin go build -o delete-content
zip delete-content-macos64.zip delete-content

rm delete-content

echo "building for windows64"
GOOS=windows go build -o delete-content.exe
zip delete-content-windows64.zip delete-content.exe

rm delete-content.exe
