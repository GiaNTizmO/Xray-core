package udp

import (
	"github.com/giantizmo/xray-core/common"
	"github.com/giantizmo/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
