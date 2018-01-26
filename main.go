package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	cmd := flag.NewFlagSet("options", flag.ContinueOnError)
	filePath := cmd.String("file", "", "Path to file containing CSV list")
	leftVoice := cmd.String("left", "Thomas", "Path to file containing CSV list")
	rightVoice := cmd.String("right", "Alex", "Path to file containing CSV list")
	reverse := cmd.Bool("reverse", false, "Reverse column order")

	if err := cmd.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			return
		}
		panic(err)
	}

	reader, err := makeReader(*filePath)

	if err != nil {
		fmt.Println(err.Error())
		cmd.PrintDefaults()
		return
	}

	prompt := MakePrompt()

	for {
		leftVal, rightVal, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if *reverse {
			prompt.Speak(rightVal, *rightVoice)
			prompt.Speak(leftVal, *leftVoice)
		} else {
			prompt.Speak(leftVal, *leftVoice)
			prompt.Speak(rightVal, *rightVoice)
		}

		fmt.Println()
	}
}

func makeReader(filePath string) (*Reader, error) {

	var reader *Reader
	var err error

	if filePath != "" {
		reader, err = FileReader(filePath)
	} else if pipedStdin() {
		reader = StdinReader()
	} else {
		err = fmt.Errorf("Missing input file path")
	}

	return reader, err
}

func pipedStdin() bool {
	info, err := os.Stdin.Stat()

	if err != nil {
		panic(err.Error)
	}

	return (info.Mode() & os.ModeCharDevice) != os.ModeCharDevice
}
