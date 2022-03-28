package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dfinity/keysmith/account"
	"github.com/dfinity/keysmith/crypto"
//	"github.com/dfinity/keysmith/seed"
)

const ACCOUNT_CMD = "account"

type AccountCmd struct {
	FlagSet *flag.FlagSet
	Args    *AccountCmdArgs
}

type AccountCmdArgs struct {
	XPub  *string
	Index     *uint
	Protected *bool
}

func NewAccountCmd() *AccountCmd {
	fset := flag.NewFlagSet(ACCOUNT_CMD, flag.ExitOnError)
	args := &AccountCmdArgs{
		XPub:  fset.String("x", "", "XPub"),
		Index:     fset.Uint("i", 0, "Derivation index."),
		Protected: fset.Bool("p", false, "Password protection."),
	}
	return &AccountCmd{fset, args}
}

func (cmd *AccountCmd) Run() error {
	cmd.FlagSet.Parse(os.Args[2:])
    /*
	seed, err := seed.Load(*cmd.Args.SeedFile, *cmd.Args.Protected)
	if err != nil {
		return err
	}
	masterXPrivKey, err := crypto.DeriveMasterXPrivKey(seed)
	if err != nil {
		return err
	}
    */
    masterXPubKey, err := crypto.LoadXPubKey(*cmd.Args.XPub)
	_, grandchildECPubKey, err := crypto.DeriveGrandchildECKeyPair(
		masterXPubKey,
		uint32(*cmd.Args.Index),
	)
	if err != nil {
		return err
	}
	accountId, err := account.FromECPubKey(grandchildECPubKey)
	if err != nil {
		return err
	}
	fmt.Println(accountId.String())
	return nil
}
