package builtins_test

import (
	"errors"
	"testing"

	"github.com/Dicathen/CSCE4600/Project2/builtins"
)

func TestKill(t *testing.T) {
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
			name:    "error too many args",
			args:    args{args: []string{"abc", "def"}},
			wantErr: errors.New("kill: expected one argument"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//testing
			if err := builtins.Kill(tt.args.args...); tt.wantErr != nil {
				if tt.wantErr == nil {
					t.Errorf("Kill() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
		})
	}
}
