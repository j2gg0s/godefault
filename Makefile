.PHONY: codegen-defaults
codegen-defaults:
	dlv debug ./cmd/godefault \
		-- \
		--v 6 \
		--go-header-file=${HOME}/go/src/k8s.io/gengo/boilerplate/no-boilerplate.go.txt \
		-i ./apis/user/v1/ \
		-i ./apis/core/v1/ \
		--trim-path-prefix=github.com/j2gg0s/godefault \
		-o ./
