# contracts-ai

Canonical source for MuxCore **ai.*** domain event type strings, JSON payloads, and shared inference request types.

Import:

```go
import (
  aievents "github.com/Muxcore-Media/contracts-ai/events"
  "github.com/Muxcore-Media/contracts-ai/infer"
)
```

This repository is a **library**, not a sidecar. Feature AI modules (`ai-runtime`, `ai-subtitles`, `ai-recommend`, `ai-tickets`, `ai-filter`, `ai-librarian`) implement behavior; existing non-AI modules must only emit or consume these contracts.

## Event contract rules

1. Never rename an event type string once published.
2. Payload changes are additive — new JSON fields must be optional.
3. Paths, transcripts, and reporter identities are PII. Notification UIs must not print raw `output_path`, `audio_path`, or `reporter` values.

## Modules

| Event prefix | Publisher |
|--------------|-----------|
| `ai.runtime.*` | ai-runtime |
| `ai.subtitle.*` | ai-subtitles |
| `ai.recommend.*` | ai-recommend |
| `ai.ticket.*` | ai-tickets |
| `ai.filter.*` | ai-filter |
| `ai.librarian.*` | ai-librarian |

Cross-module media hooks (empty subtitle search) live in `contracts-media` as `media.subtitles.search.failed`.
