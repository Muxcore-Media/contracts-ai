package events

// AI domain event types and JSON payloads.
// Feature modules publish these; existing non-AI modules must not contain model logic.

const (
	EventRuntimeJobCompleted = "ai.runtime.job.completed"
	EventRuntimeJobFailed    = "ai.runtime.job.failed"

	EventSubtitleGenerated        = "ai.subtitle.generated"
	EventSubtitleGenerationFailed = "ai.subtitle.generation.failed"
	EventSubtitleSynced           = "ai.subtitle.synced"
	EventSubtitleSyncFailed       = "ai.subtitle.sync.failed"
	EventSubtitleTranslated       = "ai.subtitle.translated"

	EventRecommendReady = "ai.recommend.ready"

	EventTicketOpened   = "ai.ticket.opened"
	EventTicketResolved = "ai.ticket.resolved"
	EventTicketFailed   = "ai.ticket.failed"

	EventFilterProfileUpdated = "ai.filter.profile.updated"
	EventFilterApplied        = "ai.filter.applied"

	EventLibrarianFindingsReady = "ai.librarian.findings.ready"
)

// RuntimeJobPayload is emitted when an inference job finishes.
type RuntimeJobPayload struct {
	JobID    string `json:"job_id"`
	Kind     string `json:"kind"`
	Provider string `json:"provider"`
	Error    string `json:"error,omitempty"`
}

// SubtitleJobPayload is emitted for generate / sync / translate outcomes.
type SubtitleJobPayload struct {
	JobID       string  `json:"job_id"`
	MediaID     string  `json:"media_id,omitempty"`
	MediaFileID string  `json:"media_file_id,omitempty"`
	Language    string  `json:"language,omitempty"`
	OutputPath  string  `json:"output_path,omitempty"` // PII: filesystem path
	OffsetMs    int64   `json:"offset_ms,omitempty"`
	Scale       float64 `json:"scale,omitempty"`
	Error       string  `json:"error,omitempty"`
}

// RecommendReadyPayload is a completed suggestion batch.
type RecommendReadyPayload struct {
	BatchID    string `json:"batch_id"`
	Count      int    `json:"count"`
	LibraryIDs int    `json:"library_ids"`
}

// TicketPayload is an opened, resolved, or failed household ticket.
type TicketPayload struct {
	TicketID string `json:"ticket_id"`
	Intent   string `json:"intent,omitempty"`
	MediaID  string `json:"media_id,omitempty"`
	Reporter string `json:"reporter,omitempty"` // PII: household identity
	Action   string `json:"action,omitempty"`
	Title    string `json:"title,omitempty"`
	Language string `json:"language,omitempty"`
	Error    string `json:"error,omitempty"`
}

// FilterAppliedPayload is a computed filter plan for playback/transcode peers.
type FilterAppliedPayload struct {
	PlanID     string `json:"plan_id"`
	MediaID    string `json:"media_id"`
	ProfileID  string `json:"profile_id"`
	BleepCount int    `json:"bleep_count"`
	CutCount   int    `json:"cut_count"`
}

// LibrarianFindingsPayload is a completed library QA scan.
type LibrarianFindingsPayload struct {
	ScanID  string `json:"scan_id"`
	Count   int    `json:"count"`
	Summary string `json:"summary,omitempty"`
}

// All returns every AI event type string.
func All() []string {
	return []string{
		EventRuntimeJobCompleted,
		EventRuntimeJobFailed,
		EventSubtitleGenerated,
		EventSubtitleGenerationFailed,
		EventSubtitleSynced,
		EventSubtitleSyncFailed,
		EventSubtitleTranslated,
		EventRecommendReady,
		EventTicketOpened,
		EventTicketResolved,
		EventTicketFailed,
		EventFilterProfileUpdated,
		EventFilterApplied,
		EventLibrarianFindingsReady,
	}
}

// Known reports whether eventType is a defined AI event constant.
func Known(eventType string) bool {
	switch eventType {
	case EventRuntimeJobCompleted,
		EventRuntimeJobFailed,
		EventSubtitleGenerated,
		EventSubtitleGenerationFailed,
		EventSubtitleSynced,
		EventSubtitleSyncFailed,
		EventSubtitleTranslated,
		EventRecommendReady,
		EventTicketOpened,
		EventTicketResolved,
		EventTicketFailed,
		EventFilterProfileUpdated,
		EventFilterApplied,
		EventLibrarianFindingsReady:
		return true
	default:
		return false
	}
}
