package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name    string
		project string
		exp     string
		expErr  error
	}{
		{
			name:    "success",
			project: "./testdata/tool/",
			exp:     "go build SUCCESS\ngo test OK\ngo format OK\n",
			expErr:  nil,
		},
		{
			name:    "fail",
			project: "./testdata/tool-err/",
			exp:     "",
			expErr: &stepErr{
				step: "go build",
			},
		},
		{
			name:    "fail",
			project: "./testdata/toolFmtErr/",
			exp:     "",
			expErr: &stepErr{
				step: "go fmt",
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			out := new(bytes.Buffer)

			err := run(tC.project, out)

			if tC.expErr != nil {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if !errors.Is(err, tC.expErr) {
					t.Fatalf("expected %v, got %v", tC.expErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got := out.String(); got != tC.exp {
				t.Fatalf("expected %q, got %q", tC.exp, got)
			}
		})
	}
}
