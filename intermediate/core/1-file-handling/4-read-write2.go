/*
go doc os openfile

func OpenFile(name string, flag int, perm FileMode) (*File, error)
    OpenFile is the generalized open call; most users will use Open or Create
    instead. It opens the named file with specified flag (O_RDONLY etc.). If the
    file does not exist, and the O_CREATE flag is passed, it is created with
    mode perm (before umask). If successful, methods on the returned File can be
    used for I/O. If there is an error, it will be of type *PathError.


$ go doc os writestring
package os // import "os"

func (f *File) WriteString(s string) (n int, err error)
    WriteString is like Write, but writes the contents of string s rather than a
    slice of bytes.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.OpenFile("/users/ganesh/temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	/* (perm) Permissions of 600 mean that the owner has full read and write access to the file,
	while no other user can access the file. Permissions of 644 mean that the owner of the
	file has read and write access, while the group members and other users on the system only have read access.*/

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		_, err := file.WriteString("Hope you had a good day!!!")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	defer file.Close()
}
