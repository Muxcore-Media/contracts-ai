package infer

// Shared inference request/response types for ai-runtime and feature modules.
// Feature modules call ai-runtime; they do not embed model clients themselves.

const (
	KindComplete      = "complete"
	KindEmbed         = "embed"
	KindTranscribe    = "transcribe"
	KindClassify      = "classify"
	KindTranslate     = "translate"
	ProviderHeuristic = "heuristic"
	ProviderOpenAI    = "openai"
	ProviderWhisper   = "whisper-cli"
)

// CompleteRequest is a local or OpenAI-compatible chat completion.
type CompleteRequest struct {
	JobID     string `json:"job_id,omitempty"`
	Provider  string `json:"provider,omitempty"`
	System    string `json:"system,omitempty"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens,omitempty"`
}

// CompleteResponse is model text plus the provider that produced it.
type CompleteResponse struct {
	JobID    string `json:"job_id"`
	Provider string `json:"provider"`
	Text     string `json:"text"`
}

// EmbedRequest turns text into a dense vector for similarity.
type EmbedRequest struct {
	JobID    string `json:"job_id,omitempty"`
	Provider string `json:"provider,omitempty"`
	Text     string `json:"text"`
}

// EmbedResponse is a unit-length (or heuristic) embedding.
type EmbedResponse struct {
	JobID      string    `json:"job_id"`
	Provider   string    `json:"provider"`
	Embedding  []float64 `json:"embedding"`
	Dimensions int       `json:"dimensions"`
}

// TranscribeRequest is ASR input. TranscriptHint lets tests and offline
// households supply known text without a cloud model.
type TranscribeRequest struct {
	JobID          string  `json:"job_id,omitempty"`
	Provider       string  `json:"provider,omitempty"`
	MediaID        string  `json:"media_id,omitempty"`
	AudioPath      string  `json:"audio_path,omitempty"` // PII: filesystem path
	Language       string  `json:"language,omitempty"`
	DurationSec    float64 `json:"duration_sec,omitempty"`
	TranscriptHint string  `json:"transcript_hint,omitempty"`
}

// TranscriptCue is one timed line.
type TranscriptCue struct {
	StartMs int64  `json:"start_ms"`
	EndMs   int64  `json:"end_ms"`
	Text    string `json:"text"`
}

// TranscribeResponse is timed cues plus detected language.
type TranscribeResponse struct {
	JobID    string          `json:"job_id"`
	Provider string          `json:"provider"`
	Language string          `json:"language"`
	Cues     []TranscriptCue `json:"cues"`
}

// ClassifyRequest labels text or a scene description.
type ClassifyRequest struct {
	JobID    string   `json:"job_id,omitempty"`
	Provider string   `json:"provider,omitempty"`
	Text     string   `json:"text"`
	Labels   []string `json:"labels,omitempty"`
	MediaID  string   `json:"media_id,omitempty"`
}

// ClassifyHit is one label with confidence in [0,1].
type ClassifyHit struct {
	Label      string  `json:"label"`
	Confidence float64 `json:"confidence"`
}

// ClassifyResponse is ranked labels.
type ClassifyResponse struct {
	JobID    string        `json:"job_id"`
	Provider string        `json:"provider"`
	Hits     []ClassifyHit `json:"hits"`
}

// TranslateRequest converts subtitle or ticket text between languages.
type TranslateRequest struct {
	JobID    string `json:"job_id,omitempty"`
	Provider string `json:"provider,omitempty"`
	Text     string `json:"text"`
	Source   string `json:"source,omitempty"`
	Target   string `json:"target"`
}

// TranslateResponse is the translated text.
type TranslateResponse struct {
	JobID    string `json:"job_id"`
	Provider string `json:"provider"`
	Text     string `json:"text"`
	Source   string `json:"source"`
	Target   string `json:"target"`
}
