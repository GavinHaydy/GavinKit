package tiktok

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
)

func getDevices(networkCardType string) string {
	// 获取网卡
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	var netCard string
	for _, device := range devices {
		if networkCardType == "wifi" && strings.Contains(device.Description, "Wi-Fi") {
			netCard = device.Name
		}
		if networkCardType == "line" && strings.Contains(device.Description, "Connection") {
			netCard = device.Name
		}
	}
	return netCard
}

func GetStreamAddress(networkCardType string) {
	dri := getDevices(networkCardType)
	handle, err := pcap.OpenLive(dri, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter("tcp port 1935"); err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if app := packet.ApplicationLayer(); app != nil {
			data := string(app.Payload())
			if strings.Contains(data, "rtmp://") {
				fmt.Printf("推流地址:%s", data)
			}
			if strings.Contains(data, "stream-") {
				fmt.Printf("推流码：%s", data)
			}
		}
	}
}
