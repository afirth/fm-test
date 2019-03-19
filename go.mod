module github.com/afirth/fm-test

go 1.12

require (
	github.com/afirth/fm-test/api v0.0.2
	github.com/afirth/fm-test/gbdx v0.0.2 // indirect
	github.com/afirth/fm-test/transcode v0.0.2 // indirect
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/gorilla/mux v1.7.0
	github.com/stretchr/testify v1.3.0 // indirect
)

replace github.com/afirth/fm-test/api v0.0.2 => ./pkg/api

replace github.com/afirth/fm-test/transcode v0.0.2 => ./pkg/transcode

replace github.com/afirth/fm-test/gbdx v0.0.2 => ./pkg/gbdx
