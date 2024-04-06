package errutils

import "fmt"

func Wrap(caller string, err error) error {
	fmt.Printf("Error: %s: %s", caller, err)
	return err
}
