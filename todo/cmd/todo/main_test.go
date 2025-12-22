package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	filename = ".todo.json"
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if out, err := build.CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, "Build failed:\n", string(out))
		os.Exit(1)
	}

	code := m.Run()

	os.Remove(binName)
	os.Remove(filename)
	os.Exit(code)
}

func TestTodoCli(t *testing.T) {
	task := "test task 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	run := func(t *testing.T, name string, args ...string) string {
		t.Helper()
		cmd := exec.Command(cmdPath, args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%s failed: %v\noutput:\n%s", name, err, out)
		}
		return string(out)
	}

	t.Run("Add New Task", func(t *testing.T) {
		run(t, "add", "-add", task)
	})

	t.Run("Add New Task via STDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		pipe, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}

		io.WriteString(pipe, "task2")
		pipe.Close()

		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("stdin add failed: %v\noutput:\n%s", err, out)
		}
	})

	t.Run("List Tasks", func(t *testing.T) {
		out := run(t, "list", "-list")

		expected :=
			fmt.Sprintf("1 [ ] %s\n", task) +
				fmt.Sprintf("2 [ ] task2")

		if strings.TrimSpace(out) != expected {
			t.Fatalf("list mismatch:\nexpected:\n%s\ngot:\n%s",
				expected, strings.TrimSpace(out))
		}
	})

	t.Run("Delete the added task", func(t *testing.T) {
		run(t, "delete", "-delete", "2")
	})

	t.Run("Complete Task", func(t *testing.T) {
		run(t, "add dummy", "-add", "dummy")
		run(t, "complete", "-complete", "2")

		out := run(t, "list after complete", "-list")

		expected :=
			fmt.Sprintf("1 [ ] %s\n", task) +
				"2 [x] dummy"

		if strings.TrimSpace(out) != expected {
			t.Fatalf("complete mismatch:\nexpected:\n%s\ngot:\n%s",
				expected, strings.TrimSpace(out))
		}
	})
}
