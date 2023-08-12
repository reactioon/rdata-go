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
	_,_ = c.Connection.Write([]byte(ncmd))

	buf := make([]byte, 0, 4096)
    tmp := make([]byte, 256)
    for {

        n,_ := c.Connection.Read(tmp)
        buf = append(buf, tmp[:n]...)

		if tmp[n-1] == 10 {
			break;
		}

    }
	
	return string(buf)

}

func (c CONN) Close() {

	c.Connection.Close()

}
