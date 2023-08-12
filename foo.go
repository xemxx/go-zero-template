package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"

	"github.com/xemxx/go-zero-template/internal/config"
	"github.com/xemxx/go-zero-template/internal/handler"
	"github.com/xemxx/go-zero-template/internal/svc"
	"github.com/xemxx/go-zero-template/pkg/mysql"
)

var configFile = flag.String("f", "etc/foo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logc.MustSetup(c.Log)

	mysql.Init(c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
