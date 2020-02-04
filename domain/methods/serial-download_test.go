package methods

import (
	"github.com/siddhant81/FileDownloadManager/domain/download"
	"github.com/siddhant81/FileDownloadManager/utils/errors"
	"reflect"
	"testing"
	"time"
)

func TestRequestSerialDownload(t *testing.T) {
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
		want *errors.RestErr
	}{
		{
			name : "works fine",
			args : args { &tempDownload},
			want : nil,
		},
		{
			name : "Rest Error",
			args : args { &tempDownload1},
			want : errors.NewBadRequestError("Could not download one or more files"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestSerialDownload(tt.args.thisDownload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestSerialDownload() = %v, want %v", got, tt.want)
			}
		})
	}
}