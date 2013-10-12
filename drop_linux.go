// © 2013 Steve McCoy under the MIT license.

package privs

/*
#define _GNU_SOURCE
#include <unistd.h>

int just_seteuid(uid_t e){
	int err = setresuid(getuid(), e, geteuid());
	if(err != 0 || geteuid() != e)
		return -1;
	return 0;
}
*/
import "C"

import (
	"errors"
)

// Drop is equivalent to calling DropTo with the current real user ID.
func Drop() error {
	return DropTo(int(C.getuid()))
}

// DropTo temporarily drops privileges by moving the privileged UID
// to the saved UID and assigning newID to the effective UID.
func DropTo(newID int) error {
	e := C.just_seteuid(C.uid_t(newID))
	if e != 0 {
		return errors.New("failed to set effective UID")
	}
	return nil
}
