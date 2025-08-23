package executer

import (
	"os/exec"

	"github.com/DigitalCatServices-io/claimav-wrapper/src/configuration"
)

func ExecuteClamAVScanFile(filename string) error {
	return exec.Command(
		configuration.CLAMAV.SCAN_EXECUTION_PATH,
		filename,
	).Run()
}
