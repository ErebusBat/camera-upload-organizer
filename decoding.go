package photorg

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

type DecoderFunc func(*MoveInfo) error
type Decoder struct {
	Func     DecoderFunc
	Priority int
	Ext      string
	Name     string
	isSystem bool
}

type decoderSlice []Decoder
type decoderMap map[string]decoderSlice

// Len is part of sort.Interface.
func (s decoderSlice) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s decoderSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s decoderSlice) Less(i, j int) bool {
	lhs := s[i]
	rhs := s[j]
	return lhs.Priority < rhs.Priority
}

var decodeMap decoderMap

func init() {
	if decodeMap == nil {
		init_decoding()
	}
}

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

func registerSystemDecoders(handledExts []string, priority int, name string, dFunc DecoderFunc) {
	for _, ext := range handledExts {
		RegisterDecoderInst(&Decoder{
			Priority: priority,
			isSystem: true,
			Ext:      ext,
			Name:     name,
			Func:     dFunc,
		})
	}
}

// Registers a new decoder
func RegisterDecoder(ext string, name string, dFunc DecoderFunc) error {
	ext = normalizeDecoderExtension(ext)
	newDecoder := Decoder{
		isSystem: false,
		Func:     dFunc,
		Priority: math.MaxInt32,
		Ext:      ext,
		// Name:     fmt.Sprintf("Anonymous %s decoder", ext),
		Name: name,
	}
	return RegisterDecoderInst(&newDecoder)
}

func RegisterDecoderInst(decoder *Decoder) error {
	if decodeMap == nil {
		init_decoding()
	}
	ext := decoder.Ext
	var decoders decoderSlice

	if decoder.Name == "" {
		return fmt.Errorf("You must specify a name for your encoder!")
	}
	// log.Printf("Registering Decoder %s for %s\n", decoder.Name, decoder.Ext)

	decoder.Ext = normalizeDecoderExtension(decoder.Ext)

	// Check to make sure it isn't already registerd
	decoders = decodeMap[ext]
	if decoders == nil {
		decoders = make(decoderSlice, 0, 1)
	}
	decoders = append(decoders, *decoder)
	sort.Sort(decoders)
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
		if decoder.isSystem != true {
			log.Printf("[%s]: Invoking decoder: %s\n", moveInfo.fileName, decoder.Name)
		}
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

func GetDecoders(ext string) (decoders decoderSlice, ok bool) {
	ext = normalizeDecoderExtension(ext)
	decoders, ok = decodeMap[ext]
	return
}
