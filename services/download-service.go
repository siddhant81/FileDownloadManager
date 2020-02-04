package services

import (
	"github.com/siddhant81/FileDownloadManager/domain/download"
	"github.com/siddhant81/FileDownloadManager/domain/methods"
	"github.com/siddhant81/FileDownloadManager/utils/errors"
)

func CreateDownload(initDownload download.NewDownload) (*download.Download, *errors.RestErr) {
	result, err := move(initDownload)
	if err != nil {
		return result, err
	}

	validateErr := (*result).Validate()
	return result, validateErr
}

func StartDownload(thisDownload *download.Download) *errors.RestErr {
	if thisDownload.DownloadType == "serial" {
		return methods.RequestSerialDownload(thisDownload)
	}
	methods.RequestConcurrentDownload(thisDownload)
	return nil
}

func GetDownload(downloadID string) (*download.Download, *errors.RestErr){
	return download.Get(downloadID)
}

func SuccessfulDownloads() []string {
	return download.Done()
}