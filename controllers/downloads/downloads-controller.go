package downloads

import (
	"github.com/gin-gonic/gin"
	"github.com/siddhant81/FileDownloadManager/domain/download"
	"github.com/siddhant81/FileDownloadManager/services"
	"github.com/siddhant81/FileDownloadManager/utils/errors"
	"net/http"
)

func StartDownloads(c *gin.Context) {
	var initDownload download.NewDownload
	err := c.ShouldBindJSON(&initDownload)

	if err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status,restErr)
		return
	}

	result, createErr := services.CreateDownload(initDownload)

	if createErr != nil {
		c.JSON(createErr.Status,createErr)
		return
	}

	startErr := services.StartDownload(result)

	if startErr != nil {
		c.JSON(startErr.Status,startErr)
		return
	}

	c.JSON(200, result.Id)
	return
}

func DownloadStatus(c *gin.Context) {
	downloadId := c.Param("download_id")

	requestedDownload, getErr := services.GetDownload(downloadId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, requestedDownload)
}

func Browse(c *gin.Context) {
	result := services.SuccessfulDownloads()
	c.JSON(http.StatusOK, result)
}