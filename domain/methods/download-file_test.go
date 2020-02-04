package methods

import "testing"

func TestDownloadFile(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name : "works fine",
			args : args { "https://images.pexels.com/photos/414612/pexels-photo-414612.jpeg" },
			want : "pexels-photo-414612.jpeg",
			wantErr : false,
		},
		{
			name : "Does not work",
			args : args { "https://imags.pexels.com/photos/414612/pexels-phot612.jpeg" },
			want : "URL not working",
			wantErr : true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err :=  DownloadFile(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}