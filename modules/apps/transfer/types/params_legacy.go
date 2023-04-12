/*
NOTE: Usage of x/params to manage parameters is deprecated in favor of x/gov
controlled execution of MsgUpdateParams messages. These types remains solely
for migration purposes and will be removed in a future release.
[#3621](https://github.com/cosmos/ibc-go/issues/3621)
*/
package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
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

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySendEnabled, &p.SendEnabled, validateEnabledTypeLegacy),
		paramtypes.NewParamSetPair(KeyReceiveEnabled, &p.ReceiveEnabled, validateEnabledTypeLegacy),
		paramtypes.NewParamSetPair(KeySlashPrefix, &p.SlashPrefix, validatePrefixLegacy),
	}
}

// Validate all ibc-transfer module parameters
func (p Params) Validate() error {
	if err := validateEnabledTypeLegacy(p.SendEnabled); err != nil {
		return err
	}

	if err := validatePrefixLegacy(p.SlashPrefix); err != nil {
		return err
	}

	return validateEnabledTypeLegacy(p.ReceiveEnabled)
}

func validateEnabledTypeLegacy(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validatePrefixLegacy(i interface{}) error {
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
