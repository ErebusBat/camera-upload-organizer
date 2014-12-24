package photorg

import (
	"fmt"
	"strings"
)

type Decoder func(*MoveInfo) error

var decodeFuncs map[string]Decoder

// Can't use init() because it isn't invoked before RegisterDecoder is called from
// other files in this package
func init_decoding() {
	decodeFuncs = make(map[string]Decoder)
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
func RegisterDecoder(ext string, dFunc Decoder) error {
	if decodeFuncs == nil {
		init_decoding()
	}
	ext = normalizeDecoderExtension(ext)

	// Check to make sure it isn't already registerd
	if _, alreadyReg := decodeFuncs[ext]; alreadyReg {
		return fmt.Errorf("A decoder for >%s< has already been registered!", ext)
	}
	decodeFuncs[ext] = dFunc
	return nil
}

// Returns a Decoder function
func GetDecoder(ext string) (dFunc Decoder, wasFound bool) {
	ext = normalizeDecoderExtension(ext)
	dFunc, wasFound = decodeFuncs[ext]
	return
}
