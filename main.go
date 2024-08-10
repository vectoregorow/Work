/**
* Environment: Go 1.22.2
*
* To install a package, you need to run `go get <package_name>` in the shell.
* This should update the package and dependencies in your `go.mod` and `go.sum` files.
* 
* You can execute tests using the Test command in the "Run" button.
* You can add more commands via ci-config.json file.
*/

package main

import (
	"fmt"
  "ci/greeter"
)

func main() {
	ciGreeter := greeter.New()

	fmt.Println(ciGreeter.Greet());
}
