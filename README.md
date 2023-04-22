<h1 align="center">Welcome to rofi-snippet 👋</h1>
<p>
</p>

> `rofi-snippet` is a text snippet tool for Linux

## Requirements

`rofi-snippet` relies on the following command line tools

- `xdotool`
- `xsel`
- `rofi`

In addition, `rofi-snippet` requires [`golang`](https://go.dev/) for compilation.

For example, on `Ubuntu`, you can get what you need with the following commands:

```bash
sudo apt install xsel
sudo apt install xdotool
sudo apt install rofi
sudo apt install golang-go
```

## Installation

1. Clone this repo: `git clone https://github.com/tkancf/rofi-snippet`
2. Change into rofi-snippet: `cd rofi-snippet`
3. Install: `make install`
4. A binary is produced in `~/go/bin`. Make sure `GOPATH` is in your environment variables. To check, run `go env`.
5. Run `rofi-snippet` in the terminal to test.

### Remarks

By default, `config.toml` is installed to `/etc/rofi-snippet/config.toml`. If you want to change this default location, modify `confPath := "/etc" + "/rofi-snippet" + "/config.toml"` in `main.go` and following two lines in `Makefile` to your desired location:

```sh
    @sudo mkdir -p /etc/rofi-snippet/
    @sudo cp ./config.toml /etc/rofi-snippet/
```

## 🚀 Usage

You probably want to bind `rofi-snippet` to a global shortcut. If you're using `i3`, you can add the following line to `~/.config/i3`:

bindsym $mod+Shift+d exec --no-startup-id "rofi-snippet"

Restart `i3`. By pressing `$mod+Shift+d`, it'll bring up the `rofi-snippet`.

## Author

👤 **tkancf**

* Github: [@tkancf](https://github.com/tkancf)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_

