
# generate client events
./protoc.exe --proto_path=./schemas --csharp_out=./output ./schemas/*.proto
sleep 10000