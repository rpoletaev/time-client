package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const epochDelta = 2208988800

func main() {
	errLogger := log.New(os.Stderr, "error:", log.LstdFlags)
	netAddress := getServerAddress()

	conn, err := net.Dial("tcp", netAddress)
	if err != nil {
		errLogger.Fatal(err)
	}

	defer conn.Close()

	buf := make([]byte, 4)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		errLogger.Fatal(err)
	}

	unixTime := int64(binary.BigEndian.Uint32(buf) - epochDelta)
	fmt.Println(unixTime)
}

func getServerAddress() string {
	if len(os.Args) < 3 {
		fmt.Print("you must specify the host and port:\n\ttime-client time.host.com 37\n")
		os.Exit(1)
	}

	return fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])
}
