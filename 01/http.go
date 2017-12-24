package main 

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", ":7777")
	defer listener.Close()

	if err != nil {
		fmt.Println("Fail to start server!")
	}
	
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go clientHandler(clientConn)
	}
}


func clientHandler(clientConn net.Conn) {
	defer clientConn.Close()

	buf := make([]byte, 1024)

    _, err := clientConn.Read(buf)

    if err != nil {
       fmt.Println("Fail to get request data!")
    }

    fmt.Println(string(buf))

	const responseData = `
HTTP/1.1 200 OK

Hello, World!
`
	clientConn.Write([]byte(responseData))
}