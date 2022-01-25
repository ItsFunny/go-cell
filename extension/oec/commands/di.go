package commands

import "github.com/itsfunny/go-cell/di"

var (
	OecCommands = di.RegisterCommandConstructor(
		newRegisterCommand,
		newDeployContractCmd,
		newRegisterAndDeployCmd,
		newTransferCommand,
		newBalanceCommand,
		newImportAccountCommand,
	)
)
