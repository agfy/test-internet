package main

import (
	"flag"
	"fmt"
	"github.com/agfy/test-internet/speedtest"
)

func main() {
	fast := flag.Bool("fast", false, "use fast.com for internet speed evaluation")
	speedTest := flag.Bool("speedtest", false, "use speedtest.com for internet speed evaluation")
	flag.Parse()
	if !*fast && !*speedTest {
		fmt.Println("Error: You must specify -fast or -speedtest.")
		return
	}

	var downloadSpeed, uploadSpeed float64
	if *speedTest {
		fmt.Println("Started evaluation on speedtest.net")
		user, err := speedtest.FetchUserInfo()
		if err != nil {
			fmt.Println("Error: Cannot fetch user information. http://www.speedtest.net/speedtest-config.php is temporarily unavailable. " + err.Error())
			return
		}

		servers, err := speedtest.FetchServers(user)
		if err != nil {
			fmt.Println("Error: Cannot fetch servers. " + err.Error())
			return
		}

		downloadSpeed, uploadSpeed, err = speedtest.StartTest(servers)
		if err != nil {
			fmt.Println("Error: Failed to test speed. " + err.Error())
			return
		}
		fmt.Printf("Download: %5.2f Mbit/s\n", downloadSpeed)
		fmt.Printf("Upload: %5.2f Mbit/s\n\n", uploadSpeed)
	}
}
