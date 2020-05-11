package app

import "testing"

func TestApp_Start(t *testing.T) {
	if app := New("spotify"); app != nil {
		err := app.Start()
		if err != nil {
			t.Error(err)
		}

		return
	}

	t.Error("App not available")
}