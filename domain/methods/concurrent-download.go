package methods

import (
	"github.com/siddhant81/FileDownloadManager/domain/download"
	"sync"
	"time"
)

func RequestConcurrentDownload(thisDownload *download.Download) {
	(*thisDownload).Status = "Queued"
	(*thisDownload).StartTime = time.Now()

	maxConcurrentGoroutines := 5
	concurrentGoroutines := make(chan struct{}, maxConcurrentGoroutines)

	var wg sync.WaitGroup

	for url := range (*thisDownload).Files{
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			concurrentGoroutines <- struct{}{}

			filepath, err := DownloadFile(url)

			(*thisDownload).Files[url] = filepath
			if err != nil {
				(*thisDownload).Status = "Failed"
			}

			<-concurrentGoroutines
		}(url)
	}
	wg.Wait()

	if (*thisDownload).Status != "Failed" {
		(*thisDownload).Status = "Successful"
	}

	(*thisDownload).EndTime = time.Now()
	(*thisDownload).Save()
	return
}

