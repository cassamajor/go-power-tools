package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var pmsetOutput = regexp.MustCompile(`(\d+)%`)

// ParsePmsetOutput parses the output of the `pmset -g pm` command.
func ParsePmsetOutput(text string) (Status, error) {
	matches := pmsetOutput.FindStringSubmatch(text)

	if len(matches) < 2 {
		return Status{}, fmt.Errorf("no charge percent found in output: %q", text)
	}

	charge, err := strconv.Atoi(matches[1])

	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percent: %q with error %w", matches[1], err)
	}

	return Status{ChargePercent: charge}, nil
}

func GetPmsetOutput() (string, error) {
	data, err := exec.Command("pmset", "-g", "ps").CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("error running pmset: %w", err)
	}

	return string(data), nil
}
