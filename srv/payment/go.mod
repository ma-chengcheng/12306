module github.com/mamachengcheng/12306/srv/payment

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/iGoogle-ink/gopay v1.5.30
	github.com/mamachengcheng/12306/common v0.0.0-20210304041356-64a3f07123ae // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/micro/micro/v3 v3.0.0
	google.golang.org/genproto v0.0.0-20210303154014-9728d6b83eeb // indirect
	google.golang.org/protobuf v1.25.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
