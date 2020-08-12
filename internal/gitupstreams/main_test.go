package gitupstreams

import (
	"reflect"
	"testing"
)

func TestGetBrowerURL(t *testing.T) {
	tests := []struct {
		name            string
		branch          string
		remoteURL       string
		domainOverwrite string
		want            string
		wantErr         bool
	}{
		{
			name:      "gh: basic",
			branch:    "master",
			remoteURL: "git@github.com:user/repo.git",
			want:      "https://github.com/user/repo",
			wantErr:   false,
		},
		{
			name:      "gh: basic with branch",
			branch:    "develop",
			remoteURL: "git@github.com:user/repo.git",
			want:      "https://github.com/user/repo/tree/develop",
			wantErr:   false,
		},
		{
			name:      "gh: basic http",
			branch:    "master",
			remoteURL: "http://github.com/user/repo.git",
			want:      "http://github.com/user/repo",
			wantErr:   false,
		},
		{
			name:            "gh: basic with domain overwrite",
			branch:          "master",
			remoteURL:       "git@github.com:user/repo.git",
			domainOverwrite: "repo.git.com",
			want:            "https://repo.git.com/user/repo",
			wantErr:         false,
		},
		{
			name:      "azure devops: basic",
			branch:    "master",
			remoteURL: "https://CORP@dev.azure.com/v3/CORP/Project/GitRepo",
			want:      "https://dev.azure.com/CORP/Project/_git/GitRepo",
			wantErr:   false,
		},
		{
			name:      "azure devops: ssh",
			branch:    "master",
			remoteURL: "git@ssh.dev.azure.com:v3/CORP/Project/GitRepo",
			want:      "https://dev.azure.com/CORP/Project/_git/GitRepo",
			wantErr:   false,
		},

		{
			name:      "bitbucket: basic",
			branch:    "master",
			remoteURL: "https://bitbucket.org/User/GitRepo",
			want:      "https://bitbucket.org/User/GitRepo",
			wantErr:   false,
		},
		{
			name:      "bitbucket: ssh",
			branch:    "master",
			remoteURL: "git@bitbucket.org:/User/GitRepo.git",
			want:      "https://bitbucket.org/User/GitRepo",
			wantErr:   false,
		},
		{
			name:      "bitbucket: ssh with branch",
			branch:    "develop",
			remoteURL: "git@bitbucket.org:/User/GitRepo.git",
			want:      "https://bitbucket.org/User/GitRepo/src/HEAD/?at=develop",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guh := NewGitURLHandler()
			got, err := guh.GetBrowerURL(tt.remoteURL, tt.domainOverwrite, tt.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBrowerURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBrowerURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
