package main

import (
	"context"
	"github.com/aobco/log"
	"github.com/mittwald/goharbor-client/v5/apiv2"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
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
	harborClient, _ = apiv2.NewRESTClientForHost("http://192.168.10.201:7508/api", "admin", "*******", opts)
	ctx             = context.Background()
)

func main() {
	listUsers, err := harborClient.ListUsers(ctx)
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	for i, user := range listUsers {
		println(i, user.Username)
	}
}
