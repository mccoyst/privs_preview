// © 2013 Steve McCoy under the MIT license.

/*
Package privs provides simpler functions to manipulate
"set-user-ID" privileges than provided by POSIX.

These functions are implemented with advice from "Setuid Demystified":
http://www.cs.berkeley.edu/~daw/papers/setuid-usenix02.pdf
*/
package privs

//  BUG(mccoyst): The linux setuid system calls only affect the current thread; not the whole process. https://code.google.com/p/go/issues/detail?id=1435
