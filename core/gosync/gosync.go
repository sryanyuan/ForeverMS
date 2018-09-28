package gosync

import (
	"context"
	"sync"
)

// Context provides routine synchronize components
type Context struct {
	ctx      context.Context
	wg       *sync.WaitGroup
	cancelFn context.CancelFunc
}

func NewContextWithCancel() *Context {
	ctx, cancelFn := context.WithCancel(context.Background())
	return &Context{
		ctx:      ctx,
		wg:       &sync.WaitGroup{},
		cancelFn: cancelFn,
	}
}

func (c *Context) Add(delta int) {
	c.wg.Add(delta)
}

func (c *Context) Done() {
	c.wg.Done()
}

func (c *Context) Cancelled() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Wait() {
	c.wg.Wait()
}

func (c *Context) Cancel() {
	c.cancelFn()
}
