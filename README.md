常见的 grpc demo 本来是想给会员中心做准备的，不过，我估计等不到那个时候了
如果你见到了这篇文档，希望能帮到你
made in 武鑫宇
======================= openssl 证书生成 版本 windows openSSl-win64 创建时间为 2020 4 7 过了几年可以重新更新
openssl
genrsa -out ca.key 2048
req -new -x509 -days 3650 -key ca.key -out ca.pem
服务端
genrsa -out server.key 2048
req -new -key server.key -out server.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
客户端
ecparam -genkey -name secp384r1 -out client.key
req -new -key client.key -out client.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
=======================生成文件
files 里设置传递参数
生成service 参数
生成prod.pb.go
protoc --go_out=plugins=grpc:../service Prod.proto
生成prod.pb.gw.go
protoc --grpc-gateway_out=logtostderr=true:../service Prod.proto
=====================
gen.bat 可以直接生成所需要service
===================== geb,bat 备份
cd files && protoc --go_out=plugins=grpc:../service Prod.proto
protoc --go_out=plugins=grpc:../service Orders.proto
protoc --go_out=plugins=grpc:../service --validate_out=lang=go:../service Models.proto
protoc --grpc-gateway_out=logtostderr=true:../service Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../service Orders.proto
protoc --go_out=plugins=grpc:../service Users.proto
protoc --grpc-gateway_out=logtostderr=true:../service Users.proto
cd ..
