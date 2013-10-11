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

// Lose permanently drops privileges by assigning newID to
// all of the real, effective, and saved UIDs.
func LoseTo(newID int) error {
	olde := C.geteuid()

	err := C.setreuid(newID, newID)
	if err != nil {
		return err
	}
	if C.getuid() != newID {
		return errors.New("failed to set real UID")
	}
	if C.geteuid() != newID {
		return errors.New("failed to set effective UID")
	}
	savedUID = olde
	// hope that the suid, if it exists, was set
	return nil
}
