package sonr

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
)

// ^ authStreamMessage is for Auth Stream Request ^
type authStreamMessage struct {
	Subject  string        `json:"Subject"`
	Decision bool          `json:"Decision"`
	PeerInfo lobby.Peer    `json:"PeerInfo,omitempty"`
	Metadata file.Metadata `json:"Metadata,omitempty"`
}

// ^ Auth Stream Struct ^ //
type authStreamConn struct {
	self       *Node
	readWriter *bufio.ReadWriter
	stream     network.Stream
	callback   Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create a buffer stream for non blocking read and write.
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
		self:       sn,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
		self:       sn,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Write Message on Stream ^ //
func (asc *authStreamConn) Write(authMsg authStreamMessage) error {
	fmt.Println("Auth Msg Struct: ", authMsg)

	// Convert Request to JSON String
	var jsonData []byte
	jsonData, err := json.Marshal(authMsg)
	if err != nil {
		fmt.Println("Error Converting Meta to JSON", err)
		return err
	}
	println("Auth Msg Bytes: ", jsonData)
	fmt.Println("Auth Msg String: ", string(jsonData))

	// Write Message with "Delimiter"=(Seperator for Message Values)
	_, err = asc.readWriter.WriteString(fmt.Sprintf("%s\n", string(jsonData)))
	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}

	// Write buffered data
	err = asc.readWriter.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer")
		return err
	}
	return nil
}

// ^ Read Data from Msgio ^ //
func (asc *authStreamConn) Read() {
	for {
		// ** Read the Buffer **
		str, err := asc.readWriter.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}

		// ** Empty String **
		if str == "" {
			return
		}

		// ** Contains Data **
		// Construct message
		fmt.Println("Received String Message:", str)
		if str != "\n" {
			asm := new(authStreamMessage)
			err := json.Unmarshal([]byte(str), asm)
			if err != nil {
				fmt.Println("Error Unmarshalling Auth Stream Message ", err)
			}

			// Check Message Subject
			switch asm.Subject {
			// @ Request to Invite
			case "Request":
				// Callback the Invitation
				asc.callback.OnInvited(asm.PeerInfo.String(), asm.Metadata.String())

			// @ Response to Invite
			case "Response":
				// Callback to Proxies
				asc.callback.OnResponded(asm.Decision)

				// Handle Decision
				if asm.Decision {
					fmt.Println("Auth Accepted")
				} else {
					// Reset
					fmt.Println("Auth Declined")
				}

			// ! Invalid Subject
			default:
				fmt.Printf("%s.\n", asm.Subject)
			}
		}
	}
}
