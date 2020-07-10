/* 
Package mylib is my own package
*/
package mylib

// Average returns the average of a series of numbers
func Average(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return int(total/len(s))
}