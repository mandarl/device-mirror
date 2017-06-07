package browser

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

//Launch will launch a browser window pointing to the local HTTP server.
//TODO: It will check try to launch in order chrome, FF, IE, Safari.
//It will use appropriate launch method based on platform darwin, win, linux.
func Launch() {
	switch runtime.GOOS {
	case "darwin":
		launchDarwin()
	case "windows":
		launchWindows()
	case "linux":
		launchLinux()
	}
}

func launchDarwin() {
	var chrome = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	cmd := exec.Command(chrome, "--app=http://127.0.0.1:51232/")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()

}

func launchWindows() {
	cmd := exec.Command(os.Getenv("ProgramFiles(x86)")+"\\Google\\Chrome\\Application\\chrome",
		"--app=http://127.0.0.1:51232/", "--user-data-dir", "--no-proxy-server")
	op, err := cmd.Output()
	log.Printf("cmd: %v%v%v", cmd, op, err)
}

func launchLinux() {

}
