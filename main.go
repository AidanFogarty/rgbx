package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const usageText = `rgbx - Convert RGB values to hexadecimal color code

Usage:
  rgbx [options] <red> <green> <blue>

Arguments:
  red     Red component (0-255)
  green   Green component (0-255)
  blue    Blue component (0-255)

Options:
  -help      Show this help message`

const RGBLength = 3

func Usage() {
	fmt.Fprintln(flag.CommandLine.Output(), usageText)
	os.Exit(0)
}

func main() {
	help := flag.Bool("help", false, "Show help instructions")

	flag.Parse()

	if *help {
		Usage()
	}

	args := flag.Args()

	if len(args) != RGBLength {
		fmt.Println("error: exactly 3 arguments required: <red> <green> <blue>")
		os.Exit(1)
	}

	red, err := parseComponent(args[0], "red")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	green, err := parseComponent(args[1], "green")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	blue, err := parseComponent(args[2], "blue")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	hexCode := rgbToHex(red, green, blue)
	fmt.Println(hexCode, "\ncopied to clipboard ✅")

	err = copyToClipboard(hexCode)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func parseComponent(arg string, name string) (int, error) {
	value, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("%s component must be an integer", name)
	}
	if value < 0 || value > 255 {
		return 0, fmt.Errorf("%s component must be between 0 and 255", name)
	}
	return value, nil
}

func rgbToHex(red, green, blue int) string {
	return fmt.Sprintf("#%02X%02X%02X", red, green, blue)
}

func copyToClipboard(text string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo "%s" | pbcopy`, text))
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
