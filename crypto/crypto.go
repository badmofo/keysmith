package crypto

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/dfinity/go-hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
)

const DerivationPath = "m/44'/223'/0'"

func LoadXPubKey(xpub string) (*hdkeychain.ExtendedKey, error) {
    xPrivKey, err := hdkeychain.NewKeyFromString(xpub)
    return xPrivKey, err
}

func DeriveMasterXPrivKey(seed []byte) (*hdkeychain.ExtendedKey, error) {
	masterXPrivKey, err := hdkeychain.NewMaster(
		seed,
		&chaincfg.MainNetParams,
	)
	if err != nil {
		return nil, err
	}
	path, err := accounts.ParseDerivationPath(DerivationPath)
	if err != nil {
		return nil, err
	}
	for _, i := range path {
		masterXPrivKey, err = masterXPrivKey.Derive(i)
		if err != nil {
			return nil, err
		}
	}
	return masterXPrivKey, nil
}

func DeriveGrandchildECKeyPair(
	masterXPrivKey *hdkeychain.ExtendedKey,
	i uint32,
) (*btcec.PrivateKey, *btcec.PublicKey, error) {
	// First apply the change.
	childXPrivKey, err := masterXPrivKey.Derive(0)
	if err != nil {
		return nil, nil, err
	}
	grandchildXPrivKey, err := childXPrivKey.Derive(i)
	if err != nil {
		return nil, nil, err
	}
	grandchildECPrivKey, err := grandchildXPrivKey.ECPubKey()
	if err != nil {
		return nil, nil, err
	}
    return nil, grandchildECPrivKey, nil
}
