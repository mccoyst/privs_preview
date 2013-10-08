// © Steve McCoy under the MIT license.

package privs

import (
	"errors"
	"syscall"
)

// Drop is equivalent to calling DropTo with the current real user ID.
func Drop() error {
	return DropTo(syscall.Getuid())
}

// DropTo temporarily drops privileges by moving the privileged UID
// to the saved UID and assigning newID to the effective UID.
func DropTo(newID int) error {
	olde := syscall.Geteuid()

	err := syscall.Setreuid(syscall.Getuid(), olde)
	if err != nil {
		return err
	}
	err = syscall.Seteuid(newID)
	if err != nil {
		return err
	}
	if syscall.Geteuid() != newID {
		return errors.New("failed to set effective UID")
	}
	savedUID = olde
	return nil
}
