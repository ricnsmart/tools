package iot

import (
	"context"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	. "platform/config"
	"time"
)

var cred credentials.TransportCredentials

func Connect(host string) {
	// 获取公钥凭证用于grpc
	var err error

	cred, err = credentials.NewClientTLSFromFile("config/ricnsmart.pem", host)

	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
}

func BatchCheckDeviceNames(request *BatchDeviceInfo) (*BatchRegisterReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &BatchRegisterReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.BatchCheckDeviceNames(ctx, request)
}

func BatchRegisterDeviceWithApplyId(request *BatchRegisterRequest) (*BatchRegisterReply, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &BatchRegisterReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.BatchRegisterDeviceWithApplyId(ctx, request)
}

func DeleteDevice(device *Device) (*NullReply, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &NullReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.DeleteDevice(ctx, device)
}

func GetAllDevices(request *GetRequest) (*GetReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问域名微服务失败: %v", err)
		return &GetReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.GetAllDevices(ctx, request)
}

func DisableThing(device *Device) (*NullReply, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &NullReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.DisableThing(ctx, device)
}

func EnableThing(device *Device) (*NullReply, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &NullReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.EnableThing(ctx, device)
}

func Pub(request *PubRequest) (*PubReply, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &PubReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.Pub(ctx, request)
}

func GetDeviceStatus(device *Device) (*GetSingleDeviceStatusReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &GetSingleDeviceStatusReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.GetDeviceStatus(ctx, device)
}

func BatchGetDeviceState(request *BatchDeviceInfo) (*GetDeviceStatusReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &GetDeviceStatusReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.BatchGetDeviceState(ctx, request)
}

func QueryDeviceDetail(device *Device) (*QueryDeviceReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &QueryDeviceReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.QueryDeviceDetail(ctx, device)
}

func QueryPageByApplyId(request *QueryPageRequest) (*QueryPageReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(IoT.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问物联网平台微服务失败: %v", err)
		return &QueryPageReply{}, err
	}

	c := NewIOTClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.QueryPageByApplyId(ctx, request)
}