# sandboxes-core

Status: pre-alpha.

Domain model, state machines, operation journal contracts, workspace rules, and
ports implemented by infrastructure modules.

## Boundaries

- no SQLite, containerd, HTTP, or os/exec;
- no dependency on runtime, gateway, daemon, or SDK;
- domain invariants live here and are enforced by the daemon.
