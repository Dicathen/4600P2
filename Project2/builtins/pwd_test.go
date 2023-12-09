package builtins_test

import (
	"errors"
	"os"
	"testing"

	"github.com/Dicathen/CSCE4600/Project2/builtins"
)

func TestPrintWorkingDirectory(t *testing.T) {
	tmp := t.TempDir()
	type args struct {
		args []string
	}
	//build tests
	tests := []struct {
		name    string
		args    args
		wantDir string
		wantErr error
	}{
		{
			name:    "error too many args",
			args:    args{args: []string{"abc", "def"}},
			wantErr: errors.New("pwd: expected zero arguments"),
		},
		{
			name:    "no args should print working directory",
			wantDir: tmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//setup
			if err := os.Chdir(tmp); err != nil {
				t.Fatalf("failed to change directory: %v", err)
			}
			//testing
			if err := builtins.PrintWorkingDirectory(tt.args.args...); err != nil {
				if tt.wantErr == nil {
					t.Errorf("PrintWorkingDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			//if something goes wrong
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("failed to get working directory: %v", err)
			}
			if wd != tt.wantDir {
				t.Errorf("PrintWorkingDirectory() = %v, want %v", wd, tt.wantDir)
			}
		})
	}
}
