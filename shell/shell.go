package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)

	// error if args is empty
	if len(args) == 0 {
		return nil, errors.New("input cannot be empty")
	}

	return exec.Command(args[0], args[1:]...), nil
}

func Main() int {
	output := WithOutput(os.Stdout)
	input := WithInput(os.Stdin)
	errput := WithErrput(os.Stderr)

	session, err := NewSession(input, output, errput)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	session.Run()
	return 0
}

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	DryRun         bool
}

func (s *Session) Run() {
	fmt.Fprintf(s.Stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)

		if err != nil {
			fmt.Fprintf(s.Stdout, "> ")
			continue
		}

		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s\n> ", line)
			continue
		}

		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Fprintln(s.Stderr, "error: ", err)
		}

		fmt.Fprintf(s.Stdout, "%s", output)
		fmt.Fprintf(s.Stdout, "\n> ")
	}
	fmt.Fprintln(s.Stdout, "\nBe seeing you!")
}

type option func(session *Session) error

func NewSession(opts ...option) (*Session, error) {
	c := &Session{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		DryRun: false,
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithInput(r io.Reader) option {
	return func(s *Session) error {
		if r == nil {
			return errors.New("nil is not a valid reader")
		}
		s.Stdin = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(s *Session) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		s.Stdout = w
		return nil
	}
}

func WithErrput(w io.Writer) option {
	return func(s *Session) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		s.Stderr = w
		return nil
	}
}

func WithDryRun(b bool) option {
	return func(s *Session) error {
		s.DryRun = b
		return nil
	}
}
