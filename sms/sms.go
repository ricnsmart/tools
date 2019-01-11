package sms

import (
	"context"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

var (
	cred       credentials.TransportCredentials
	smsAddress string
)

func Connect(host, address string) {
	// 获取公钥凭证用于grpc
	var err error

	smsAddress = address

	cred, err = credentials.NewClientTLSFromFile("config/ricnsmart.pem", host)

	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
}

func SendSMS(PhoneNumbers, TemplateCode string, TemplateParam string) (err error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(smsAddress, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := NewSMSClient(conn)

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	_, err = c.SendSms(ctx, &SMSRequest{PhoneNumbers: PhoneNumbers, TemplateParam: TemplateParam, TemplateCode: TemplateCode})

	return
}
