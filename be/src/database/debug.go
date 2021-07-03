package database

import (
	"bufio"
	"context"
	"os"
)

type Debug struct {
	ctx <-chan struct{}
}

func NewDebug() *Debug {
	s := &Debug{}
	go s.workers()
	return s
}

func (s *Debug) Stop(ctx context.Context) {
	s.ctx = ctx.Done()
}

func (s *Debug) workers() {
	for {
		select {
		case <-s.ctx:
			return
		default:
			println("SQL >")
			s := bufio.NewScanner(os.Stdin)
			if s.Scan() {
				if err := DB.Exec(s.Text()).Error; err != nil {
					println(err.Error())
					continue
				}
			}
		}
	}
}
