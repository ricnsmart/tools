package vms

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

var (
	cred credentials.TransportCredentials
	addr string
)

func Connect(host, address string) {
	// 获取公钥凭证用于grpc
	var err error

	cred, err = credentials.NewClientTLSFromFile("config/ricnsmart.pem", host)

	if err != nil {
		log.Fatal(err)
	}
	addr = address
}

func Call(CalledNumber, TtsCode, TtsParam string) (*Reply, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(cred))
	if err != nil {
		return nil, err
	}
	c := NewVMSClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reply, err := c.Call(ctx, &Request{CalledNumber: CalledNumber, TtsCode: TtsCode, TtsParam: TtsParam})
	if err != nil {
		return reply, err
	}
	if reply.Code != "OK" {
		return reply, errors.New(string(reply.Message))
	}
	return reply, nil
}
