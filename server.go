package mockXray

import (
	"net"
	"log"
)

func StartMockXrayServer() {
	go func() {
		log.Println("Starting Mock XRAY Logger on localhost 2000")
		addr := net.UDPAddr{
			Port: 2000,
			IP:   net.ParseIP("127.0.0.1"),
		}

		conn, err := net.ListenUDP("udp", &addr)

		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		var buf [2048]byte
		log.Println("Waiting for XRAY Traffic")
		for {
			l, _, _ := conn.ReadFromUDP(buf[:])
			log.Print("Received a XRAY Packet:")
			log.Println(string(buf[:l]))
			log.Println("-------------------------------------")
		}
	}()
}
