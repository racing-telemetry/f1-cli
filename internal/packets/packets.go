package packets

import (
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	Id          int
	Name        string
	Description string
}

type IgnoredPackets []Packet

func (p IgnoredPackets) IsIgnored(pId int) bool {
	for _, packet := range p {
		if packet.Id == pId {
			return true
		}
	}

	return false
}

func (p IgnoredPackets) HasAllIgnored() bool {
	return len(p) == len(packets)
}

func (p IgnoredPackets) Pretty() string {
	items := make([]string, 0, len(p))
	for _, pack := range p {
		items = append(items, fmt.Sprintf("%s(%d)", pack.Name, pack.Id))
	}

	return strings.Join(items, ", ")
}

func GetIgnoredPackets(s []int) IgnoredPackets {
	var packs IgnoredPackets // ignored packets array
	for _, pId := range s {
		p := contains(pId)
		if p != nil {
			packs = append(packs, *p)
		}
	}

	return packs
}

func Pretty() string {
	items := make([]string, 0, len(packets)+2)
	items = append(items, "ID- Packet - Description")
	items = append(items, "------------------------")
	for _, packet := range packets {
		i := strconv.Itoa(packet.Id)
		if packet.Id < 10 {
			i = i + " "
		}

		items = append(items, fmt.Sprintf("%s- %s - %s", i, packet.Name, packet.Description))
	}

	return strings.Join(items, "\n")
}

func contains(pId int) *Packet {
	for _, p := range packets {
		if p.Id == pId {
			return &p
		}
	}
	return nil
}
