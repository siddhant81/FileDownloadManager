package app

import (
	"github.com/siddhant81/FileDownloadManager/controllers/downloads"
	"github.com/siddhant81/FileDownloadManager/controllers/health"
)

func mapURLs() {
	router.GET("/health", health.Health)
	router.POST("/downloads", downloads.StartDownloads)
	router.GET("/downloads/:download_id", downloads.DownloadStatus)
	router.GET("/files", downloads.Browse)
}
