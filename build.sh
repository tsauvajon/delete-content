export CGO_ENABLED=0

rm -f dist/*

# arguments :
# $1 GOOS
# $2 GOARCH
# $3 textual platform
# $4 zipped executable name
# $5 unzipped executable name
build_for()
{
    echo "building for $1-$2"
    GOOS=$1 GOARCH=$2 go build -o dist/$4
    zip -j dist/delete-content-$3-$2.zip dist/$4
    mv dist/$4 dist/$5
}

build_for linux amd64 linux delete-content delete-content-linux-amd64
build_for darwin amd64 macos delete-content delete-content-macos-amd64
build_for windows amd64 windows delete-content.exe delete-content-windows-amd64.exe
