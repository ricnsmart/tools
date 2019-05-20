module github.com/ricnsmart/tools

require (
	github.com/go-redis/redis v6.15.0+incompatible
	github.com/golang/protobuf v1.2.0
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/influxdata/influxdb v1.7.0
	github.com/influxdata/platform v0.0.0-20181222011335-94b7c3cea0ef // indirect
	github.com/labstack/gommon v0.2.7
	github.com/onsi/ginkgo v1.6.0 // indirect
	github.com/onsi/gomega v1.4.1 // indirect
	github.com/ricnsmart/plugins v0.0.10
	github.com/ricnsmart/rules v0.0.10
	github.com/satori/go.uuid v1.2.0
	github.com/streadway/amqp v0.0.0-20181205114330-a314942b2fd9
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	google.golang.org/grpc v1.19.0
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.38.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190510104115-cbcb75029529
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190507092727-e4e5bf290fec
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190509164839-32b2708ab171
	golang.org/x/net => github.com/golang/net v0.0.0-20190509222800-a4d6f7feada5
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190511041617-99f201b6807e
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.5.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190508193815-b515fa19cec8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
	labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20180705113738-7446a0344b78
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
