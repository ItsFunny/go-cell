package commands

import "github.com/itsfunny/go-cell/base/core/options"

var (
	moniker     = "moniker"
	from        = "from"
	to          = "to"
	amount      = "amount"
	blockNumber = "blockNumber"
	prvHex      = "prvHex"

	toLimitCount = "toLimitCount"

	transactionLimit = "transactionLimit"
	accountLimit     = "accountLimit"

	hexBlockHash = "hexBlockHash"
)
var (
	monikerOption      = options.StringOption(moniker, "moniker", "别名")
	fromOption         = options.StringOption(from, from, "from")
	toOption           = options.StringOption(to, to, to)
	amountOption       = options.IntOption(amount, amount, amount).WithDefault(1)
	blockNumberOption  = options.Int64Option(blockNumber, blockNumber, blockNumber).WithRequired(false).WithDefault(int64(0))
	prvHexOption       = options.StringOption(prvHex, prvHex, prvHex).WithRequired(true)
	toLimitCountOption = options.IntOption(toLimitCount).WithDefault(10)

	transactionLimitOption = options.IntOption(transactionLimit, transactionLimit, "交易总数限制").WithDefault(10)
	accountLimitOption     = options.IntOption(accountLimit, accountLimit, "账户数限制")

	hexBlockHashOption = options.StringOption(hexBlockHash, hexBlockHash, "block的hex hash")
)
