package fasttest

import (
	"github.com/ddo/go-fast"
)

func StartTest(urls []string, fastCom *fast.Fast) (float64, float64, error) {
	KbpsChan := make(chan float64)
	downloadSpeeds := make([]float64, 0)

	go func() {
		for Kbps := range KbpsChan {
			downloadSpeeds = append(downloadSpeeds, Kbps/1000)
		}
	}()

	err := fastCom.Measure(urls, KbpsChan)
	if err != nil {
		return 0.0, 0.0, err
	}

	var downloadSpeed float64
	for _, v := range downloadSpeeds {
		downloadSpeed += v
	}
	downloadSpeed = downloadSpeed / float64(len(downloadSpeeds))
	return downloadSpeed, 0.0, nil
}
