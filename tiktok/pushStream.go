package tiktok

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"regexp"
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

func GetStreamAddress(networkCardType string) []string {
	dri := getDevices(networkCardType)
	handle, err := pcap.OpenLive(dri, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter("tcp port 1935"); err != nil {
		log.Fatal(err)
	}
	var (
		rtmpURL   string
		streamKey string
	)

	urlSplit := func(startWith, url string) string {
		var pat string
		if startWith == "r" {
			pat = `rtmp[^\s]*?game`
		} else {
			pat = `stream-[^\s]*?True`
		}

		re := regexp.MustCompile(pat)
		return re.FindString(url)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if app := packet.ApplicationLayer(); app != nil {
			data := string(app.Payload())
			if rtmpURL == "" && strings.Contains(data, "rtmp://") {
				rtmpURL = urlSplit("r", data)
			}
			if streamKey == "" && strings.Contains(data, "stream-") {
				streamKey = urlSplit("s", data)
			}

			// 如果都抓到了，就返回
			if rtmpURL != "" && streamKey != "" {
				return []string{rtmpURL, streamKey}
			}
		}
	}
	return []string{}
}
