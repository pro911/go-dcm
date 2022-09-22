package main

import (
	"apipost-dcm/internal"
	"apipost-dcm/internal/app/router"
	"apipost-dcm/internal/pkg/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	internal.InitProjects()

	r := gin.New()

	router.RegisterRouter(r)

	if err := r.Run(fmt.Sprintf(":%d", conf.Conf.Http.Port)); err != nil {
		panic(fmt.Errorf("httpServer run fail: %w", err))
	}
}
