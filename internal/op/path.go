package op

import (
	"github.com/alist-org/alist/v3/internal/errs"
	"strings"

	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// GetStorageAndActualPath Get the corresponding storage and actual path
// for path: remove the mount path prefix and join the actual root folder if exists
func GetStorageAndActualPath(rawPath string) (storage driver.Driver, actualPath string, err error) {
	rawPath = utils.FixAndCleanPath(rawPath)
	storage = GetBalancedStorage(rawPath)
	if storage == nil {
		if rawPath == "/" {
			err = errs.NewErr(errs.StorageNotFound, "please add a storage first")
			return
		}
		err = errs.NewErr(errs.StorageNotFound, "rawPath: %s", rawPath)
		return
	}
	log.Debugln("use storage: ", storage.GetStorage().MountPath)

	actualPath = GetActualWithStorage(rawPath, storage)
	return
}

// GetActualWithStorage balance cache or main will return the same, do better
func GetActualWithStorage(rawPath string, storage driver.Driver) string {
	mountPath := utils.GetActualMountPath(storage.GetStorage().MountPath)
	return utils.FixAndCleanPath(strings.TrimPrefix(rawPath, mountPath))
}
