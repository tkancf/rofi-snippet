package main

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/atotto/clipboard"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Settings []Setting `toml:"setting"`
}

type Setting struct {
	Application string `toml:"application"`
	Data        []Data `toml:"data"`
}

type Data struct {
	Desc string `toml:"desc"`
	Text string `toml:"text"`
}

var conf Config

func main() {
	Init()
	list := listAllDesc()
	reader := strings.NewReader(list)

	preClip, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	err, desc := getResult("rofi -dmenu -sep '|'", reader)
	if err != nil {
		log.Fatal(err)
	}
	text := descToText(desc)
	clipboard.WriteAll(text)

	exec.Command("sh", "-c", "xdotool key shift+Insert").Run()

	clipboard.WriteAll(preClip)
}

func descToText(desc string) string {
	desc = strings.TrimRight(desc, "\n")
	text := ""
	for _, value := range conf.Settings {
		for _, v := range value.Data {
			if desc == v.Desc {
				return v.Text
			}
		}
	}
	return text
}

func getResult(command string, r io.Reader) (error, string) {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdin = r
	out, err := cmd.Output()
	result := strings.TrimRight(string(out), "\n")
	return err, result
}

func run(command string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}

func listAllDesc() string {
	var all []byte

	for _, value := range conf.Settings {
		for _, v := range value.Data {
			all = append(all, []byte(v.Desc)...)
			all = append(all, []byte("|")...)
		}
	}

	return string(all)
}

func Init() {
	confPath := "/etc" + "/rofi-snippet" + "/config.toml"
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		log.Fatal(err)
	}

}
