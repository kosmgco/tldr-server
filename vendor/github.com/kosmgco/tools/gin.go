package tools

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type GinApp struct {
	Name        string `default:"$PROJECT_NAME" json:"name,omitempty"`
	IP          string `default:"" json:"ip,omitempty"`
	Port        int    `default:"80" docker:"@lock" json:"port,omitempty"`
	SwaggerPath string `default:"./swagger.json" json:"swaggerPath,omitempty"`
	app         *gin.Engine
}

func (a *GinApp) Init() {
	a.app = gin.New()
	a.app.Use(gin.Recovery(), WithServiceName(a.Name), Logger())
}

type GinEngineFunc func(router *gin.Engine)

func (a *GinApp) Register(ginEngineFunc GinEngineFunc) {
	a.Init()
	ginEngineFunc(a.app)
}

func (a *GinApp) Start() {
	err := a.app.Run(a.getAddr())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Server run failed[%s]\n", err.Error())
		os.Exit(1)
	}
}

func (a GinApp) getAddr() string {
	return fmt.Sprintf("%s:%d", a.IP, a.Port)
}
