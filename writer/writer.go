package writer

import (
	"fmt"
	"os"
)

func WriteToFile(path string, data []byte) error {
	perm := os.FileMode(0o600)

	err := os.WriteFile(path, data, perm)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return os.Chmod(path, perm)
}
