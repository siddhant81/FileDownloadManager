package methods

import (
	"github.com/siddhant81/FileDownloadManager/domain/download"
	"github.com/siddhant81/FileDownloadManager/utils/errors"
	"time"
)

func RequestSerialDownload(thisDownload *download.Download) *errors.RestErr{
	flag := 0
	thisDownload.Status = "Queued"
	thisDownload.StartTime = time.Now()

	for url := range thisDownload.Files {
		name, err := DownloadFile(url)
		if err != nil {
			thisDownload.Files[url] = name
			flag = 1
			continue
		}

		thisDownload.Files[url] = name
	}

	if flag == 1 {
		thisDownload.Status = "Failed"
		return errors.NewBadRequestError("Could not download one or more files")
	}

	thisDownload.Status = "Successful"
	thisDownload.EndTime = time.Now()
	(*thisDownload).Save()
	return nil
}

