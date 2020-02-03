package download

import (
	"github.com/siddhant81/FileDownloadManager/utils/errors"
	"time"
)

type Download struct {
	Id 			   string		`json:"id"`
	StartTime 	   time.Time 	`json:"start_time"`
	EndTime 	   time.Time	`json:"end_time"`
	Status 		   string	`json:"status"`
	DownloadType   string `json:"download_type"`
	Files          map[string]string `json:"files"`
}

type NewDownload struct {
	Type  string  `json:"type"`
	Urls  []string  `json:"urls"`
}

func (download *Download) Validate() *errors.RestErr {
	if download.DownloadType != "serial" && download.DownloadType != "concurrent" {
		return errors.NewBadRequestError("Invalid Download Type")
	}
	return nil
}