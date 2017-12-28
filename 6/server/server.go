package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","),
		"gzip",
	) != -1
}

var contents = []string{
	"これは、私が小さいときに、村のもへいというおじいさんから聞いたお話です。",
	"むかしは、私達の村の近くの、中山というところに小さいお城があって、",
	"中山様というおとのさまが、おられたそうです。",
	"ごんは、ひとりぼっちの小狐で、しだのーぱいしげった森のなかに穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へでてきて、いたずらばかりしました。",
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept: %v\n", conn.RemoteAddr())

	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain; charset=UTF-8",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
