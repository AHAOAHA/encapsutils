/*
 * @berief 扩展Golang flag包
 */

package encapsutils

import "fmt"

type StringSlice []string

func (f *StringSlice) String() string {
	return fmt.Sprintf("%v", []string(*f))
}

func (f *StringSlice) Set(value string) error {
	*f = append(*f, value)
	return nil
}
