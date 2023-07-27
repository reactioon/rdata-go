package rdata

import (
	
	"fmt"
	"os"
	"net"

)

type CLIENT struct {
	
	Host string
	Port string

}

type CONN struct {
	
	Connection net.Conn
	Server CLIENT

}

var err error
var buffer []byte
var buffetLength int

func _alert(msg string, exit bool) {
	
	fmt.Println(msg)
	
	if exit {
		os.Exit(0)
	}

}

func (c CLIENT) Load(host string, port string) CLIENT {

	dc := CLIENT{
		Host: host,
		Port: port,
	}

	return dc

}

func (c CLIENT) Connect() CONN {

	host := c.Host
	port := c.Port

	connection, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		_alert("Can't connect on server. The server is offline.\n\r", true)
	}

	conn := CONN{
		Connection: connection,
		Server: c,
	}

	return conn

}

func (c CONN) Send(cmd string) string {
	
	ncmd := cmd + "\n"

	_, err = c.Connection.Write([]byte(ncmd))
	buffer = make([]byte, 1024)
	buffetLength, err = c.Connection.Read(buffer)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	answer := string(buffer[:buffetLength])

	buffer = buffer[:0]

	return answer

}

func (c CONN) Close() {

	c.Connection.Close()

}
