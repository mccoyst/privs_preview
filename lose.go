// © Steve McCoy under the MIT license.

package privs

import (
	"syscall"
)

// Lose is equivalent to calling LoseTo with the current real user ID.
func Lose() error {
	return LoseTo(syscall.Getuid())
}
