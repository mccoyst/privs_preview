// © 2013 Steve McCoy under the MIT license.

package privs

/*
#include <sys/types.h>
#include <unistd.h>
*/
import "C"

import (
	"errors"
)

// Drop is equivalent to calling DropTo with the current real user ID.
func Drop() error {
	return DropTo(syscall.Getuid())
}

// DropTo temporarily drops privileges by moving the privileged UID
// to the saved UID and assigning newID to the effective UID.
func DropTo(newID int) error {
	olde := C.seteuid()

	err := C.setreuid(syscall.Getuid(), olde)
	if err != nil {
		return err
	}
	err = C.seteuid(newID)
	if err != nil {
		return err
	}
	if C.geteuid() != newID {
		return errors.New("failed to set effective UID")
	}
	savedUID = olde
	return nil
}
