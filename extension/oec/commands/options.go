package commands

import "github.com/itsfunny/go-cell/base/core/options"

var (
	moniker       = "moniker"
	monikerOption = options.StringOption(moniker, "moniker", "别名")

	from         = "from"
	to           = "to"
	amount       = "amount"
	fromOption   = options.StringOption(from, from, "from")
	toOption     = options.StringOption(to, to, to)
	amountOption = options.IntOption(amount, amount, amount).WithDefault(1)

	blockNumber       = "blockNumber"
	blockNumberOption = options.Int64Option(blockNumber, blockNumber, blockNumber).WithRequired(false).WithDefault(int64(0))

	prvHex       = "prvHex"
	prvHexOption = options.StringOption(prvHex, prvHex, prvHex).WithRequired(true)
)
