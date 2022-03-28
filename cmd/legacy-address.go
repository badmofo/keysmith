package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dfinity/keysmith/crypto"
//	"github.com/dfinity/keysmith/seed"
	eth "github.com/ethereum/go-ethereum/crypto"
)

const LEGACY_ADDRESS_CMD = "legacy-address"

type LegacyAddressCmd struct {
	FlagSet *flag.FlagSet
	Args    *LegacyAddressCmdArgs
}

type LegacyAddressCmdArgs struct {
	XPub  *string
	Index     *uint
	Protected *bool
}

func NewLegacyAddressCmd() *LegacyAddressCmd {
	fset := flag.NewFlagSet(LEGACY_ADDRESS_CMD, flag.ExitOnError)
	args := &LegacyAddressCmdArgs{
		XPub:  fset.String("x", "", "XPub."),
		Index:     fset.Uint("i", 0, "Derivation index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &LegacyAddressCmd{fset, args}
}

func (cmd *LegacyAddressCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
	//seed, err := seed.Load(*cmd.Args.XPub, *cmd.Args.Protected)
	masterXPubKey, err := crypto.LoadXPubKey(*cmd.Args.XPub)
	if err != nil {
		return err
	}
	_, grandchildECPubKey, err := crypto.DeriveGrandchildECKeyPair(
		masterXPubKey,
		uint32(*cmd.Args.Index),
	)
	if err != nil {
		return err
	}
	address := eth.PubkeyToAddress(*grandchildECPubKey.ToECDSA())
	output := strings.ToLower(strings.TrimPrefix(address.String(), "0x"))
	fmt.Println(output)
	return nil
}
