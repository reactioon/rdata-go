package rdata

import (
	
	"fmt"
	"os"
	"net"
	"runtime"
	"strings"


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

	defer runtime.GC()
	
	println(msg)
	
	if exit {
		os.Exit(0)
	}

}

func Load(host string, port string) CLIENT {

	defer runtime.GC()

	dc := CLIENT{
		Host: host,
		Port: port,
	}

	return dc

}

func (c CLIENT) Connect() (CONN, error) {

	defer runtime.GC()

	var host, port string

	host = c.Host
	port = c.Port

	connection, err := net.Dial("tcp", host+":"+port)
	
	if err != nil {
		connection = nil
		return CONN{}, fmt.Errorf("%s", "[rdata] Can't connect on server. The server is offline.")
	}
	
	conn := CONN{
		Connection: connection,
		Server: c,
	}

	//
	// clear caches
	//
	connection = nil

	return conn, err

}

func (c CONN) Send(cmd string) string {

	defer runtime.GC()

	var buf, tmp []byte
	var ncmd, strBuf, strTrim string
	var n int 
	var errRead error

	if c.Connection == nil {
		return ""
	}

	ncmd = cmd + "\n"
	c.Connection.Write([]byte(ncmd))

	buf = make([]byte, 0, 4096)
    tmp = make([]byte, 256)

    for {

        n, errRead = c.Connection.Read(tmp)

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

	strTrim = string(buf)
	strBuf = strings.TrimSpace(strTrim)

	ncmd = ncmd[:0]
	strTrim = strTrim[:0]
	tmp = nil
	buf = nil
	n = 0
	
	return strBuf

}

func (c CONN) Close() {

	if c.Connection != nil {

		defer runtime.GC()
		c.Connection.Close()

	}

}
