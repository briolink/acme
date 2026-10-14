# acme

Briolink's P0 acceptance project: desk bookings, kept small. P0 exits when
the [runner](https://github.com/briolink/briolink/tree/main/runner) completes
every task here unattended ([roadmap](https://github.com/briolink/briolink/blob/main/docs/5-roadmap.md#stages)).
Its run records also feed the tier spike
([briolink#16](https://github.com/briolink/briolink/issues/16)). The rules
are examples for measuring the pipeline, not a client's business rules.

It is laid out as any project repo: `spec/` holds the model (one `Booking`
entity with a flat statechart, its guards, actions and invariants, four pure
functions, seven routes), `tasks/` the task contracts, one issue each. The
coder writes `domain/acme/`; the generators write `gen/`.

`internal/speccheck` checks the set with briolink's format v0 decoders
(pinned in `go.mod`): every file decodes and validates, routes resolve
against the model, every `spec_ref` resolves, and every function has one
test task and one logic task that depends on it.

## Expected outcomes

Predictions, written before any run, for the spike to compare against.

| Tasks | Kind | Executor | Expected |
| --- | --- | --- | --- |
| 0001–0004 | test | reasoner (`routine`) | Pass; boundary tests from the rules |
| 0006–0010, 0012, 0015–0018 | logic | coder | First or second attempt: one comparison or one field set |
| 0005 `deposit_covers` | logic | coder | Possible escalation: rounding up a percentage |
| 0011 `quote` | logic | coder | Likely escalation: started blocks from a duration |
| 0013 `refund_due` | logic, medium | coder | Likely escalation: tiered time thresholds, inclusive bounds, rounding down |
| 0014 `overstay_charge` | logic, medium | coder | Likely escalation: a grace inside the first block, the rate shared with `quote` |

Tasks behind a route (0011, 0012, 0015–0018) also run the generated handler
tests; tasks behind a transition also run its generated transition tests.

## Open

- Storage: the handlers need load and save; the `sqlite` plugin and its
  queries join once the generator spike
  ([briolink#13](https://github.com/briolink/briolink/issues/13)) settles them.
- The plugin choices are not yet resolved against manifests: plugins ship
  no `plugin.json` yet.
- The `mutate` argv follows the format fixture; `quality/mutate` fixes its
  real flags when it is built.
