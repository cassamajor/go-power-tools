package battery_test

import (
	"github.com/cassamajor/battery"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func Test_ParsePmsetOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/pmset.txt")

	if err != nil {
		t.Fatal(err)
	}

	want := battery.Status{ChargePercent: 100}
	got, err := battery.ParsePmsetOutput(string(data))

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
