package rdata

import (
	
	"fmt"
	"os"
	"net"
	"runtime"


)

type CLIENT struct {
	
	Host string
	Port string

}

type CONN struct {
	
	Connection net.Conn
	Server CLIENT

}

func _alert(msg string, exit bool) {
	
	fmt.Println(msg)
	
	if exit {
		os.Exit(0)
	}

}

func (c CLIENT) Load(host string, port string) CLIENT {

	var dc = CLIENT{}

	dc = CLIENT{
		Host: host,
		Port: port,
	}

	return dc

}

func (c CLIENT) Connect() (CONN, error) {

	defer runtime.GC()

	var conn = CONN{}
	var connection net.Conn
	var host, port string
	var err error

	host = c.Host
	port = c.Port

	connection, err = net.Dial("tcp", host+":"+port)
	
	if err != nil {
		connection = nil
		return CONN{}, fmt.Errorf("%s", "[rdata] Can't connect on server. The server is offline.")
	}
	
	conn = CONN{
		Connection: connection,
		Server: c,
	}

	connection = nil

	return conn, err

}

func (c CONN) Send(cmd string) string {

	defer runtime.GC()

	var buf, tmp []byte

	if c.Connection == nil {
		return ""
	}

	ncmd := cmd + "\n"
	c.Connection.Write([]byte(ncmd))

	buf = make([]byte, 0, 4096)
    tmp = make([]byte, 256)
    for {

        n, errRead := c.Connection.Read(tmp)

		if errRead == nil {
			
			buf = append(buf, tmp[:n]...)

			if tmp[n-1] == 10 {
				break;
			}

		}

		if errRead != nil {
			buf = buf[:0]
			break;
		}

    }

	tmp = tmp[:0]
	
	return string(buf)

}

func (c CONN) Close() {

	defer runtime.GC()

	if c.Connection != nil {
		c.Connection.Close()
		c = CONN{}
	}

}
