// © 2013 Steve McCoy under the MIT license.

package privs

import (
	"errors"
	"syscall"
)

// Lose permanently drops privileges by assigning newID to
// all of the real, effective, and saved UIDs.
func LoseTo(newID int) error {
	olde := syscall.Geteuid()

	err := syscall.Setreuid(newID, newID)
	if err != nil {
		return err
	}
	if syscall.Getuid() != newID {
		return errors.New("failed to set real UID")
	}
	if syscall.Geteuid() != newID {
		return errors.New("failed to set effective UID")
	}
	savedUID = olde
	// hope that the suid, if it exists, was set
	return nil
}
