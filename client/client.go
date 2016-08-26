package main

import (
	// "encoding/binary"
	"flag"
	"github.com/golang/glog"
	// "fmt"
	// "encoding/binary"
	// "github.com/golang/protobuf/proto"
	"github.com/oikomi/FishChatServer2/libnet"
	// mybinary "github.com/oikomi/FishChatServer2/libnet/binary"
	// "fmt"
	// myproto "github.com/oikomi/FishChatServer2/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/oikomi/FishChatServer2/codec"
	"github.com/oikomi/FishChatServer2/protocol"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func checkErr(err error) {
	if err != nil {
		glog.Error(err)
	}
}

func clientLoop(session *libnet.Session) {
	err := session.Send(&protocol.Cmd{
		Cmd:    proto.Int32(1001),
		CmdStr: proto.String("hello"),
	})
	checkErr(err)

	// rsp, err := session.Receive()
	// checkErr(err)
	// glog.Info(rsp)

}

func main() {
	var addr string

	flag.StringVar(&addr, "addr", "127.0.0.1:17000", "echo server address")
	flag.Parse()
	protobuf := codec.Protobuf()
	// session, err := libnet.Connect("tcp", addr, libnet.Packet(2, 1024*1024, 1024, binary.BigEndian, TestCodec{}))
	client, err := libnet.Connect("tcp", addr, protobuf, 0)
	checkErr(err)
	clientLoop(client)

	// go func() {
	// 	var msg string
	// 	for {
	// 		if err := session.Receive(&msg); err != nil {
	// 			glog.Error("session.Receive error: ", err)
	// 			break
	// 		}
	// 		fmt.Printf("%s\n", msg)
	// 	}
	// }()
}