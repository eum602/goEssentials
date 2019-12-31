package hello

import "rsc.io/quote" //extracts many packages from this official dependency

func Hello() string {
    return quote.Hello()
}

/*
1. go help
2. go env
3. go get -d github.com/... (pulls a repository and automatically locates this on $GOPATH/src/github folder)
4. go version
5. go fmt ./.. (allows for formatting files in a whole folder from where you run this command)
6. go fmt => it is to format a file
7. go run filename.go => runs a program
8. go run *.go => runs all  go files in a folder you are located.
9. go build filename.go => creates an executable filename => we can run that executable by using ./main
10. go install filename.go => creates an INSTALLER which is located in $GOPATH/bin => that can be run by simply filename (you must be in $GOPATH/bin or reference to that before running it)
11. go install *.go => creates installers for all the go files that exist on the current folder you are located.
12. go install is also for a package but the executable is located in $GOPATH/pkg
13. go mod init eum602.com/repo (to create a go module)
    go mod tidy ==> gets all required depedencies
    go mod why  ==>  gets all required depedencies and also lists those dependencies
14. go test (run tests => file somename_test.go must exist with some logic before runnign the test => after running it creates a go.mod file which indicates direct and indirect dependencies)
15. go list -m all (shows all direct and indirect dependencies required)
16. go get the_dependency_to_pull (eg go get golang .org/x/get)
17. go list -m -versions the_package_to_check_the_versions (eg. rsc.io/sampler) => this is to see all versions
18. go get  the_dependency_to_pull@va.b.c (eg. go get rsc.io/sampler@v1.3.1)
*/
