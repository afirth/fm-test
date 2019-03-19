module github.com/afirth/fm-test

go 1.12

require (
	github.com/afirth/fm-test/api v0.0.1
	github.com/afirth/fm-test/gbdx v0.0.1 // indirect
	github.com/afirth/fm-test/transcode v0.0.1 // indirect
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/gorilla/mux v1.7.0
)

replace github.com/afirth/fm-test/api v0.0.1 => ./pkg/api

replace github.com/afirth/fm-test/transcode v0.0.1 => ./pkg/transcode

replace github.com/afirth/fm-test/gbdx v0.0.1 => ./pkg/gbdx
