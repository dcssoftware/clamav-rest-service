package claimav

import (
	"os"

	"github.com/DigitalCatServices-io/claimav-wrapper/src/helper/executer"
)

func (s *ClamAVService) ScanFile(fileContent []byte) error {

	tmpFile, tmpFileErr := os.CreateTemp("", "clamav-filescanner")
	if tmpFileErr != nil {
		return tmpFileErr
	}

	defer os.Remove(tmpFile.Name())

	tmpFile.Chmod(0755)

	_, writeErr := tmpFile.Write(fileContent)
	if writeErr != nil {
		return writeErr
	}

	return executer.ExecuteClamAVScanFile(tmpFile.Name())
}
