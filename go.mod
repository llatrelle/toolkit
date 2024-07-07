module github.com/llatrelle/toolkit

go 1.21.5

replace github.com/llatrelle/toolkit => ./toolkit/logger

require (
	github.com/go-chi/chi v1.5.5
	github.com/rs/zerolog v1.33.0
	github.com/stretchr/testify v1.9.0
	github.com/unrolled/render v1.6.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
