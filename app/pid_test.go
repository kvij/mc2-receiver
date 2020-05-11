package app

import "testing"

func TestPidByCmdline(t *testing.T) {
	pids := PidByCmdline("spotify %U")

	if pids == nil {
		t.Log("Not running")
	} else {
		t.Log(pids)
	}
}
