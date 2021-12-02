package soundplayer

import (
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/pkg/errors"
)

type SoundPlayer struct {
	buffer *beep.Buffer
}

func New() *SoundPlayer {
	return &SoundPlayer{}
}

func (soundPlayer *SoundPlayer) Load(rc io.ReadCloser) error {
	streamer, format, err := mp3.Decode(rc)

	if err != nil {
		return errors.Wrap(err, "unable to decode mp3 file")
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
	soundPlayer.buffer = buffer

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	return nil
}

func (soundPlayer *SoundPlayer) Play() {
	streamer := soundPlayer.buffer.Streamer(0, soundPlayer.buffer.Len())
	speaker.Play(streamer)
}
