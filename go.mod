module github.com/ThePiachu/TestGoPrivate

go 1.22.8

require (
	github.com/ThePiachu/Go v0.0.0-20241015200357-f69786ad4933
	golang.org/x/net v0.30.0
	google.golang.org/appengine/v2 v2.0.6
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/ThePiachu/TestGoPrivate/Work => ./Work
