package main

import (
	"context"
	"fmt"
	"github.com/aobco/harbor-client/v5/apiv2"
	"github.com/aobco/harbor-client/v5/apiv2/pkg/config"
	"github.com/aobco/log"
	"time"
)

func int64ptr(in int64) *int64 {
	return &in
}

var (
	opts = &config.Options{
		PageSize: 100,
		Page:     0,
		Timeout:  60 * time.Second,
		Sort:     "",
		Query:    "",
	}
	harborClient, _ = apiv2.NewRESTClientForHost("http://192.168.10.201:7508/api", "admin", "Xsky@123", opts)
	ctx             = context.Background()
)

func main() {
	println("================================= ListUsers")
	listUsers, err := harborClient.ListUsers(ctx)
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	for i, user := range listUsers {
		println(i, user.Username)
	}
	println("================================= GetUserByName")
	u, err := harborClient.GetUserByName(ctx, "lb")
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	println(u.Username)
	println("================================= Head Project")
	h, err := harborClient.HeadProject(ctx, "cqh5")
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	fmt.Printf("%v\n", h)
}
