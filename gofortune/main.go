package main

import (
	"bufio"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var files []string

	rand.Seed(time.Now().Unix())

	fortuneCommand := exec.Command("fortune", "-f")
	pipe, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}

	fortuneCommand.Start()

	outputStream := bufio.NewScanner(pipe)
	outputStream.Scan()

	line := outputStream.Text()
	path := line[strings.Index(line, "/"):]

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, "/off/") {
			return nil
		}

		if filepath.Ext(path) == ".dat" {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	i := randomInt(1, len(files))
	randomFile := files[i]

	file, err := os.Open(randomFile)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	quotes := string(b)
	quotesSice := strings.Split(quotes, "%")
	j := randomInt(1, len(quotesSice))

	println(quotesSice[j])
}
