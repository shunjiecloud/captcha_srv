package main

import (
	"github.com/micro/go-micro/v2"
	captcha_proto "github.com/shunjiecloud-proto/captcha/proto"
	"github.com/shunjiecloud/captcha-srv/modules"
	"github.com/shunjiecloud/captcha-srv/services"
	"github.com/shunjiecloud/pkg/log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.captcha"),
		micro.WrapHandler(log.LogWrapper),
	)

	// Init
	modules.Setup()
	service.Init()

	// Register Handlers
	captcha_proto.RegisterCaptchaHandler(service.Server(), new(services.CaptchaService))

	// Run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
