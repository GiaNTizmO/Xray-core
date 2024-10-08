package all

import (
	"github.com/giantizmo/xray-core/main/commands/all/api"
	"github.com/giantizmo/xray-core/main/commands/all/convert"
	"github.com/giantizmo/xray-core/main/commands/all/tls"
	"github.com/giantizmo/xray-core/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
