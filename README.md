DEPRECATED
==========
go-extimport was originally developed to ensure the consistency of the vendor directory for projects. This was necessary
when the vendor directory was managed manually or using basic tools such as [govendor[(https://github.com/kardianos/govendor).
However, tools such as [dep](https://github.com/golang/dep) and the support for [modules in Go](https://blog.golang.org/using-go-modules) 
have made the functionality provided by this tool unnecessary. As such, active development on this project has ended.

go-extimport
============
go-extimport is a program that verifies that there are no imports that reference external packages in a Go project.

A package is considered external if it is not in the standard Go library and not resolvable within the project directory
itself (the package is not in the project or vendored in the project). An import is considered to be directly external
if it imports an external package. An import is considered transitively external if the imported package itself is not
external, but one of its dependent packages is external.

Given a package, go-extimport checks the imports of all of the Go files and Go test files in that package. However, when
checking transitive external package dependencies, only non-test go files are considered (that is, the check will not
fail if a test file of an imported package has an external dependency).
