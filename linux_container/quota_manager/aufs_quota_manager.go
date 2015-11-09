package quota_manager

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/pivotal-golang/lager"
)

type AUFSQuotaManager struct{}

func (AUFSQuotaManager) SetLimits(logger lager.Logger, containerRootFSPath string, limits garden.DiskLimits) error {
	return nil
}

func (AUFSQuotaManager) GetLimits(logger lager.Logger, containerRootFSPath string) (garden.DiskLimits, error) {
	return garden.DiskLimits{}, nil
}

func (AUFSQuotaManager) GetUsage(logger lager.Logger, containerRootFSPath string) (garden.ContainerDiskStat, error) {
	_, err := os.Stat(containerRootFSPath)
	if os.IsNotExist(err) {
		return garden.ContainerDiskStat{}, fmt.Errorf("get usage: %s", err)
	}

	command := fmt.Sprintf("df -B 1 | grep %s | awk -v N=3 '{print $N}'", filepath.Base(containerRootFSPath))
	outbytes, err := exec.Command("sh", "-c", command).CombinedOutput()
	if err != nil {
		return garden.ContainerDiskStat{}, fmt.Errorf("get usage: df: %s, %s", err, string(outbytes))
	}

	var bytesUsed uint64
	if _, err := fmt.Sscanf(string(outbytes), "%d", &bytesUsed); err != nil {
		return garden.ContainerDiskStat{}, fmt.Errorf("get usage: Sscanf: %s", err)
	}

	return garden.ContainerDiskStat{
		ExclusiveBytesUsed: bytesUsed,
	}, nil

	return garden.ContainerDiskStat{}, nil
}

func (AUFSQuotaManager) Setup() error {
	return nil
}

func (AUFSQuotaManager) IsEnabled() bool {
	return false
}
