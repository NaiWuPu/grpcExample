cd files && protoc --go_out=plugins=grpc:../service Prod.proto
protoc --go_out=plugins=grpc:../service Orders.proto
protoc --go_out=plugins=grpc:../service --validate_out=lang=go:../service Models.proto
protoc --grpc-gateway_out=logtostderr=true:../service Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../service Orders.proto
protoc --go_out=plugins=grpc:../service Users.proto
protoc --grpc-gateway_out=logtostderr=true:../service Users.proto
cd ..
