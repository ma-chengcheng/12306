module order

go 1.15

require (
	github.com/mamachengcheng/12306/srv/seat v0.0.0-20210307072655-2a49a5650ef7 // indirect
	github.com/micro/micro/v3 v3.0.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
