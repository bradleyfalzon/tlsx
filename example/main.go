package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bradleyfalzon/tlsx"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	iface := flag.String("iface", "eth0", "Network interface to capture on")
	flag.Parse()

	handle, err := pcap.OpenLive(*iface, 1500, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	err = handle.SetBPFFilter("(dst port 443)")
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Listening on", *iface)
	for packet := range packetSource.Packets() {
		go readPacket(packet)
	}
}

func readPacket(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			log.Println("Could not decode TCP layer")
			return
		}
		if tcp.SYN {
			// Connection setup
		} else if tcp.FIN {
			// Connection teardown
		} else if tcp.ACK && len(tcp.LayerPayload()) == 0 {
			// Acknowledgement packet
		} else if tcp.RST {
			// Unexpected packet
		} else {
			// data packet
			readData(packet)
		}
	}
}

func readData(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		t, _ := tcpLayer.(*layers.TCP)

		var hello = tlsx.ClientHello{}

		err := hello.Unmarshall(t.LayerPayload())

		switch err {
		case nil:
		case tlsx.ErrHandshakeWrongType:
			return
		default:
			log.Println("Error reading Client Hello:", err)
			log.Println("Raw Client Hello:", t.LayerPayload())
			return
		}
		log.Printf("Client hello from port %s to %s", t.SrcPort, t.DstPort)
		fmt.Println(hello)
	} else {
		log.Println("Client Hello Reader could not decode TCP layer")
		return
	}
}
