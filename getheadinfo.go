/* GetHeadInfo
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	tcpAdrr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAdrr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	//result, err := ioutil.ReadAll(conn)
	result, err := io.ReadAll(conn)
	fmt.Println(string(result))
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("get head info Fatal error: %s", err.Error())
	}
}
