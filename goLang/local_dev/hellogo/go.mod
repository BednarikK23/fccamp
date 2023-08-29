module github.com/BednarikK23/fccamp/tree/master/goLang/local_dev/hellogo

go 1.20

// go¨s dependency managment is heavilly based on git and remote url's
// so normally what would you do would be push mystring directory into github and import it from there
// what we are doing here with replace and require?:
// little hack to get things working locally without hoving to pubvblish to git...
// we re saying: I want you to take this string, this path (github/...)
// and dont goo look for it on the internet instead just resolve it on this path ../mystrings
replace github.com/BednarikK23/fccamp/tree/master/goLang/local_dev/mystrings v0.0.0 => "../mystrings"

require (
	github.com/BednarikK23/fccamp/tree/master/goLang/local_dev/mystrings v0.0.0
)
