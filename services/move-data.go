package services

import (
	"github.com/rs/xid"
	"github.com/siddhant81/FileDownloadManager/domain/download"
	"github.com/siddhant81/FileDownloadManager/utils/errors"
)

func move(initDownload download.NewDownload) (*download.Download,*errors.RestErr) {
	var tempDownload download.Download

	tempDownload.Id = xid.New().String()
	tempDownload.Status = "Hasn't started yet!"
	tempDownload.DownloadType = initDownload.Type
	tempDownload.Files = make(map[string]string)

	for i := 0; i<len(initDownload.Urls); i++ {
		url := initDownload.Urls[i]
		tempDownload.Files[url] = "Not Downloaded Yet!"
	}

	err := tempDownload.Save()
	if err != nil {
		return &tempDownload, err
	}

	return &tempDownload, nil
}