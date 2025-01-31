package main

import (
	"errors"
	"flag"
	"math"
	"runtime"
	"strings"
	"time"

	logr "github.com/GRVYDEV/S.A.T.U.R.D.A.Y/log"
	"github.com/GRVYDEV/S.A.T.U.R.D.A.Y/stt/engine"
	whisper "github.com/ggerganov/whisper.cpp/bindings/go"
	"github.com/rs/zerolog"
)

var debug = flag.Bool("debug", false, "print debug logs")

type WhisperModel struct {
	ctx    *whisper.Context
	params whisper.Params
}

func NewWhisperModel() (*WhisperModel, error) {
	flag.Parse()
	if !*debug {
		logr.SetGlobalOptions(logr.GlobalConfig{V: int(zerolog.DebugLevel)})
	}
	ctx := whisper.Whisper_init("./models/ggml-base.bin")
	if ctx == nil {
		return nil, errors.New("failed to initialize whisper")
	}

	params := ctx.Whisper_full_default_params(whisper.SAMPLING_GREEDY)
	params.SetPrintProgress(false)
	params.SetPrintSpecial(false)
	params.SetPrintRealtime(false)
	params.SetPrintTimestamps(false)
	params.SetSingleSegment(false)
	params.SetMaxTokensPerSegment(32)
	params.SetThreads(int(math.Min(float64(4), float64(runtime.NumCPU()))))
	params.SetSpeedup(false)
	params.SetLanguage(ctx.Whisper_lang_id("ko"))

	logger.Infof("Initialized whisper model with params:\n %s", params.String())

	return &WhisperModel{ctx: ctx, params: params}, nil
}

func (w *WhisperModel) Process(samples []float32) (error, engine.Transcription) {
	start := time.Now()
	transcription := engine.Transcription{}
	if err := w.ctx.Whisper_full(w.params, samples, nil, nil); err != nil {
		return err, transcription
	} else {
		segments := w.ctx.Whisper_full_n_segments()
		for i := 0; i < segments; i++ {
			trasncriptionSegment := engine.TranscriptionSegment{}

			trasncriptionSegment.StartTimestamp = uint32(w.ctx.Whisper_full_get_segment_t0(i) * 10)
			trasncriptionSegment.EndTimestamp = uint32(w.ctx.Whisper_full_get_segment_t1(i) * 10)

			trasncriptionSegment.Text = strings.TrimLeft(w.ctx.Whisper_full_get_segment_text(i), " ")

			transcription.Transcriptions = append(transcription.Transcriptions, trasncriptionSegment)
		}
	}
	elapsed := time.Since(start)
	logger.Debugf("Process took %s", elapsed)
	return nil, transcription
}
