package command

import (
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/spf13/cobra"
)

var (
	_ reactor.ICommand = (*CLICommand)(nil)
)

type CLICommand struct {
	cobra.Command
}

func (C *CLICommand) ID() reactor.ProtocolID {
	//TODO implement me
	panic("implement me")
}

func (C *CLICommand) Execute(ctx reactor.IBuzzContext) {
	//TODO implement me
	panic("implement me")
}

func (C *CLICommand) SupportRunType() reactor.RunType {
	//TODO implement me
	panic("implement me")
}

func (C *CLICommand) ToSwaggerPath() *reactor.PathItemWrapper {
	//TODO implement me
	panic("implement me")
}

func (C *CLICommand) GetOptions() []options.Option {
	//TODO implement me
	panic("implement me")
}
