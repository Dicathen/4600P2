package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
		{
			name: "help command",
			args: args{
				r: strings.NewReader("help\n"),
			},
		},
		{
			name: "history command",
			args: args{
				r: strings.NewReader("history\n"),
			},
		},
		{
			name: "kill command",
			args: args{
				r: strings.NewReader("kill \n"),
			},
		},
		{
			name: "pwd command",
			args: args{
				r: strings.NewReader("pwd\n"),
			},
		},
		{
			name: "echo command",
			args: args{
				r: strings.NewReader("echo hello world\n"),
			},
		},
		{
			name: "echo command with *",
			args: args{
				r: strings.NewReader("echo *\n"),
			},
		},
		{
			name: "cd command",
			args: args{
				r: strings.NewReader("cd\n"),
			},
		},
		{
			name: "env command",
			args: args{
				r: strings.NewReader("env\n"),
			},
		},
		{
			name: "command in bin path command",
			args: args{
				r: strings.NewReader("test\n"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			}
		})
	}
}
