package main

import "fmt"

func Square(x int) int {
	return x * x
}

func main() {
	fmt.Println("Square: ", Square(4))
}

/*
$ go test -v
$ go test -v -cover
$ go test -v -cover -coverprofile=c.out
$ go tool cover -html=c.out -o coverage.html

------------------------------------------------------
chaos testing
system testing
load/performance testing
end-to-end testing
integration testing
unit testing
---------------------------------------------------------
Goals:
* extreme values
* input validation
* race condition
* error conditions
* boundary conditions
* pre- and post- conditions
* randomized data (fuzzing)
* configuration & deployment
* interfaces to other software
---------------------------------------------------------
TestHello(t *testing.T){}
Errors are reported through parameter t and fail the test

Table-driven test
Table-driven subtest (t.Run(name f))//closure


---------------------------------------------------------
I can deﬁne a root function for all testing; it will then
run all tests from this point
func TestMain(m *testing.M)
stop, err := startEmulator()
if err != nil {
	log.Println("*** FAILED TO START EMULATOR ***")
	os.Exit(-1)
}
result := m.Run() // run all UTs
stop()
os.Exit(result)
}

---------------------------------------------------------
Special test-only packages:
If you need to add test-only code as part of a package, you can place it in a package that ends in _test


*/
