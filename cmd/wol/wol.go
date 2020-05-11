package main

import (
	"bytes"
	"encoding/hex"
	"net"
)

func main() {
	// Wake from sleep needs enter to remove lockscreen
	conn, err := net.Dial("udp", "192.168.1.255:9")
	if err != nil {
		return
	}

	defer conn.Close()

	mac, err := hex.DecodeString("b8aeed729759")
	if err != nil {
		return
	}

	marker, err := hex.DecodeString("ffffffffffff")
	if err != nil {
		return
	}

	msg := bytes.NewBuffer(marker)
	for i := 0; i < 16; i++ {
		msg.Write(mac)
	}

	conn.Write(msg.Bytes())
}
