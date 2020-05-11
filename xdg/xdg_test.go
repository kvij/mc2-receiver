package xdg

import "testing"

func TestAddEntries(t *testing.T) {
	entries = make(map[string]*Entry)
	AddEntries("/opt/spotify")
}

func TestGet(t *testing.T) {
	id := "spotify"
	e := Get(id)
	if e == nil {
		t.Error("Not found")
	}

	if e.Id != id {
		t.Error("Id not set match")
	}

	t.Log(*e)
}