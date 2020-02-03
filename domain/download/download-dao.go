package download

import (
	"fmt"
	"github.com/siddhant81/FileDownloadManager/utils/errors"
)

var (
	downloadDB = make(map[string]*Download)
)

func (download Download) Get() (*Download, *errors.RestErr) {
	result := downloadDB[download.Id]
	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("download ID %d not found", download.Id))
	}

	return result, nil
}

func (download Download) Save() *errors.RestErr {
	current := downloadDB[download.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("Download ID %d already exists, please retry", download.Id))
	}

	downloadDB[download.Id] = &download
	return nil
}
