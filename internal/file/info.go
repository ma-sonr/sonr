package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/h2non/filetype"
	md "github.com/sonr-io/core/internal/models"
)

// ^ Struct returned on GetInfo() Generate Preview/Metadata
type FileInfo struct {
	Mime    *md.MIME
	Payload md.Payload
	Name    string
	Path    string
	Size    int32
	IsImage bool
}

// ^ Method Returns File Info at Path ^ //
func GetFileInfo(path string) FileInfo {
	// Initialize
	var mime *md.MIME
	var payload md.Payload
	var subtype string

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}
	defer file.Close()

	// Get Info
	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}

	// @ 2. Create Mime Protobuf
	// Adjust JPEG
	if kind.MIME.Subtype == "jpg" || kind.MIME.Subtype == "jpeg" {
		subtype = "jpg"
	} else {
		subtype = kind.MIME.Subtype
	}

	// Create Data
	mime = &md.MIME{
		Type:    md.MIME_Type(md.MIME_Type_value[kind.MIME.Type]),
		Subtype: subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Find Payload
	if mime.Type == md.MIME_image || mime.Type == md.MIME_video || mime.Type == md.MIME_audio {
		payload = md.Payload_MEDIA
	} else {
		// Get Extension
		ext := filepath.Ext(path)

		// Cross Check Extension
		if ext == "pdf" {
			payload = md.Payload_PDF
		} else if ext == "ppt" || ext == "pptx" {
			payload = md.Payload_PRESENTATION
		} else if ext == "xls" || ext == "xlsm" || ext == "xlsx" || ext == "csv" {
			payload = md.Payload_SPREADSHEET
		} else if ext == "txt" || ext == "doc" || ext == "docx" || ext == "ttf" {
			payload = md.Payload_TEXT
		} else {
			payload = md.Payload_UNDEFINED
		}
	}

	// Return Object
	return FileInfo{
		Mime:    mime,
		Payload: payload,
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: filetype.IsImage(head),
	}
}

// ^ Method Generates new Transfer Card from Contact^ //
func NewCardFromContact(p *md.Profile, c *md.Contact, status md.TransferCard_Status) *md.TransferCard {
	return &md.TransferCard{
		// SQL Properties
		Payload:  md.Payload_CONTACT,
		Received: int32(time.Now().Unix()),
		Preview:  p.Picture,
		Platform: p.Platform,

		// Transfer Properties
		Status: status,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		// Data Properties
		Contact: c,
	}
}

// ^ Method Generates new Transfer Card from URL ^ //
func NewCardFromUrl(p *md.Profile, s string, status md.TransferCard_Status) *md.TransferCard {
	// Return Card
	return &md.TransferCard{
		// SQL Properties
		Payload:  md.Payload_URL,
		Received: int32(time.Now().Unix()),
		Platform: p.Platform,

		// Transfer Properties
		Status: status,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		// Data Properties
		Url: s,
	}
}

// ^ Method Creates AuthInvite from Request ^ //
func NewInviteFromRequest(req *md.InviteRequest, p *md.Peer) md.AuthInvite {
	// Initialize
	var card *md.TransferCard
	var payload md.Payload

	// Determine Payload
	if req.Type == md.InviteRequest_Contact {
		payload = md.Payload_CONTACT
		card = NewCardFromContact(p.Profile, req.Contact, md.TransferCard_DIRECT)
	} else if req.Type == md.InviteRequest_URL {
		payload = md.Payload_URL
		card = NewCardFromUrl(p.Profile, req.Url, md.TransferCard_DIRECT)
	} else {
		payload = md.Payload_UNDEFINED
	}

	// Return Protobuf
	return md.AuthInvite{
		From:     p,
		Payload:  payload,
		Card:     card,
		IsDirect: req.IsDirect,
	}
}
