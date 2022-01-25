package commands

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/extension/oec/contract"
)

type oecCommand struct {
	*reactor.Command
	service contract.IContractService
}

func newOecCommand(s contract.IContractService, cmd *reactor.Command) reactor.ICommand {
	ret := &oecCommand{
		Command: cmd,
		service: s,
	}
	return ret
}
