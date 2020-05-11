package app

import (
	"github.com/kvij/remotemc/xdg"
	"os"
	"os/exec"
)

type App struct {
	xdg.Entry
	Running bool
	pids []int
}

func New(id string) *App {
	if e := xdg.Get(id); e != nil {
		pids := PidByCmdline(e.Exec)
		return &App{
			Entry:   *e,
			Running: pids != nil,
			pids:    pids,
		}
	}
	return nil
}

func (a *App) Start() error {
	cmd := exec.Command("gtk-launch", a.Id)
	err := cmd.Run()
	return err
}

func (a *App) Stop() {
	a.signal(os.Interrupt)
}

func (a *App) Kill() {
	a.signal(os.Kill)
}


func (a *App) signal(sig os.Signal) {
	for _, pid := range a.pids {
		p, err := os.FindProcess(pid)
		if err != nil {
			continue
		}

		p.Signal(sig)
	}
}