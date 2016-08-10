// Go app template
package main

import (
	"fmt"
	"os"

	"github.com/mandarl/device-mirror/browser"
	"github.com/mandarl/go-selfupdate/selfupdate"
	"github.com/mkideal/cli"
)

var VERSION string = "dev"

type argT struct {
	cli.Helper
	Argument string `cli:"a,arg" usage:"some argument"`
	Version  bool   `cli:"!v,version" usage:"print the current version"`
}

func main() {
	cli.Run(&argT{}, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		run(argv)
		return nil
	})
}

func run(args *argT) {
	if args.Version {
		fmt.Printf("appnameplaceholder: Verison: %s\n", VERSION)
		os.Exit(0)
	}
	runUpdate()

	browser.LaunchChrome()

	fmt.Printf("Hello Mr. %s\n", args.Argument)
}

func runUpdate() {
	var updater = &selfupdate.Updater{
		CurrentVersion: VERSION,
		ApiURL:         "http://dipoletech.com/projects/dist/",
		//u.fetch(u.ApiURL + u.CmdName + "/" + plat + ".json")
		BinURL: "http://dipoletech.com/projects/dist/",
		//u.BinURL + u.CmdName + "/" + u.Info.Version
		//  + "/" + plat + ".gz"
		DiffURL: "",
		//u.fetch(u.DiffURL + u.CmdName + "/" + u.CurrentVersion
		//  + "/" + u.Info.Version + "/" + plat)
		Dir:     "update/",
		CmdName: "appnameplaceholder", // this is added to apiurl to get json
	}

	if updater != nil {
		go updater.BackgroundRun()
	}
}
