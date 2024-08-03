go build -buildmode=exe -ldflags="-s -w" -o ./app/MutoolUI
cp ./bin/liblcl.so ./app
cp ./bin/mutool ./app
cp -r ./docs ./app

