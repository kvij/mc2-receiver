package xdg

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Entry struct {
	Id string
	Name string
	Exec string
	Icon string
}

var entries map[string]*Entry

func init() {
	entries = make(map[string]*Entry)

	homedir := os.Getenv("HOME")
	AddEntries("/usr/share/applications")
	AddEntries("/usr/local/share/applications")
	AddEntries( homedir + "/.local/share/applications")
}

// Key = filename.desktop
// Name: first occurrence Name=
// Executable: first occurrence Exec=[^\s]+
// Match /proc/$PID/cmdline with ^Executable for status and pids

func List() map[string]*Entry {
	return entries
}

func Get(id string) *Entry {
	return entries[id]
}

func AddEntries(dir string) {
	paths, err := filepath.Glob(dir + "/*.desktop")
	if err != nil {
		log.Print(err)
		return
	}

	for _, path := range paths {
		e := parseEntry(path)
		entries[e.Id] = e
	}
}



func parseEntry(path string) (e *Entry) {
	e = &Entry{
		Id:   getId(path),
	}

	e.readEntryFile(path)
	return
}

func getId(path string) string {
	runes := []rune(filepath.Base(path))
	index := len(runes)-8 // index of .desktop
	return string(runes[:index])
}

func (e *Entry) readEntryFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")

		if parts[0] == "Name" && e.Name == "" {
			e.Name = parts[1]
		}

		if parts[0] == "Exec" && e.Exec == "" {
			e.Exec = parts[1]
		}

		if parts[0] == "Icon" && e.Icon == "" {
			e.Icon = parts[1]
		}
	}

	return
}