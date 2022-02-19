package udp

import "sync"

type Counter struct {
	sync.Mutex

	recv uint64 // Success Client Packet Count
	err  uint64 // Fail Client Packet Count
	ign  uint64 // Ignored Client Packet Count
}

func (c *Counter) RecvCount() uint64 {
	return c.recv
}

func (c *Counter) ErrCount() uint64 {
	return c.err
}

func (c *Counter) IgnoredCount() uint64 {
	return c.ign
}

func (c *Counter) IncRecv() {
	c.Lock()
	c.recv += 1
	c.Unlock()
}

func (c *Counter) IncErr() {
	c.Lock()
	c.err += 1
	c.Unlock()
}

func (c *Counter) IncIgnored() {
	c.Lock()
	c.ign += 1
	c.Unlock()
}
