module github.com/manuelbilbao/cli/ignite/pkg/gomodule

go 1.21.0

require (
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.8.2
	github.com/ignite/modules v1.0.0
)

replace github.com/ignite/modules => ../local-module-fork
