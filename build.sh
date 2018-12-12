export GOARCH=amd64 CGO_ENABLED=0

rm -f dist/*

echo "building for linux64"
GOOS=linux go build -o dist/delete-content
zip -j dist/delete-content-linux64.zip dist/delete-content
mv dist/delete-content dist/delete-content-linux

echo "building for macOS64"
GOOS=darwin go build -o dist/delete-content
zip -j dist/delete-content-macos64.zip dist/delete-content
mv dist/delete-content dist/delete-content-macos

echo "building for windows64"
GOOS=windows go build -o dist/delete-content.exe
zip -j dist/delete-content-windows64.zip dist/delete-content.exe
mv dist/delete-content.exe dist/delete-content-windows.exe
