package amap

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

const (
	connectFailed  = "Failed to connect to Location Service"
	connectSucceed = "Location Service connected!"
)

var (
	cred   credentials.TransportCredentials
	target string
)

func Connect(host, address string) {
	var err error

	target = address

	cred, err = credentials.NewClientTLSFromFile("config/ricnsmart.pem", host)

	util.FatalOnError(err, connectFailed, " Address:", target)

	log.Info(connectSucceed, " Address:", target)
}

func ForwardGeocode(request *Address) (*Location, error) {

	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Errorf("访问域名微服务失败: %v", err)
		return &Location{}, err
	}

	c := NewAmapClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.ForwardGeocode(ctx, request)
}
