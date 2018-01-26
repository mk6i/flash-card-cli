package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Prompt struct {
	tty    *os.File
	reader *bufio.Reader
}

func MakePrompt() *Prompt {

	tty, err := os.Open("/dev/tty")

	if err != nil {
		panic(err)
	}

	return &Prompt{tty, bufio.NewReader(tty)}
}

func (p *Prompt) Speak(message string, voice string) {

	FlushTTYin()

	fmt.Print(strings.TrimSpace(message) + " ")

	textToSpeech(message, voice)

	p.reader.ReadLine()
}

func textToSpeech(message string, voice string) {

	// Convert information in parens to comments so that they are ignored by the say command
	message = strings.Replace(message, "(", "[[", -1)
	message = strings.Replace(message, ")", "]]", -1)

	cmd := exec.Command("/usr/bin/say", "-v", voice, message)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: There was a problem running the voice to text command: %s\n", stderr.String())
		os.Exit(1)
	}
	cmd.Run()
}
