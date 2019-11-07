// ## Exported vs Unexported Structs
// We first declare two struct types, `secretAgent`
// and `PublicEmployee` under the `model` package.

package model

type secretAgent struct { // First letter lower case: Unexported!
	name   string
	salary int
}

type PublicEmployee struct { // First letter upper case: Exported!
	Name   string // exported
	Role   string // exported
	salary int    // unexported (lower case!)
}
