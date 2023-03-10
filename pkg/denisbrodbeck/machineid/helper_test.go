package machineid

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"strings"
	"testing"
)

func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func Test_run(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	wantStdout := "hello"
	wantStderr := ""
	if err := run(stdout, stderr, "echo", "hello"); err != nil {
		t.Error(err)
	}
	gotStdout := strings.TrimRight(stdout.String(), "\r\n")
	if gotStdout != wantStdout {
		t.Errorf("run() = %v, want %v", gotStdout, wantStdout)
	}
	if gotStderr := stderr.String(); gotStderr != wantStderr {
		t.Errorf("run() = %v, want %v", gotStderr, wantStderr)
	}
}

func Test_run_unknown(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	err := run(stdout, stderr, "echolo", "hello")
	if err == nil {
		t.Error("expected error, got none")
	}
	if strings.Contains(err.Error(), "executable file not found") == false {
		t.Error("unexpected error, expected exec not found")
	}
}

func Test_trim(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil",
			args: args{s: ""},
			want: "",
		},
		{
			name: "space",
			args: args{s: " space "},
			want: "space",
		},
		{
			name: "nl",
			args: args{s: "data\n"},
			want: "data",
		},
		{
			name: "combined",
			args: args{s: " some data \n"},
			want: "some data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trim(tt.args.s); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
