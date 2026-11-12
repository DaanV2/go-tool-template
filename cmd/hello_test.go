package cmd_test

import (
	"bytes"
	"testing"

	"github.com/DaanV2/go-tool-template/cmd"
)

func TestHello(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{args: []string{"hello"}, want: "Hello, world!\n"},
		{args: []string{"hello", "gopher"}, want: "Hello, gopher!\n"},
	}

	for _, tt := range tests {
		var out bytes.Buffer

		root := cmd.NewRootCmd()
		root.SetOut(&out)
		root.SetArgs(tt.args)

		if err := root.Execute(); err != nil {
			t.Fatalf("Execute(%v) error = %v", tt.args, err)
		}

		if got := out.String(); got != tt.want {
			t.Errorf("Execute(%v) = %q, want %q", tt.args, got, tt.want)
		}
	}
}
