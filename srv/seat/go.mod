module github.com/mamachengcheng/12306/srv/seat

go 1.15

require (
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/micro/micro/v3 v3.0.0
	google.golang.org/genproto v0.0.0-20210303154014-9728d6b83eeb // indirect
	gorm.io/gorm v1.21.2 // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
