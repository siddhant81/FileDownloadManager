package download

import (
	"fmt"
	"github.com/siddhant81/FileDownloadManager/utils/errors"
)

var (
	downloadDB = make(map[string]*Download)
)

func Get(downloadId string) (*Download, *errors.RestErr) {
	result := downloadDB[downloadId]
	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("download ID %s not found", downloadId))
	}

	return result, nil
}

func Done() []string {
	var result []string

	for k,v := range downloadDB {
		if v.Status == "Successful" {
			result = append(result, k)
		}
	}

	return result
}

func (download Download) Save() *errors.RestErr {
	downloadDB[download.Id] = &download
	return nil
}
