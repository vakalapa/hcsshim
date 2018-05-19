// +build rs5

package uvm

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("HCSSHIM_TEST_DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}

func TestCreateBadOS(t *testing.T) {
	opts := &UVMOptions{
		OperatingSystem: "foobar",
	}
	_, err := Create(opts)
	if err == nil || (err != nil && err.Error() != `unsupported operating system "foobar"`) {
		t.Fatal(err)
	}
}

func TestCreateBadKirdPath(t *testing.T) {
	opts := &UVMOptions{
		OperatingSystem: "linux",
		KirdPath:        `c:\does\not\exist\I\hope`,
	}
	_, err := Create(opts)
	if err == nil || (err != nil && err.Error() != `kernel 'c:\does\not\exist\I\hope\bootx64.efi' not found`) {
		t.Fatal(err)
	}
}

func TestCreateLCOW(t *testing.T) {
	opts := &UVMOptions{
		OperatingSystem: "linux",
	}
	_, err := Create(opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateWCOWBadLayerFolders(t *testing.T) {
	opts := &UVMOptions{
		OperatingSystem: "windows",
	}
	_, err := Create(opts)
	if err == nil || (err != nil && err.Error() != `at least 2 LayerFolders must be supplied`) {
		t.Fatal(err)
	}
}
