package recorder

import (
	"bytes"
	"context"
	"fmt"
	"github.com/racing-telemetry/f1-dump/cmd/flags"
	"github.com/racing-telemetry/f1-dump/internal"
	"github.com/racing-telemetry/f1-dump/internal/text/printer"
	"github.com/racing-telemetry/f1-dump/internal/udp"
	"github.com/racing-telemetry/f1-dump/pkg/opts"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Recorder struct {
	sync.Mutex

	buf  *bytes.Buffer
	serv *udp.Client

	ctx  context.Context
	stop context.CancelFunc

	flags *flags.Flags
	Stats *udp.Counter
}

func NewRecorder(flags *flags.Flags) (*Recorder, error) {
	serv, err := udp.Serve(flags.UDPAddr())
	if err != nil {
		return nil, err
	}

	return newRecorder(flags, serv), nil
}

func newRecorder(flags *flags.Flags, serv *udp.Client) *Recorder {
	ctx, fn := context.WithCancel(context.Background())
	return &Recorder{
		Stats: new(udp.Counter),
		buf:   new(bytes.Buffer),
		serv:  serv,
		flags: flags,
		ctx:   ctx,
		stop:  fn,
	}
}

func (r *Recorder) Start() {
	hasAnyPacketIgnored := len(r.flags.Packs) != 0
	for {
		select {
		case <-r.ctx.Done():
			return
		default:
		}

		buf, err := r.serv.ReadSocket()
		if err != nil {
			if opts.Verbose {
				log.Println(err)
			}

			r.Stats.IncErr()
			continue
		}

		if hasAnyPacketIgnored {
			header := new(udp.Header)
			if header.Read(buf) != nil {
				if opts.Verbose {
					printer.PrintError("header read error: %s", err.Error())
				}
			}

			if r.flags.Packs.IsIgnored(int(header.PacketID)) {
				r.Stats.IncIgnored()
				continue
			}
		}

		r.Stats.IncRecv()

		r.buf.Grow(internal.BufferSize)
		r.buf.Write(buf)
	}
}

func (r *Recorder) Stop() {
	r.stop()
}

func (r *Recorder) Save() (*os.File, error) {
	if r.buf.Len() == 0 {
		return nil, fmt.Errorf("no data found to save")
	}

	file := r.flags.File
	if file == "" {
		file = fmt.Sprintf(internal.OutFileFormat, time.Now().Unix())
	}

	if err := os.MkdirAll(filepath.Dir(file), 0770); err != nil {
		return nil, err
	}

	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(file, r.buf.Bytes(), 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}
