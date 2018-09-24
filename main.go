package main

import (
	"github.com/BurntSushi/toml"
	"github.com/atotto/clipboard"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"os/exec"
	"strings"
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
	listAllDesc()
	err, out := getResult("rofi -dmenu -sep '|'", strings.NewReader(listAllDesc()))
	if err != nil {
		log.Fatal(err)
	}
	clipboard.WriteAll(descToText(out))

	exec.Command("sh", "-c", "xdotool key shift+Insert").Run()
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
	if _, err := toml.DecodeFile("/home/tkancf/.config/config.toml", &conf); err != nil {
		log.Fatal(err)
	}

}
