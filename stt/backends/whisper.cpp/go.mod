module github.com/GRVYDEV/S.A.T.U.R.D.A.Y/stt/backends/whisper.cpp

go 1.20

replace (
	github.com/GRVYDEV/S.A.T.U.R.D.A.Y/log => ../../../log
	github.com/GRVYDEV/S.A.T.U.R.D.A.Y/stt/engine => ../../engine
)

require (
	github.com/GRVYDEV/S.A.T.U.R.D.A.Y/log v0.0.0-20230726165431-f1d355a1a072
	github.com/GRVYDEV/S.A.T.U.R.D.A.Y/stt/engine v0.0.0-20230726165431-f1d355a1a072
	github.com/ggerganov/whisper.cpp/bindings/go v0.0.0-20240103173633-9962371f71b9
)

require golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
