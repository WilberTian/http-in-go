package simple_http

import (
	"bufio"
	"strings"
	"bytes"
	"fmt"
)

type Request struct {
	Method string

	URI string

	Proto string 

	Header Header

	ContentLength int64
}


func parseRequestLine(requestLine string) (method, reuqestURI, proto string, ok bool) {
	s1 := strings.Index(requestLine, " ")
	s2 := strings.Index(requestLine[s1+1:], " ")
	if s1 < 0 || s2 < 0 {
		return
	}

	s2 += s1 + 1
	return requestLine[:s1], requestLine[s1+1 : s2], requestLine[s2+1:], true
}

func parseRequestHeader(bufReader *bufio.Reader) (Header, error) {
	header := Header{}

	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			return nil, err
		}

		if len(line) == 0 {
			return header, nil
		}

		i := bytes.IndexByte(line, ':')
		key := strings.Trim(string(line[:i]), " ")
		value := strings.Trim(string(line[i + 1:]), " ")

		header.Set(key, value)
	}
}

func ParseRequest(requestStr string) {
	strReader := strings.NewReader(requestStr)
	bufReader := bufio.NewReader(strReader)
	req := new(Request)

	requestLine, _, err := bufReader.ReadLine()	

	if err != nil {
		fmt.Println("Fail to get request line!")
	}

	var ok bool
	req.Method, req.URI, req.Proto, ok = parseRequestLine(string(requestLine))	

	if !ok {
		fmt.Println("Fail to parse request line!")
	}

	req.Header, err = parseRequestHeader(bufReader)
	
}

