package main

import "godemo/godns/request"

func main() {
	//fmt.Println(request.ParseDomainName("www.baidu.com"))
	data := request.SendData("192.168.1.1:53", "www.baidu.com")
	request.ParseRespond(data)
}
