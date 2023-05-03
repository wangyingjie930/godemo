/**
  @author: wangyingjie
  @since: 2023/5/2
  @desc: 实现ping命令, go run main.go www.baidu.com
**/

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size    int
	count   int
	typ     uint8 = 8
	code    uint8 = 0
)

// ICMP ICMP报文头部信息, 需要按照icmp报文规定好的顺序
type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	ID          uint16
	SequenceNum uint16
}

func main() {
	GetCommonArgs()
	desIp := os.Args[len(os.Args)-1]

	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	fmt.Println(remoteAddr)

	for i := 0; i < count; i++ {
		icmp := &ICMP{
			Type:        typ,
			Code:        code,
			Checksum:    0,
			ID:          uint16(i),
			SequenceNum: uint16(i),
		}
		var buffer bytes.Buffer
		binary.Write(&buffer, binary.BigEndian, icmp)
		data := make([]byte, size)
		//往切片里面写入数据
		copy(data, "hello world")
		//放入待发送的数据中
		buffer.Write(data)

		data = buffer.Bytes()
		//数据加入签名
		checksum := CheckSum(data)
		data[2] = byte(checksum >> 8) //高8位放入
		data[3] = byte(checksum)      //第8位放入

		//设置超时时间
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		n, err := conn.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 65535)
		startTime := time.Now()
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		//buf中的数据是从ip层的数据开始计算
		fmt.Printf("来自 %d.%d.%d.%d 的回复: 字节=%d 时间=%dms TTL=%d\n",
			buf[12], buf[13], buf[14], buf[15],
			n-28, //IPV4头部20个字节,ICMP头部8个, 通过抓包看到的
			time.Since(startTime).Milliseconds(), buf[8])
	}
}

func GetCommonArgs() {
	flag.Int64Var(&timeout, "w", 1000, "超时时间")
	flag.IntVar(&size, "l", 32, "发送字节数")
	flag.IntVar(&count, "n", 4, "请求次数")
	flag.Parse()
}

//CheckSum 校验和计算
func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length = len(data)
		index  int
	)
	//两两拼接并且求和
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	//奇数剩余1位, 单个加入求和
	if length > 0 {
		sum += uint32(data[index])
	}
	hi := sum >> 16
	for hi != 0 {
		//高16位+低16位
		sum = hi + uint32(uint16(sum)) //因为sum为32位的, uint16(sum)直接取低16位的, 溢出部分不计入
		hi = sum >> 16
	}
	return uint16(^sum)
}
