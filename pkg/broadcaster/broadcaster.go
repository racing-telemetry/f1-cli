package broadcaster

import (
	"context"
	"github.com/racing-telemetry/f1-dump/internal"
	"github.com/racing-telemetry/f1-dump/internal/text/printer"
	"github.com/racing-telemetry/f1-dump/internal/udp"
	"github.com/racing-telemetry/f1-dump/pkg/opts"
	"io"
	"net"
	"os"
	"time"
)

type Broadcaster struct {
	serv *udp.Server

	ctx  context.Context
	stop context.CancelFunc

	Stats *udp.Counter
}

func NewBroadcaster(addr *net.UDPAddr) (*Broadcaster, error) {
	serv, err := udp.Dial(addr)
	if err != nil {
		return nil, err
	}

	return newBroadcaster(serv), nil
}

func newBroadcaster(serv *udp.Server) *Broadcaster {
	ctx, fn := context.WithCancel(context.Background())
	return &Broadcaster{
		serv:  serv,
		ctx:   ctx,
		stop:  fn,
		Stats: new(udp.Counter),
	}
}

func (b *Broadcaster) Start(file string, instant bool) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	offset := int64(0)
	t := 0
	for {
		select {
		case <-b.ctx.Done():
			break
		default:
		}

		buf := make([]byte, internal.BufferSize)

		n, err := f.ReadAt(buf, offset)
		if err != nil {
			if err == io.EOF {
				break
			}

			if opts.Verbose {
				printer.PrintError("reading file: %s", err.Error())
			}
		}

		offset += int64(n)

		if !instant {
			header := new(udp.Header)
			if header.Read(buf) != nil {
				if opts.Verbose {
					printer.PrintError("header read error: %s", err.Error())
				}
			}

			d := int(header.SessionTime * 100000)
			if d != t {
				time.Sleep(time.Nanosecond * time.Duration(d))
				t = d
			}
		}

		err = b.serv.WriteSocket(buf)
		if err != nil {
			if opts.Verbose {
				printer.PrintError("socket write error: %s", err.Error())
			}

			b.Stats.IncErr()
		} else {
			b.Stats.IncRecv()
		}
	}

	return nil
}

func (b *Broadcaster) Stop() {
	b.stop()

	err := b.serv.Close()
	if err != nil {
		if opts.Verbose {
			printer.PrintError("socket closing: %s", err.Error())
		}
	}
}
