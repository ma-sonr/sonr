package sonr

import (
	"bufio"
	"fmt"

	"io/ioutil"

	"github.com/libp2p/go-libp2p-core/network"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Auth Stream Struct ^ //
type authStreamConn struct {
	stream   network.Stream
	callback Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		stream:   stream,
		callback: sn.Callback,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		stream:   stream,
		callback: sn.Callback,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Write Message on Stream ^ //
func (asc *authStreamConn) Write(authMsg *pb.AuthMessage) error {
	// Initialize
	writer := bufio.NewWriter(asc.stream)
	fmt.Println("Auth Msg Struct: ", authMsg)
	marshaledBytes, err := proto.Marshal(authMsg)
	if err != nil {
		fmt.Println("Error Marshalling bytes: ", err)
	}

	bytesWritten, err := writer.Write(marshaledBytes)
	if err != nil {
		fmt.Println("Write() returned err: ", err)
		return err
	}
	fmt.Println("Number of bytes written: ", bytesWritten)
	writer.Flush()
	return nil
}

// ^ Read Data from Msgio ^ //
func (asc *authStreamConn) Read() {
	bufReader := bufio.NewReader(asc.stream)
	for {
		// ** Read the Buffer **
		readBytes, err := ioutil.ReadAll(bufReader)
		if err != nil {
			fmt.Println("Error reading msg: ", err)
		}
		fmt.Println("Read Bytes: ", readBytes)
		authMsg := pb.AuthMessage{}
		err = proto.Unmarshal(readBytes, &authMsg)
		if err != nil {
			fmt.Println("Error unmarshaling msg: ", err)
		}

		// ** Contains Data **
		// Construct message
		if authMsg.Subject != pb.AuthMessage_NONE {
			// Check Message Subject
			switch authMsg.Subject {
			// @ Request to Invite
			case pb.AuthMessage_REQUEST:
				// Retreive Values
				data, err := proto.Marshal(&authMsg)
				if err != nil {
					fmt.Println("Error Marshaling RefreshMessage ", err)
				}

				// Callback the Invitation
				asc.callback.OnInvited(data)

			// @ Peer Accepted Response to Invite
			case pb.AuthMessage_ACCEPT:
				fmt.Println("Auth Accepted")
				// Callback to Proxies
				asc.callback.OnResponded(true)

			// @ Peer Accepted Response to Invite
			case pb.AuthMessage_DECLINE:
				fmt.Println("Auth Declined")
				// Callback to Proxies
				asc.callback.OnResponded(false)
			}
		}
		// ! Invalid Subject
		fmt.Println("Not a subject", authMsg.Subject)
	}
}
