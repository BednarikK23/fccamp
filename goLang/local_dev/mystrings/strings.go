// by convention, we name our package the same as the directory
package mystrings

// Reverse reverses a string left to right
// Notice that we NEED to capitalize the first letter of the function
// If we don't then we won't be able to access this function outside of the
// mystrings package
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

/*
Note that there is no main.go or func main() in this package.

go build won't build an executable from a library package.
However, go build will still compile the package and save it to our local build cache.
It's useful for checking for compile errors.

-> go build
*/
