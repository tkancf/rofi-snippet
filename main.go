package main

import (
	"bytes"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	err, out := getResult("rofi -dmenu -sep '|'", strings.NewReader(printKeys()))
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = run("rofi -dmenu", strings.NewReader(string(out)), &buf)
	if err != nil {
		log.Fatal(err)
	}
}

func printKeys() string {
	m := map[string]string{"testん;ほげ": "test文字列", "hoge": "hogeohhoge", "bar": "bar"}
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	keyString := strings.Join(keys, "|")
	return keyString
}

func getResult(command string, r io.Reader) (error, []byte) {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdin = r
	out, err := cmd.Output()
	return err, out
}

func run(command string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}
