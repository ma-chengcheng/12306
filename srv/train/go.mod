module github.com/mamachengcheng/12306/srv/train

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mamachengcheng/12306/common v0.0.0-20210304041356-64a3f07123ae // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/stretchr/testify v1.6.1 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20210302174412-5ede27ff9881 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.6

)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
