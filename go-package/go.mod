module test-go-module-package

go 1.25.0
// go 1.25 // works like this too when go version is 1.25.0 but for "go1.19.2 run .", only go 1.25 will work
// go 1.24.4  //cannot work with patch version when doing "go1.19.2 run . but works when doing go run . using go version 1.24.4"
// go 1.24  //works with "go1.19.2 run ."
// go 1.29  //works with "go1.19.2 run ."
// go 1.19.2  // go directive cannot use patch version only major and minor version, gives compilation error, "invalid go version '1.19.2': must match format 1.23"
// go 1.19.0     // go directive cannot use patch version only major and minor version, gives compilation error, "invalid go version '1.19.0': must match format 1.23"

// toolchain go1.20.0  // cannot work with toolChain go1.20.0
// toolchain go1.25.0  // cannot work with toolChain go1.25.0
// toolchain go1.21.0 // works
// toolchain go1.20.0 // Does not work,no toolchain before go1.21 works


require github.com/spf13/cobra v1.9.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)
