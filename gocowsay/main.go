package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"strings"
)

func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ballon []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ballon = append(ballon, top)

	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ballon = append(ballon, s)
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ballon = append(ballon, s)

		index := 1
		for ; index < count-1; index++ {
			s := fmt.Sprintf("%s %s %s", borders[4], lines[index], borders[4])
			ballon = append(ballon, s)
		}

		s = fmt.Sprintf("%s %s %s", borders[2], lines[index], borders[3])
		ballon = append(ballon, s)
	}

	ballon = append(ballon, bottom)

	return strings.Join(ballon, "\n")
}

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}

	return ret
}

func calculateMaxWidth(lines []string) int {
	w := 0

	for _, l := range lines {
		len := utf8.RuneCountInString(l)

		if len > w {
			w = len
		}
	}

	return w
}

func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string

	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}

	return ret
}

func printFigure(name string) {
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	var stegosaurus = `         \                      .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

	`

	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	default:
		fmt.Println("unknonw")
	}
}

func main() {

	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes")
		fmt.Println("usage: fortune | gocowsay")
	}

	reader := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		input, _, err := reader.ReadLine()

		if err != nil && err == io.EOF {
			break
		}

		lines = append(lines, string(input))
	}

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)

	var figure string
	flag.StringVar(&figure, "f", "cow", "select a figure to speak the phrases")

	flag.Parse()

	fmt.Println(balloon)
	printFigure(figure)
	fmt.Println()
}
