package flags

import (
	"errors"
	"fmt"
	"github.com/racing-telemetry/f1-cli/internal"
	"github.com/racing-telemetry/f1-cli/internal/packets"
	"github.com/spf13/cobra"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Flags struct {
	Port  int
	Host  string
	File  string
	Packs packets.IgnoredPackets // ignored packet ids
}

func (f *Flags) UDPAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP(f.Host),
		Port: f.Port,
	}
}

func Add(cmd *cobra.Command) {
	cmd.Flags().IntP("port", "p", internal.DefaultPort, "Port address to listen on UDP.")
	cmd.Flags().String("ip", internal.DefaultHost, "Address to listen on UDP.")
	cmd.Flags().StringP("file", "f", "", "I/O file path or name to save and read packets. (sample: foo/bar/output.bin)")
	cmd.Flags().IntSlice("ignore", []int{}, "Ignore packages by packet ID on recording or broadcasting.")
}

func Build(cmd *cobra.Command) (*Flags, error) {
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return nil, err
	}

	host, _ := cmd.Flags().GetString("ip")
	if host == "" {
		host = internal.DefaultHost
	}

	iArr, err := cmd.Flags().GetIntSlice("ignore")
	if err != nil {
		return nil, err
	}

	packs := packets.GetIgnoredPackets(iArr)
	if packs.HasAllIgnored() {
		return nil, errors.New("all packets are ignored")
	}

	path, _ := cmd.Flags().GetString("file")
	path = strings.TrimSpace(path)
	if path != "" {
		ext := filepath.Ext(path)
		if ext != ".bin" {
			return nil, errors.New("file extension must be ends with .bin")
		}

		_, err = os.Stat(path)
		switch cmd.Name() {
		case "record":
			if err == nil {
				return nil, fmt.Errorf("file already exists: %s", path)
			}

		case "broadcast":
			if err != nil {
				return nil, fmt.Errorf("file doesnt exist: %s", path)
			}
		}
	}

	return &Flags{Port: port, Host: host, File: path, Packs: packs}, nil
}
