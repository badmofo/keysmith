module github.com/dfinity/keysmith

go 1.16

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2 // indirect
	github.com/dfinity/go-hdkeychain v1.1.0
	github.com/ethereum/go-ethereum v1.10.12
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d
)

replace (
	github.com/tyler-smith/go-bip39 => github.com/tyler-smith/go-bip39 v1.1.0
)
