/*
 * @berief 扩展Golang flag包
 */

package encapsutils

import "fmt"

// StringSlice flag包扩展数组格式
type StringSlice []string

// String flag string.
func (f *StringSlice) String() string {
	return fmt.Sprintf("%v", []string(*f))
}

// Set flag set.
func (f *StringSlice) Set(value string) error {
	*f = append(*f, value)
	return nil
}
