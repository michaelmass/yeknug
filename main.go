package main

import (
	"bytes"
	"embed"
	"io"
	"log"

	"github.com/michaelmass/yeknug/pkg/deviceevents"
	"github.com/michaelmass/yeknug/pkg/soundplayer"
	"github.com/pkg/errors"
)

//go:embed assets
var assets embed.FS

func main() {
	nugPlayer := NewPlayer("assets/nug.mp3")
	rechargePlayer := NewPlayer("assets/recharge.mp3")

	for event := range Listener() {
		switch event.Kind {
		case deviceevents.KeyDown:
			nugPlayer.Play()
		case deviceevents.MouseDown:
			rechargePlayer.Play()
		}
	}
}

func Listener() <-chan deviceevents.Event {
	key, err := deviceevents.New()

	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to initialise deviceevents"))
	}

	return key.Events()
}

func NewPlayer(filename string) *soundplayer.SoundPlayer {
	nug, err := assets.ReadFile(filename)

	if err != nil {
		log.Fatal(errors.Wrap(err, "reading mp3 file"))
	}

	player := soundplayer.New()
	err = player.Load(io.NopCloser(bytes.NewReader(nug)))

	if err != nil {
		log.Fatal(errors.Wrap(err, "loading mp3 file"))
	}

	return player
}
