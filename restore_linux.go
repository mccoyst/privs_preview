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

// I currently need to track it myself,
// because there is no syscall.Getresuid or syscall.Getsuid,
// nor do they exist at all on OSX.
var savedUID int

// Restore restores privileges by copying the privileged UID from
// the saved UID to the effective UID.
func Restore() error {
	err := C.seteuid(savedUID)
	if err != nil {
		return err
	}
	if C.geteuid() != savedUID {
		return errors.New("failed to reset effective UID")
	}
	return nil
}
