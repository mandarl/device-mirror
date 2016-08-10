package browser

import (
	"log"
	"os"
	"os/exec"
)

func LaunchChrome() {
	//echo := exec.Command("echo", os.Getenv("PATH"))
	cmd := exec.Command(os.Getenv("ProgramFiles(x86)")+"\\Google\\Chrome\\Application\\chrome",
		"--app=http://127.0.0.1:51232/", "--app-shell-host-window-size=800x600")
	op, err := cmd.Output()
	log.Printf("cmd: %v%v%v", cmd, op, err)
	//("chrome --app=http://127.0.0.1:51232/")
}
