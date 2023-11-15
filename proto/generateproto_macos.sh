# generate client events
rm -rf ./output/*
./protoc --proto_path=./schemas --csharp_out=./output ./schemas/*.proto