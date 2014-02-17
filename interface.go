package main

import "os/exec"

var keys map[string]string = map[string]string{
	"left":   "Left",
	"right":  "Right",
	"up":     "Up",
	"down":   "Down",
	"a":      "z",
	"b":      "x",
	"l":      "a",
	"r":      "s",
	"start":  "Return",
	"select": "Backspace",
}

func pressVbaKey(key string) {
	cmd := exec.Command("xdotool", "search", "VBA-M", "windowactivate", "--sync", "key", "--delay", "100", key)
	_ = cmd.Start()
}

func VBAInterface(in <-chan string) {
	for key := range in {
		pressVbaKey(keys[key])
	}
}
