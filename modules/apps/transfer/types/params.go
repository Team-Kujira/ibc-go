package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
	// DefaultReceiveEnabled enabled
	DefaultReceiveEnabled = true
	// DefaultSlashPrefix factory
	DefaultSlashPrefix = "factory"
)

var (
	// KeySendEnabled is store's key for SendEnabled Params
	KeySendEnabled = []byte("SendEnabled")
	// KeyReceiveEnabled is store's key for ReceiveEnabled Params
	KeyReceiveEnabled = []byte("ReceiveEnabled")
	// KeySlashPrefix is store's key for SlashPrefix Params
	KeySlashPrefix = []byte("SlashPrefix")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the ibc transfer module
func NewParams(enableSend, enableReceive bool, slashPrefix string) Params {
	return Params{
		SendEnabled:    enableSend,
		ReceiveEnabled: enableReceive,
		SlashPrefix:    slashPrefix,
	}
}

// DefaultParams is the default parameter configuration for the ibc-transfer module
func DefaultParams() Params {
	return NewParams(DefaultSendEnabled, DefaultReceiveEnabled, DefaultSlashPrefix)
}

// Validate all ibc-transfer module parameters
func (p Params) Validate() error {
	if err := validateEnabled(p.SendEnabled); err != nil {
		return err
	}

	if err := validatePrefix(p.SlashPrefix); err != nil {
		return err
	}

	return validateEnabled(p.ReceiveEnabled)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySendEnabled, p.SendEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeyReceiveEnabled, p.ReceiveEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeySlashPrefix, p.SlashPrefix, validatePrefix),
	}
}

func validateEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validatePrefix(i interface{}) error {
	p, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	err := sdk.ValidateDenom(p)
	if err != nil {
		return err
	}

	if strings.Contains(p, "/") {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
