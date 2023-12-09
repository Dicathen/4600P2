package builtins_test

import (
	"os"
	"testing"

	"github.com/Dicathen/CSCE4600/Project2/builtins"
)

func TestEcho(t *testing.T) {
	type args struct {
		args []string
	}
	//build tests
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "no args should print newline",
			wantErr: nil,
		},
		{
			name:    "* arg should print files in current directory",
			args:    args{args: []string{"*"}},
			wantErr: nil,
		},
		{
			name:    "args should print args",
			args:    args{args: []string{"abc", "def"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//testing
			//change directory to temp directory
			tmp := t.TempDir()
			if err := os.Chdir(tmp); err != nil {
				t.Fatalf("failed to change directory: %v", err)
			}
			//create a temp file
			f, err := os.Create("test.txt")
			if err != nil {
				t.Fatalf("failed to create file: %v", err)
			}
			//close the file
			if err := f.Close(); err != nil {
				t.Fatalf("failed to close file: %v", err)
			}

			if err := builtins.Echo(tt.args.args...); tt.wantErr != nil {
				if tt.wantErr == nil {
					t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
		})
	}
}
