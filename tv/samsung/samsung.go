package samsung

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/wpjunior/friday-home/tv"
)

type samsungTV struct{}

func New() tv.TV {
	return &samsungTV{}
}

func (t *samsungTV) TurnOn(ctx context.Context) error {
	return t.runCommand("as")
}

func (t *samsungTV) TurnOff(ctx context.Context) error {
	return t.runCommand("standby 0")
}

func (t *samsungTV) VolumeUp(ctx context.Context) error {
	return t.runCommand("volup")
}

func (t *samsungTV) VolumeDown(ctx context.Context) error {
	return t.runCommand("voldown")
}

func (t *samsungTV) runCommand(cmd string) error {
	c := exec.Command("cec-client", "-s")
	c.Stdin = strings.NewReader(cmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
