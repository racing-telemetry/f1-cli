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

func (b *Broadcaster) Start(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	offset := int64(0)
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
