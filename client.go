package main

import (
	"log"

	client "adarocket/rocket/client"

	"google.golang.org/grpc"
)

// const (
// 	username        = "admin1"
// 	password        = "secret"
// 	refreshDuration = 30 * time.Second
// )

// const serverURL = "127.0.0.1:5300"
const serverURL = "165.22.92.139:5300"

var userToken string

var authClient *client.AuthClient
var informClient *client.InformClient
var interceptor *client.AuthInterceptor

func setupAuthClient() {
	clientConn, err := grpc.Dial(serverURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient = client.NewAuthClient(clientConn)
}

func setupInterceptorAndClient(accessToken string) {
	transportOption := grpc.WithInsecure()

	interceptor, err := client.NewAuthInterceptor(authMethods(), accessToken)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	clientConn, err := grpc.Dial(serverURL, transportOption, grpc.WithUnaryInterceptor(interceptor.Unary()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	informClient = client.NewInformClient(clientConn)
}

// Методы для которых необходима авторизация
// Скорее всего это будет не нужно, так как для всех запросов нужна авторизация
func authMethods() map[string]bool {
	const informerServicePath = "/proto.Informer/"

	return map[string]bool{
		informerServicePath + "GetStatistic": true,
		informerServicePath + "GetNodeList":  true,
	}
}
