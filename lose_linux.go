// © 2013 Steve McCoy under the MIT license.

package privs

/*
#define _GNU_SOURCE
#include <unistd.h>

int just_setresuid(uid_t e){
	uid_t r, newe, s;
	if(setresuid(e, e, e) != 0)
		return -1;
	if(getresuid(&r, &newe, &s) != 0)
		return -1;
	if(r != e || newe != e || s != e)
		return -1;
	return 0;
}
*/
import "C"

import (
	"errors"
)

// Lose permanently drops privileges by assigning newID to
// all of the real, effective, and saved UIDs.
func LoseTo(newID int) error {
	err := C.just_setresuid(C.uid_t(newID))
	if err != 0 {
		return errors.New("failed to set user IDs")
	}
	return nil
}
