package transmit

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"mime"
	"net/textproto"
	"os"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	st "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
)

var (
	ErrParentDirNotExists = errors.New("FileItem's Parent Directory does not exist")
	ErrEmptyData          = errors.New("Passed Buffer is Empty")
)

// NewFileItem creates a new transfer file item
func NewFileItem(path string, tbuf []byte) (*st.Payload_Item, error) {
	// Extracts File Infrom from path
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Create MIME
	mime, err := NewMime(path)
	if err != nil {
		return nil, err
	}

	// Create Thumbnail on goroutine
	thumbCh := make(chan *st.Thumbnail)
	go NewThumbnail(path, tbuf, mime, thumbCh)

	// Await Thumbnail
	thumb := <-thumbCh

	// Create File Item
	fileItem := &st.FileItem{
		Mime:         mime,
		Path:         path,
		Name:         fi.Name(),
		LastModified: fi.ModTime().Unix(),
		Thumbnail:    thumb,
	}

	// Returns transfer item
	return &st.Payload_Item{
		Mime: mime,
		Data: &st.Payload_Item_File{
			File: fileItem,
		},
	}, nil
}

// Header returns the header of the FileItem
func Header(f *st.FileItem) textproto.MIMEHeader {
	cd := mime.FormatMediaType("item-data", map[string]string{
		"type":         f.GetMime().GetType().String(),
		"filename":     f.GetName(),
		"lastModified": time.Unix(f.GetLastModified(), 0).Format(time.RFC1123),
	})
	return textproto.MIMEHeader{
		"Content-Disposition": {cd},
		"Content-Type":        {f.GetMime().GetValue()},
	}
}

// ToTransferItem Returns Transfer for FileItem
func ToTransferItem(f *st.FileItem) *st.Payload_Item {
	return &st.Payload_Item{
		Thumbnail: f.GetThumbnail(),
		Mime:      f.GetMime(),
		Data: &st.Payload_Item_File{
			File: f,
		},
	}
}

// ** ───────────────────────────────────────────────────────
// ** ─── MIME Management ───────────────────────────────────
// ** ───────────────────────────────────────────────────────
// DefaultUrlMIME is the standard MIME type for URLs
func DefaultUrlMIME() *st.MIME {
	return &st.MIME{
		Type:    st.MIME_TYPE_URL,
		Subtype: ".html",
		Value:   "url/html",
	}
}

// NewMime creates a new MIME type from Path
func NewMime(path string) (*st.MIME, error) {
	// Check if path to file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	// Get MIME Type and Set Proto Enum
	mtype, err := mimetype.DetectFile(path)
	if err != nil {
		return nil, err
	}

	// Return MIME Type
	return &st.MIME{
		Type:    st.MIME_Type(st.MIME_Type_value[strings.ToUpper(mtype.Parent().String())]),
		Value:   mtype.String(),
		Subtype: mtype.Extension(),
	}, nil
}

// Ext Method adjusts extension for JPEG
func Ext(m *st.MIME) string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// IsAudio Checks if Mime is Audio
func IsAudio(m *st.MIME) bool {
	return m.Type == st.MIME_TYPE_AUDIO
}

// IsImage Checks if Mime is Image
func IsImage(m *st.MIME) bool {
	return m.Type == st.MIME_TYPE_IMAGE
}

// IsMedia Checks if Mime is any media
func IsMedia(m *st.MIME) bool {
	return IsAudio(m) || IsImage(m) || IsVideo(m)
}

// IsPDF Checks if Mime is PDF
func IsPDF(m *st.MIME) bool {
	return strings.Contains(m.GetSubtype(), "pdf")
}

// IsVideo Checks if Mime is Video
func IsVideo(m *st.MIME) bool {
	return m.Type == st.MIME_TYPE_VIDEO
}

// IsUrl Checks if Path is a URL
func IsUrl(m *st.MIME) bool {
	return m.Type == st.MIME_TYPE_URL
}

// PermitsThumbnail Checks if Mime Type Allows Thumbnail Generation.
// Image, Video, Audio, and PDF are allowed.
func PermitsThumbnail(m *st.MIME) bool {
	return IsImage(m) || IsVideo(m) || IsAudio(m) || IsPDF(m)
}

// NewThumbnail creates a new thumbnail for the given file
func NewThumbnail(path string, tbuf []byte, mime *st.MIME, ch chan *st.Thumbnail) {
	if IsImage(mime) {

		data, err := genImageThumbnail(path)
		if err == nil {
			ch <- &st.Thumbnail{
				Buffer: data,
				Mime:   mime,
			}
			return
		}
	} else {
		if tbuf != nil {

			ch <- &st.Thumbnail{
				Buffer: tbuf,
				Mime:   mime,
			}
		}
		return
	}

	ch <- nil
}

func genImageThumbnail(path string) ([]byte, error) {
	// Open Image
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	srcImage, _, _ := image.Decode(f)

	// Dimension of new thumbnail 240 X 300
	dstImage := resizeImage(srcImage, 240, 300)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, dstImage, &jpeg.Options{Quality: 75})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func resizeImage(img image.Image, length int, width int) image.Image {
	//truncate pixel size
	minX := img.Bounds().Min.X
	minY := img.Bounds().Min.Y
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y
	for (maxX-minX)%length != 0 {
		maxX--
	}
	for (maxY-minY)%width != 0 {
		maxY--
	}
	scaleX := (maxX - minX) / length
	scaleY := (maxY - minY) / width

	imgRect := image.Rect(0, 0, length, width)
	resImg := image.NewRGBA(imgRect)
	draw.Draw(resImg, resImg.Bounds(), &image.Uniform{C: color.White}, image.ZP, draw.Src)
	for y := 0; y < width; y += 1 {
		for x := 0; x < length; x += 1 {
			averageColor := getImageAverageColor(img, minX+x*scaleX, minX+(x+1)*scaleX, minY+y*scaleY, minY+(y+1)*scaleY)
			resImg.Set(x, y, averageColor)
		}
	}
	return resImg
}

func getImageAverageColor(img image.Image, minX int, maxX int, minY int, maxY int) color.Color {
	var averageRed float64
	var averageGreen float64
	var averageBlue float64
	var averageAlpha float64
	scale := 1.0 / float64((maxX-minX)*(maxY-minY))

	for i := minX; i < maxX; i++ {
		for k := minY; k < maxY; k++ {
			r, g, b, a := img.At(i, k).RGBA()
			averageRed += float64(r) * scale
			averageGreen += float64(g) * scale
			averageBlue += float64(b) * scale
			averageAlpha += float64(a) * scale
		}
	}

	averageRed = math.Sqrt(averageRed)
	averageGreen = math.Sqrt(averageGreen)
	averageBlue = math.Sqrt(averageBlue)
	averageAlpha = math.Sqrt(averageAlpha)

	averageColor := color.RGBA{
		R: uint8(averageRed),
		G: uint8(averageGreen),
		B: uint8(averageBlue),
		A: uint8(averageAlpha)}

	return averageColor
}

func imgToBytes(img image.Image) []byte {
	var opt jpeg.Options
	opt.Quality = 80

	buff := bytes.NewBuffer(nil)
	err := jpeg.Encode(buff, img, &opt)
	if err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}
