package events

import (
	"encoding/json"
	"testing"
)

func TestEventConstants(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"EventRuntimeJobCompleted", EventRuntimeJobCompleted, "ai.runtime.job.completed"},
		{"EventRuntimeJobFailed", EventRuntimeJobFailed, "ai.runtime.job.failed"},
		{"EventSubtitleGenerated", EventSubtitleGenerated, "ai.subtitle.generated"},
		{"EventSubtitleGenerationFailed", EventSubtitleGenerationFailed, "ai.subtitle.generation.failed"},
		{"EventSubtitleSynced", EventSubtitleSynced, "ai.subtitle.synced"},
		{"EventSubtitleSyncFailed", EventSubtitleSyncFailed, "ai.subtitle.sync.failed"},
		{"EventSubtitleTranslated", EventSubtitleTranslated, "ai.subtitle.translated"},
		{"EventRecommendReady", EventRecommendReady, "ai.recommend.ready"},
		{"EventTicketOpened", EventTicketOpened, "ai.ticket.opened"},
		{"EventTicketResolved", EventTicketResolved, "ai.ticket.resolved"},
		{"EventTicketFailed", EventTicketFailed, "ai.ticket.failed"},
		{"EventFilterProfileUpdated", EventFilterProfileUpdated, "ai.filter.profile.updated"},
		{"EventFilterApplied", EventFilterApplied, "ai.filter.applied"},
		{"EventLibrarianFindingsReady", EventLibrarianFindingsReady, "ai.librarian.findings.ready"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("got %q want %q", tc.got, tc.want)
			}
			if !Known(tc.got) {
				t.Fatalf("Known(%q) = false", tc.got)
			}
		})
	}
	if Known("media.movie.added") {
		t.Fatal("media events must not be claimed by contracts-ai")
	}
	if got, want := len(All()), len(tests); got != want {
		t.Fatalf("All() len = %d, want %d", got, want)
	}
}

func TestPayloadJSONRoundTrip(t *testing.T) {
	raw, err := json.Marshal(TicketPayload{TicketID: "t1", Intent: "wrong_audio_language", Language: "en"})
	if err != nil {
		t.Fatal(err)
	}
	var got TicketPayload
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.TicketID != "t1" || got.Intent != "wrong_audio_language" || got.Language != "en" {
		t.Fatalf("round-trip mismatch: %+v", got)
	}
}
