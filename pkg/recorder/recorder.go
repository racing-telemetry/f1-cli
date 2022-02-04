package recorder

import (
	"bytes"
	"context"
	"fmt"
	"github.com/racing-telemetry/f1-dump/internal"
	"github.com/racing-telemetry/f1-dump/internal/udp"
	"github.com/racing-telemetry/f1-dump/pkg/opts"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type Recorder struct {
	sync.Mutex

	buf  *bytes.Buffer
	serv *udp.Client

	ctx  context.Context
	stop context.CancelFunc

	Stats *udp.Counter
}

func NewRecorder(addr *net.UDPAddr) (*Recorder, error) {
	serv, err := udp.Serve(addr)
	if err != nil {
		return nil, err
	}

	return newRecorder(serv), nil
}

func newRecorder(serv *udp.Client) *Recorder {
	ctx, fn := context.WithCancel(context.Background())
	return &Recorder{
		Stats: new(udp.Counter),
		buf:   new(bytes.Buffer),
		serv:  serv,
		ctx:   ctx,
		stop:  fn,
	}
}

func (r *Recorder) Start() {
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

		r.Stats.IncRecv()

		r.buf.Grow(internal.BufferSize)
		r.buf.Write(buf)
	}
}

func (r *Recorder) Stop() {
	r.stop()
}

func (r *Recorder) Save(file string) (*os.File, error) {
	if r.buf.Len() == 0 {
		return nil, fmt.Errorf("no data found to save")
	}

	if file == "" {
		file = fmt.Sprintf(internal.OutFileFormat, time.Now().Unix())
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
