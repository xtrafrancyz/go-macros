package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	_ "embed"

	"github.com/haroflow/go-macros/automation"
	"github.com/haroflow/go-macros/automation/mouse"
	hook "github.com/robotn/gohook"
	"github.com/zserge/lorca"
)

var ui lorca.UI

// startUI shows www/index.html and blocks until the user closes the window.
func startUI() error {
	var err error

	ui, err = lorca.New("", "", 500, 600, "--remote-allow-origins=*")
	if err != nil {
		return err
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer ln.Close()

	http.HandleFunc("/help", helpHandler)

	http.Handle("/", http.FileServer(http.FS(fs)))

	go http.Serve(ln, nil)
	go listenHotkeys()

	initJavascriptVM()
	registerJavascriptFunctions()
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))

	<-ui.Done()

	return nil
}

func helpHandler(rw http.ResponseWriter, r *http.Request) {
	type TmplData struct {
		Commands []automation.Command
	}

	helpTemplate.Execute(rw, TmplData{
		Commands: registeredCommands,
	})
}

// listenHotkeys waits for key presses to determine what macro to run.
func listenHotkeys() {
	evChan := hook.Start()
	defer hook.End()

	for e := range evChan {
		if e.Kind != hook.KeyUp {
			continue
		}

		// raw hotkey code
		// https://www.toptal.com/developers/keycode
		if conf.RawHotkey != 0 && e.Rawcode == conf.RawHotkey {
			hotkeyPressed()
		}

		// named hotkey
		if val, ok := hook.Keycode[conf.Hotkey]; ok && val == e.Keycode {
			hotkeyPressed()
		}
	}
}

func hotkeyPressed() {
	fmt.Println("hotkey pressed")
	if isRunningMacro.Load().(bool) {
		fmt.Println("macro running... stopping macro")
		stopMacros()
	} else {
		fmt.Println("nothing running... starting macro")
		code := ui.Eval("getCode()").String()
		go executeCode(code)
	}
}

// registerJavascriptFunctions enables JS functions to call Go functions.
func registerJavascriptFunctions() {
	ui.Bind("executeCode", executeCode)
	ui.Bind("stopMacros", stopMacros)
	ui.Bind("getMousePosition", getMousePosition)
	ui.Bind("setHotkey", setHotkey)
	ui.Bind("getHotkey", getHotkey)
	ui.Bind("getSavedScripts", getSavedScripts)
	ui.Bind("saveScript", saveScript)
	ui.Bind("deleteScript", deleteScript)
}

func setHotkey(key string) {
	conf.Hotkey = key
	saveConfig(conf)
}

func getHotkey() string {
	return strings.ToUpper(conf.Hotkey)
}

func getSavedScripts() map[string]string {
	return conf.SavedScripts
}

func saveScript(name, code string) {
	conf.SavedScripts[name] = code
	saveConfig(conf)
}

func deleteScript(name string) {
	delete(conf.SavedScripts, name)
	saveConfig(conf)
}

type MousePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func getMousePosition() MousePosition {
	return MousePosition{
		mouse.GetX(),
		mouse.GetY(),
	}
}
