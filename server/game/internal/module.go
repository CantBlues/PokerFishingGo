package internal

import (
	"poker/server/base"
	"poker/server/model"
	"poker/server/protocol"
	"reflect"

	"poker/github/dolotech/leaf/room"

	"poker/github/dolotech/leaf/module"

	"github.com/golang/glog"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func init() {
	handler(&protocol.JoinRoom{}, room.OnMessage)
	handler(&protocol.LeaveRoom{}, room.OnMessage)
	handler(&protocol.Bet{}, room.OnMessage)
	handler(&protocol.SitDown{}, room.OnMessage) //
	handler(&protocol.StandUp{}, room.OnMessage) //
	handler(&protocol.Chat{}, room.OnMessage)    //
}

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	room.Init(&Creator{})
}

func (m *Module) OnDestroy() {
	glog.Errorln("OnDestroy")
}

type Creator struct{}

// 对玩家未进入房间，或者没房间数据的处理
func (this *Creator) Create(m interface{}) room.IRoom {
	if msg, ok := m.(*protocol.JoinRoom); ok {
		if len(msg.RoomNumber) == 0 {
			r := room.FindRoom()
			return r
		}
		r := room.GetRoom(msg.RoomNumber)
		if r != nil {
			return r
		}
		room := NewRoom(9, 5, 10, 1000, model.Timeout)
		room.Insert()
		return room
	}
	return nil
}
