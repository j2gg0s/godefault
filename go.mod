module github.com/j2gg0s/kserver

go 1.21.4

require (
	github.com/spf13/pflag v1.0.5
	k8s.io/gengo v0.0.0-20240129211411-f967bbeff4b4
	k8s.io/klog v1.0.0
)

require (
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
)

replace k8s.io/gengo => ./vendor/k8s.io/gengo
