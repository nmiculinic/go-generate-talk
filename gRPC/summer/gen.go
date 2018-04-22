package summer

/*
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/protoc-gen-gogoslick
go get github.com/gogo/protobuf/gogoproto
go get github.com/gogo/letmegrpc/protoc-gen-letmegrpc
*/
//go:generate bash -c "protoc -I . --gogoslick_out=plugins=grpc:. --letmegrpc_out=.  *.proto"
