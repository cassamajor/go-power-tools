package shell_test

import (
	"bytes"
	"github.com/cassamajor/shell"
	"github.com/google/go-cmp/cmp"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_CmdFromString(t *testing.T) {
	t.Run("Creates Expected Command", func(t *testing.T) {
		t.Parallel()
		input := "/bin/ls -l main.go"
		cmd, err := shell.CmdFromString(input)

		if err != nil {
			t.Fatal(err)
		}

		want := []string{"/bin/ls", "-l", "main.go"}
		got := cmd.Args

		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	})
}

func Test_CmdFromString_Error(t *testing.T) {
	t.Run("Errors on Empty Input", func(t *testing.T) {
		t.Parallel()
		_, err := shell.CmdFromString("")

		if err == nil {
			t.Error("Expected error on empty input")
		}
	})
}

func Test_NewSession(t *testing.T) {
	t.Parallel()
	want := shell.Session{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	got, err := shell.NewSession()

	if err != nil {
		t.Fatalf("%v", err)
	}

	if want != *got {
		t.Errorf("want %#v, got %#v", want, *got)
	}

	//if !cmp.Equal(want, got) {
	//	t.Error(cmp.Diff(want, got))
	//}
}

func Test_Run(t *testing.T) {
	t.Run("Run produces expected output", func(t *testing.T) {
		t.Parallel()

		input := strings.NewReader("echo hello\n\n")
		output := new(bytes.Buffer)
		errput := io.Discard

		session, err := shell.NewSession(
			shell.WithInput(input),
			shell.WithOutput(output),
			shell.WithErrput(errput),
		)

		if err != nil {
			t.Fatal(err)
		}

		session.Run()
		want := "> hello\n> > \nBe seeing you!\n"
		got := output.String()
		if !cmp.Equal(want, got) {
			t.Errorf(cmp.Diff(want, got))
		}
	})
}

func Test_RunAgain(t *testing.T) {
	t.Run("Run produces expected output", func(t *testing.T) {
		t.Parallel()

		stdin := strings.NewReader("echo hello\n\n")
		stdout := new(bytes.Buffer)
		stderr := io.Discard

		session, err := shell.NewSession(
			shell.WithInput(stdin),
			shell.WithOutput(stdout),
			shell.WithErrput(stderr),
			shell.WithDryRun(true),
		)

		if err != nil {
			t.Fatal(err)
		}

		session.Run()
		want := "> echo hello\n> > \nBe seeing you!\n"
		got := stdout.String()
		if !cmp.Equal(want, got) {
			t.Errorf(cmp.Diff(want, got))
		}
	})
}
