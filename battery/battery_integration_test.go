//go:build integration

package battery_test

import (
	"bytes"
	"github.com/cassamajor/battery"
	"os/exec"
	"testing"
)

func Test_GetPmsetOutput(t *testing.T) {
	t.Parallel()

	data, err := exec.Command("pmset", "-g", "ps").CombinedOutput()

	if err != nil {
		t.Skipf("skipping test; error running pmset: %v", err)
	}

	if !bytes.Contains(data, []byte("InternalBattery")) {
		t.Skip("skipping test; no battery found")
	}

	text, err := battery.GetPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}

	status, err := battery.ParsePmsetOutput(text)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Battery charge: %d%%", status.ChargePercent)
}
