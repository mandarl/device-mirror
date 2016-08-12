package browser

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func Launch() {
	switch runtime.GOOS {
	case "darwin":
		launch_darwin()
	case "windows":
		launch_windows()
	case "linux":
		launch_linux()
	}
}

func launch_darwin() {
	var chrome = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	cmd := exec.Command(chrome, "--app=http://127.0.0.1:51232/")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()

}

func launch_windows() {
	cmd := exec.Command(os.Getenv("ProgramFiles(x86)")+"\\Google\\Chrome\\Application\\chrome",
		"--app=http://127.0.0.1:51232/")
	op, err := cmd.Output()
	log.Printf("cmd: %v%v%v", cmd, op, err)
}

func launch_linux() {

}
