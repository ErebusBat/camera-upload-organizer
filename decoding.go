package photorg

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type DecoderFunc func(*MoveInfo) error
type Decoder struct {
	Func     DecoderFunc
	Priority int
	Ext      string
	Name     string
}

type decoderMap map[string][]Decoder

var decodeMap decoderMap

// Can't use init() because it isn't invoked before RegisterDecoder is called from
// other files in this package
func init_decoding() {
	// decodeMap = make(map[string]Decoder)
	decodeMap = make(decoderMap)
}
func normalizeDecoderExtension(ext string) string {
	ext = strings.TrimSpace(strings.ToLower(ext))

	// Make sure there is a leading dot, makes using it easier
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return ext
}

// Registers a new decoder
func RegisterDecoder(ext string, dFunc DecoderFunc) error {
	ext = normalizeDecoderExtension(ext)
	newDecoder := Decoder{
		Func:     dFunc,
		Priority: math.MaxInt32,
		Ext:      ext,
		Name:     fmt.Sprintf("Anonymous %s decoder", ext),
	}
	return RegisterDecoderInst(&newDecoder)
}

func RegisterDecoderInst(decoder *Decoder) error {
	if decodeMap == nil {
		init_decoding()
	}
	ext := decoder.Ext
	var decoders []Decoder

	if decoder.Name == "" {
		return fmt.Errorf("You must specify a name for your encoder!")
	}

	// Check to make sure it isn't already registerd
	decoders = decodeMap[ext]
	if decoders != nil {
		// if decoders, alreadyReg := decodeMap[ext]; !alreadyReg {
		// return fmt.Errorf("A decoder for >%s< has already been registered!", ext)
		decoders = make([]Decoder, 0, 1)
	}
	decoders = append(decoders, *decoder)
	decodeMap[ext] = decoders
	return nil
}

// Returns a Decoder function
func RunDecoder(moveInfo *MoveInfo) (wasRan bool) {
	ext := normalizeDecoderExtension(moveInfo.fileExt)
	decoders, wasFound := decodeMap[ext]
	if !wasFound {
		return false
	}

	for _, decoder := range decoders {
		log.Printf("[%s]: Invoking decoder: %s\n", moveInfo.fileName, decoder.Name)
		decoder.Func(moveInfo)

		if moveInfo.DateTaken != nil {
			return true
		}
	}
	log.Println("")

	// If something had ran then it would have been caught as the last
	// criteria of the for loop
	return false
}
