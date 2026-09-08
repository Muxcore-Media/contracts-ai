# Compatibility

| Component | Requirement |
|-----------|-------------|
| Go module path | `github.com/Muxcore-Media/contracts-ai` |
| Import packages | `events`, `infer` |
| Published tag | `v0.1.0` |
| Capability ID (catalog) | `contracts.ai` |
| MuxCore core | ≥ 0.5.8 |

This repository is a **library**, not a sidecar.

## Rules

1. Never rename an event type string once published.
2. Payload changes are additive.
3. Feature AI modules depend on this contract; non-AI modules must not import `infer` to run models.
