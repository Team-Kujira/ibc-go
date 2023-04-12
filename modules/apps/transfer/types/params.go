package types

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
	// DefaultReceiveEnabled enabled
	DefaultReceiveEnabled = true
	// DefaultSlashPrefix factory
	DefaultSlashPrefix = "factory"
)

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
