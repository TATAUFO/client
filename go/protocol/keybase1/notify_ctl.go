// Auto-generated by avdl-compiler v1.3.20 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_ctl.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type NotificationChannels struct {
	Session      bool `codec:"session" json:"session"`
	Users        bool `codec:"users" json:"users"`
	Kbfs         bool `codec:"kbfs" json:"kbfs"`
	Tracking     bool `codec:"tracking" json:"tracking"`
	Favorites    bool `codec:"favorites" json:"favorites"`
	Paperkeys    bool `codec:"paperkeys" json:"paperkeys"`
	Keyfamily    bool `codec:"keyfamily" json:"keyfamily"`
	Service      bool `codec:"service" json:"service"`
	App          bool `codec:"app" json:"app"`
	Chat         bool `codec:"chat" json:"chat"`
	PGP          bool `codec:"pgp" json:"pgp"`
	Kbfsrequest  bool `codec:"kbfsrequest" json:"kbfsrequest"`
	Badges       bool `codec:"badges" json:"badges"`
	Reachability bool `codec:"reachability" json:"reachability"`
	Team         bool `codec:"team" json:"team"`
}

func (o NotificationChannels) DeepCopy() NotificationChannels {
	return NotificationChannels{
		Session:      o.Session,
		Users:        o.Users,
		Kbfs:         o.Kbfs,
		Tracking:     o.Tracking,
		Favorites:    o.Favorites,
		Paperkeys:    o.Paperkeys,
		Keyfamily:    o.Keyfamily,
		Service:      o.Service,
		App:          o.App,
		Chat:         o.Chat,
		PGP:          o.PGP,
		Kbfsrequest:  o.Kbfsrequest,
		Badges:       o.Badges,
		Reachability: o.Reachability,
		Team:         o.Team,
	}
}

type SetNotificationsArg struct {
	Channels NotificationChannels `codec:"channels" json:"channels"`
}

func (o SetNotificationsArg) DeepCopy() SetNotificationsArg {
	return SetNotificationsArg{
		Channels: o.Channels.DeepCopy(),
	}
}

type NotifyCtlInterface interface {
	SetNotifications(context.Context, NotificationChannels) error
}

func NotifyCtlProtocol(i NotifyCtlInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.notifyCtl",
		Methods: map[string]rpc.ServeHandlerDescription{
			"setNotifications": {
				MakeArg: func() interface{} {
					ret := make([]SetNotificationsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SetNotificationsArg)
					if !ok {
						err = rpc.NewTypeError((*[]SetNotificationsArg)(nil), args)
						return
					}
					err = i.SetNotifications(ctx, (*typedArgs)[0].Channels)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type NotifyCtlClient struct {
	Cli rpc.GenericClient
}

func (c NotifyCtlClient) SetNotifications(ctx context.Context, channels NotificationChannels) (err error) {
	__arg := SetNotificationsArg{Channels: channels}
	err = c.Cli.Call(ctx, "keybase.1.notifyCtl.setNotifications", []interface{}{__arg}, nil)
	return
}
