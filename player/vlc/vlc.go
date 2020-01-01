package vlc

import (
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/wpjunior/friday-home/player"
)

var mu sync.Mutex
var cmd *exec.Cmd

type vlc struct{}

func New() player.Player {
	return &vlc{}
}

func (v *vlc) PlayYoutubeChannel(channel string) error {
	mu.Lock()
	if cmd != nil {
		cmd.Process.Kill()
	}

	cmd = exec.Command("cvlc", channel)
	mu.Unlock()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var err error

	go func() {
		err = cmd.Run()
	}()

	time.Sleep(time.Second)

	return err
}
