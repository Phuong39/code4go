package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	snapshotLen int32 = 65535
	err         error
	timeout     = 30 * time.Second
	handle      *pcap.Handle
	buffer      gopacket.SerializeBuffer
	options     gopacket.SerializeOptions
)

func main() {
	// 获取 libpcap 的版本
	version := pcap.Version()
	fmt.Println(version)
	// 获取网卡列表
	devices, _ := pcap.FindAllDevs()
	if len(devices) == 0 {
		return
	}
	fmt.Println("网卡列表")
	for _, v := range devices {
		fmt.Println(v.Name, "-", v.Description, "-", v.Addresses)
	}

	// 打开一个网络接口
	handle, err = pcap.OpenLive(devices[2].Name, snapshotLen, false, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Create the layers
	ethernetLayer := &layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
		DstMAC:       net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
		EthernetType: layers.EthernetTypeIPv4,
	}
	ipLayer := &layers.IPv4{
		SrcIP:   net.IP{10, 13, 152, 41},
		DstIP:   net.IP{8, 8, 8, 8},
		Version: 4,
		IHL:     5, // 20 bytes standard header size
		Length:  24,
	}
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(80),
	}
	payload := []byte{10, 20, 30, 40}

	// Serialize the layers
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(payload),
	)
	outgoingPacket := buffer.Bytes()

	// Send packet over the wire (or air)
	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal(err)
	}
}
