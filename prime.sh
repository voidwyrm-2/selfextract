exeprime() {
  echo 'building darwin/arm64...'
  GOOS=darwin GOARCH=arm64 go build -o './build/se_darwin-arm64' '.'

  if [ $? ]; then
    bytpend -o build/ build/se_darwin-arm64 magic.txt rsc.zip
    if [ $? ]; then
      rm -rf build/se_darwin-arm64
    else
      echo "appending for '' failed"
    fi
  fi

  echo 'built darwin/arm64'
}

if [ -d build ] || [ -f build ]; then
  rm -rf build
fi

mkdir build

exeprime 'darwin-arm64'

echo 'building windows/amd64...'
GOOS=windows GOARCH=amd64 go build -o './build/se_windows-amd64' '.'

if [ $? ]; then
  bytpend -o build/zse_windows-amd64 build/se_windows-amd64 magic.txt rsc.zip
  if [ $? ]; then
    rm -rf build/se_windows-amd64
  else
    echo "appending failed"
  fi
fi

echo 'built windows/amd64'
