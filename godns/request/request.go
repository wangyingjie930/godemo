package request

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
)

type Header struct {
	TransactionId uint16
	Flags         uint16
	Questions     uint16
	AnswerRRs     uint16
	AuthorityRRs  uint16
	AdditionalRRs uint16
}

// SetFlags 设置标志位
func (h *Header) SetFlags() {

}

type Query struct {
	Type  uint16
	Class uint16
}

// ParseDomainName
//
//	@Description: 按规则解析域名
//	@param domain
//	@return []byte
func ParseDomainName(domain string) []byte {
	splitDomains := strings.Split(domain, ".")
	var ret []byte
	for _, splitDomain := range splitDomains {
		//长度+字节数组
		ret = append(ret, byte(len(splitDomain)))
		ret = append(ret, []byte(splitDomain)...)
	}
	ret = append(ret, 0)
	return ret
}

// FormatData
//
//	@Description: 构造发送数据
//	@param domain
//	@return []byte
func FormatData(domain string) []byte {
	header := &Header{
		TransactionId: 0xFFFF,
		Questions:     1,
		AnswerRRs:     0,
		AuthorityRRs:  0,
		AdditionalRRs: 0,
	}
	header.SetFlags()

	query := &Query{
		Type:  1,
		Class: 1,
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, header)
	binary.Write(&buffer, binary.BigEndian, ParseDomainName(domain))
	binary.Write(&buffer, binary.BigEndian, query)

	return buffer.Bytes()
}

// SendData
//
//	@Description: 发送数据
//	@param dnsServerAddr
//	@param domain
func SendData(dnsServerAddr string, domain string) []byte {
	data := FormatData(domain)
	conn, err := net.Dial("udp", dnsServerAddr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer conn.Close()

	if _, err := conn.Write(data); err != nil {
		log.Fatal(err)
		return nil
	}

	ret := make([]byte, 1024)
	n, err := conn.Read(ret)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return ret[:n]
}

// ParseRespond todo: 解析返回的字节数组
func ParseRespond(data []byte) {
	fmt.Println(data)
}
