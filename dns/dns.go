package dns

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

func CheckDomainRecord(request *DomainRecord) (*CheckReply, error) {

	conn, err := grpc.Dial(Dns.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问域名微服务失败: %v", err)
		return &CheckReply{}, err
	}

	c := NewDNSClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.CheckDomainRecord(ctx, request)
}

func UpdateDomainRecord(request *DomainRecord) (*NullReply, error) {

	conn, err := grpc.Dial(Dns.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问域名微服务失败: %v", err)
		return &NullReply{}, err
	}

	c := NewDNSClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.UpdateDomainRecord(ctx, request)
}

func DeleteGetDomainRecord(request *DelRequest) (*NullReply, error) {

	conn, err := grpc.Dial(Dns.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问域名微服务失败: %v", err)
		return &NullReply{}, err
	}

	c := NewDNSClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.DeleteDomainRecord(ctx, request)
}

func GetDomainRecords(request *GetRequest) (*GetReply, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(Dns.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问域名微服务失败: %v", err)
		return &GetReply{}, err
	}

	c := NewDNSClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.GetDomainRecords(ctx, request)
}

func AddDomainRecord(record *DomainRecord) (*DomainRecord, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(Dns.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问域名微服务失败: %v", err)
		return &DomainRecord{}, err
	}

	c := NewDNSClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.AddDomainRecord(ctx, record)
}

func SetDomainRecordStatus(request *DomainRecord) (*NullReply, error) {

	conn, err := grpc.Dial(Dns.Address, grpc.WithTransportCredentials(cred))

	if err != nil {
		log.Fatalf("访问域名微服务失败: %v", err)
		return &NullReply{}, err
	}

	c := NewDNSClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	return c.SetDomainRecordStatus(ctx, request)

}
