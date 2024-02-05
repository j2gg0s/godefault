// Package must be named output_tests rather than _output tests to avoid
// ambiguous import conflict with main gengo package
module github.com/j2gg0s/godefault/examples/output_tests

go 1.21.4

require (
	github.com/google/go-cmp v0.5.9
	github.com/j2gg0s/godefault v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	k8s.io/gengo v0.0.0-20240129211411-f967bbeff4b4 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
)

replace github.com/j2gg0s/godefault => ../../

replace k8s.io/gengo => /Users/j2gg0s/go/src/k8s.io/gengo
