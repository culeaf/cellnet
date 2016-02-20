// Generated by github.com/davyxu/cellnet/protoc-gen-msg
// DO NOT EDIT!
// Source: coredef.proto
package coredef

import (
	"github.com/davyxu/cellnet"
)

func init() {
	cellnet.RegisterMessageMeta("coredef.SessionAccepted", (*SessionAccepted)(nil))
	cellnet.RegisterMessageMeta("coredef.SessionConnected", (*SessionConnected)(nil))
	cellnet.RegisterMessageMeta("coredef.SessionClosed", (*SessionClosed)(nil))
	cellnet.RegisterMessageMeta("coredef.PeerInit", (*PeerInit)(nil))
	cellnet.RegisterMessageMeta("coredef.PeerStart", (*PeerStart)(nil))
	cellnet.RegisterMessageMeta("coredef.PeerStop", (*PeerStop)(nil))
	cellnet.RegisterMessageMeta("coredef.UpstreamACK", (*UpstreamACK)(nil))
	cellnet.RegisterMessageMeta("coredef.CloseClientACK", (*CloseClientACK)(nil))
	cellnet.RegisterMessageMeta("coredef.DownstreamACK", (*DownstreamACK)(nil))
	cellnet.RegisterMessageMeta("coredef.RemoteCallREQ", (*RemoteCallREQ)(nil))
	cellnet.RegisterMessageMeta("coredef.RemoteCallACK", (*RemoteCallACK)(nil))
	cellnet.RegisterMessageMeta("coredef.TestEchoACK", (*TestEchoACK)(nil))
}