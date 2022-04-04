package main

import (
	"flag"
	"fmt"
	"github.com/agfy/test-internet/fasttest"
	"github.com/agfy/test-internet/speedtest"
	"github.com/ddo/go-fast"
)

func main() {
	fastFlag := flag.Bool("fast", false, "use -fast for www.fast.com speed evaluation")
	speedTestFlag := flag.Bool("speedtest", false, "use -speedtest for www.speedtest.net speed evaluation")
	flag.Parse()
	if !*fastFlag && !*speedTestFlag {
		fmt.Println("Error: You must specify -fast or -speedtest.")
		return
	}

	var downloadSpeed, uploadSpeed float64
	if *speedTestFlag {
		fmt.Println("Started evaluation on www.speedtest.net")
		user, err := speedtest.FetchUserInfo()
		if err != nil {
			fmt.Println("Error: Cannot fetch user information. www.speedtest.net/speedtest-config.php is temporarily unavailable. " + err.Error())
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

	if *fastFlag {
		fmt.Println("Started evaluation on www.fast.com")
		fastCom := fast.New()

		err := fastCom.Init()
		if err != nil {
			fmt.Println("Error: Failed to init fast. " + err.Error())
			return
		}

		urls, err := fastCom.GetUrls()
		if err != nil {
			fmt.Println("Error: Failed to get urls. " + err.Error())
			return
		}

		downloadSpeed, uploadSpeed, err = fasttest.StartTest(urls, fastCom)
		//uploadSpeed will always be zero because this go lib for fast.com doesn't provide info about upload speed.
		//To find it - we will probably need further parse of js script from fast.com
		//And maybe impossible without js stuff because to get upload links one need to push "Show more info" button.
		if err != nil {
			fmt.Println("Error: Failed to test speed. " + err.Error())
			return
		}
		fmt.Printf("Download: %5.2f Mbit/s\n", downloadSpeed)
		fmt.Printf("Upload: %5.2f Mbit/s\n\n", uploadSpeed)
	}
}
