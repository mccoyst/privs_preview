// © 2013 Steve McCoy under the MIT license.

package privs

/*
#define _GNU_SOURCE
#include <unistd.h>

int just_restore(){
	uid_t r, e, s;
	if(getresuid(&r, &e, &s) != 0){
		return -1;
	}
	if(setresuid(-1, s, -1) != 0){
		return -1;
	}
	if(geteuid() != s){
		return -1;
	}
	return 0;
}
*/
import "C"

import (
	"errors"
)

// Restore restores privileges by copying the privileged UID from
// the saved UID to the effective UID.
func Restore() error {
	err := C.just_restore()
	if err != 0 {
		return errors.New("failed to reset effective UID")
	}
	return nil
}
