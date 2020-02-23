package main

import (
	"./moodle"
	"fmt"
	"github.com/jrsearles/systray"
	"github.com/zserge/lorca"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	// systray setup
	systray.SetIcon(getIcon("../assets/folder.ico"))
	systray.SetTooltip("Look at me, I'm a tooltip!")
	appClick := systray.AddMenuItem("Open Noodle", "Opens the app")
	quitClick := systray.AddMenuItem("Quit", "Stops Noodle")

	// Loading placeholder for when the application is starting
	initialLoading := "data:text/html," + url.PathEscape(`
				<html>
					<head>
						<title>Loading</title></head>
						<style>.lds-grid{display:inline-block;position:relative;width:80px;height:80px}.lds-grid div{position:absolute;width:16px;height:16px;border-radius:50%;background:#666;animation:lds-grid 1.2s linear infinite}.lds-grid div:nth-child(1){top:8px;left:8px;animation-delay:0s}.lds-grid div:nth-child(2){top:8px;left:32px;animation-delay:-.4s}.lds-grid div:nth-child(3){top:8px;left:56px;animation-delay:-.8s}.lds-grid div:nth-child(4){top:32px;left:8px;animation-delay:-.4s}.lds-grid div:nth-child(5){top:32px;left:32px;animation-delay:-.8s}.lds-grid div:nth-child(6){top:32px;left:56px;animation-delay:-1.2s}.lds-grid div:nth-child(7){top:56px;left:8px;animation-delay:-.8s}.lds-grid div:nth-child(8){top:56px;left:32px;animation-delay:-1.2s}.lds-grid div:nth-child(9){top:56px;left:56px;animation-delay:-1.6s}@keyframes lds-grid{0%,100%{opacity:1}50%{opacity:.5}}</style>
					<body><div class="lds-grid"><div></div><div></div><div></div><div></div><div></div><div></div><div></div><div></div><div></div></div></body>
				</html>
				`)

	// Open Moodle action
	go func() {
		for {
			<-appClick.ClickedCh

			args := []string{}
			if runtime.GOOS == "linux" {
				args = append(args, "--class=Lorca")
			}
			ui, err := lorca.New(initialLoading, "", 375, 540)
			if err != nil {
				log.Fatal(err)
			}

			ui.Bind("newUser", moodle.NewUser)
			ui.Bind("checkCourses", moodle.CheckCourses)

			// Load HTML
			ln, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				log.Fatal(err)
			}
			go http.Serve(ln, http.FileServer(FS))
			ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

			// Wait until the interrupt signal arrives or browser window is closed
			sigc := make(chan os.Signal)
			signal.Notify(sigc, os.Interrupt)
			select {
			case <-sigc:
			case <-ui.Done():
			}

			// Close both the server and the UI
			ln.Close()
			ui.Close()
		}
	}()

	// Quit action
	go func() {
		<-quitClick.ClickedCh
		fmt.Println("Bye bye")
		systray.Quit()
	}()
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
