package main

import (
	"fmt"
	"godemo/pprof/data"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
