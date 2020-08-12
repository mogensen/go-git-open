package gitupstreams

import (
	"testing"
)

func TestAzureDevopsUpstream_BranchURL(t *testing.T) {
	type args struct {
		repoURL string
		branch  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "https on master",
			args: args{
				repoURL: "https://ssh.dev.azure.com/v3/CORP/Project/GitRepo",
				branch:  "master",
			},
			want:    "https://dev.azure.com/CORP/Project/_git/GitRepo",
			wantErr: false,
		},
		{
			name: "https on develop",
			args: args{
				repoURL: "https://ssh.dev.azure.com/v3/CORP/Project/GitRepo",
				branch:  "develop",
			},
			want:    "https://dev.azure.com/CORP/Project/_git/GitRepo?version=GBdevelop",
			wantErr: false,
		},
		{
			name: "git on master",
			args: args{
				repoURL: "git@ssh.dev.azure.com:v3/CORP/Project/GitRepo",
				branch:  "master",
			},
			want:    "https://dev.azure.com/CORP/Project/_git/GitRepo",
			wantErr: false,
		},
		{
			name: "git on develop",
			args: args{
				repoURL: "git@ssh.dev.azure.com:v3/CORP/Project/GitRepo",
				branch:  "develop",
			},
			want:    "https://dev.azure.com/CORP/Project/_git/GitRepo?version=GBdevelop",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := getURL(tt.args.repoURL)
			u := AzureDevopsUpstream{}
			got, err := u.BranchURL(url, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("AzureURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AzureURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
