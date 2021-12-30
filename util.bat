protoc --go_out=../Thyme/ *.proto

protoc  -I ../Thyme/  *.proto --go-grpc_out=../thyme