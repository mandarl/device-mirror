package delete

import (
	"fmt"
	"log"

	"github.com/zach-klippenstein/goadb"
)

func IsDeviceConnected() {
	client, err := adb.NewWithConfig(adb.ServerConfig{})
	if err != nil {
		log.Fatal(err)
	}

	client.KillServer()
	client.StartServer()
	devices, err := client.ListDevices()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Devices:")
	for _, device := range devices {
		fmt.Printf("\t%+v\n", *device)
	}
}
