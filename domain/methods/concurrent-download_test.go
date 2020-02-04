package methods

import (
	"github.com/siddhant81/FileDownloadManager/domain/download"
	"testing"
	"time"
)

func TestRequestConcurrentDownload(t *testing.T) {
	tempDownload := download.Download{"xyz",time.Now(),time.Now(),"Hasn't started yet",
		"serial",map[string]string{"https://images.pexels.com/photos/414612/pexels-photo-414612.jpeg":"xy"}}
	tempDownload1 := download.Download{"xyz",time.Now(),time.Now(),"Hasn't started yet",
		"serial",map[string]string{"https://images.pexels.com/photos/414612/pexels-photo-414612.jpeg":"xy",
			"https://imes.pexels.com/photos/414612/pexels-photo-414612.jpeg":"xy"}}

	type args struct {
		thisDownload *download.Download
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name : "works fine",
			args : args { &tempDownload},
		},
		{
			name : "Failed",
			args : args { &tempDownload1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RequestConcurrentDownload(tt.args.thisDownload)
			time.Sleep(2*time.Second)
			if tt.name == "works fine" && tempDownload.Status == "Failed" {
				t.Errorf("supposed to work fine but failed")
			}
			if tt.name == "Failed" && tempDownload1.Status == "Successful" {
				t.Errorf("supposed to not work but worked")
			}
		})
	}
}