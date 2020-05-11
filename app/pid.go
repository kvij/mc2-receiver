package app

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func PidByCmdline(exec string) (pids []int) {
	fields := strings.Fields(exec)
	command := filepath.Base(fields[0])

	paths, err := filepath.Glob("/proc/[0-9]*/cmdline")
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		cmdline := scanner.Text()
		file.Close()

		if cmdline == "" {
			continue
		}

		// Filter extra null byte from string
		cmdbytes := []byte(cmdline)
		cmdline = string(cmdbytes[:len(cmdbytes)-1])
		fields := strings.Fields(cmdline)
		if (filepath.Base(fields[0]) == command){
			pidString := filepath.Base(filepath.Dir(path))
			pid, err := strconv.Atoi(pidString)
			if err == nil {
				pids = append(pids, pid)
			}
		}
	}

	return
}