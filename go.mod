module github.com/ThePiachu/TestGoPrivate

go 1.22.8

require (
	github.com/ThePiachu/Go v0.0.0-20241015200357-f69786ad4933
	github.com/ThePiachu/gobcy/v2 v2.0.6-0.20241018184526-a9b44f1adb8c
	golang.org/x/net v0.30.0
	google.golang.org/appengine/v2 v2.0.6
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/ThePiachu/TestGoPrivate/Work => ./Work
