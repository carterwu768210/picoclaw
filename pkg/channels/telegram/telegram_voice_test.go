package telegram

import "testing"

func TestInferTelegramMediaFilename(t *testing.T) {
	tests := []struct {
		name        string
		defaultBase string
		original    string
		mimeType    string
		fallbackExt string
		want        string
	}{
		{
			name:        "uses original filename when present",
			defaultBase: "audio",
			original:    "meeting.m4a",
			mimeType:    "audio/mpeg",
			fallbackExt: ".mp3",
			want:        "meeting.m4a",
		},
		{
			name:        "derives extension from mime type",
			defaultBase: "voice",
			original:    "",
			mimeType:    "audio/ogg",
			fallbackExt: ".mp3",
			want:        "voice.oga",
		},
		{
			name:        "falls back to configured extension",
			defaultBase: "audio",
			original:    "clip",
			mimeType:    "",
			fallbackExt: ".mp3",
			want:        "clip.mp3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferTelegramMediaFilename(tt.defaultBase, tt.original, tt.mimeType, tt.fallbackExt)
			if got != tt.want {
				t.Fatalf("inferTelegramMediaFilename() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestContentTagForDocument(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		mimeType string
		want     string
	}{
		{name: "audio mime tagged as audio", filename: "file.bin", mimeType: "audio/ogg", want: "[audio]"},
		{name: "audio extension tagged as audio", filename: "voice.ogg", mimeType: "application/octet-stream", want: "[audio]"},
		{name: "non-audio document tagged as file", filename: "report.pdf", mimeType: "application/pdf", want: "[file]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contentTagForDocument(tt.filename, tt.mimeType)
			if got != tt.want {
				t.Fatalf("contentTagForDocument() = %q, want %q", got, tt.want)
			}
		})
	}
}
