# SRS — Greeting

Module: `greeting`
Design: [View the approved design](http://localhost:8080/design/e70e7c39-7d4c-4cce-8d6f-8e3b47b2c550)
Design system: `design/design-system.md`

> One file per module, at `docs/greeting/SRS.md`. It covers only the functions that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

Greeting module lets any visitor view one stored greeting and replace it. Without it, "Hello World Acceptance" cannot prove database persistence, Go API read/write behaviour, or Next.js rendering through team pipeline.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Anyone opening the public page; no sign-in exists | View current greeting, enter replacement greeting, save replacement greeting |

## 3. Scope

**In scope** — functions specified below, by plan title:

- Persistent greeting editor

**Out of scope** — reasonable adjacent work not included here:

- Sign-in and permissions — deliberately not built; stakeholder requested no sign-in.
- Navigation and additional sections — deliberately not built; approved design has one centered section only.
- External services — deliberately not built; stakeholder requested no external services.
- Multiple greetings or history — deliberately not built; product scope contains one current greeting only.

## 4. Functional requirements

### 4.1 Persistent greeting editor

**Requirement GREETING-001 — View persisted greeting**

*As a* Visitor, *I want to* see current stored greeting as page heading, *so that* page reflects persisted product state.

Behaviour:

1. Visitor opens greeting page.
2. Page reads current greeting through backend service.
3. Page shows current greeting as only `h1` heading in centered single-section layout.
4. When no greeting has been changed yet, displayed and editable value is `Hello, World!`.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/greeting/test-cases/persistent-greeting-editor.md`.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Stored greeting has never been changed | Visitor opens page | Heading text is `Hello, World!` |
| AC-2 | Stored greeting is `Pipeline accepted` | Visitor opens page | Heading text is `Pipeline accepted` |
| AC-3 | Stored greeting is `Hello, World!` | Visitor opens page | Text field value is `Hello, World!` |
| AC-4 | Page is open | Visitor inspects page structure | Page contains one `main` region, one centered section, one `h1`, one text field, one `Save` button, and one status message region |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Missing stored greeting | No saved greeting exists during first app use | Visitor sees `Hello, World!` as current greeting |
| Upstream failure | Current greeting cannot be loaded | No error screen is part of the approved design; API error envelope belongs in service contract |
| Permission | Visitor is not signed in | Not applicable: this module has no sign-in and all visitors may view greeting |

**Data touched**

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Stores current greeting displayed in heading and input; default product value is `Hello, World!` |

**Requirement GREETING-002 — Save replacement greeting**

*As a* Visitor, *I want to* replace greeting through text field and Save button, *so that* new greeting becomes current persisted value.

Behaviour:

1. Visitor edits greeting text field.
2. Visitor submits form by activating `Save`.
3. If trimmed greeting is not empty, page saves replacement through backend service.
4. After successful save, heading updates to saved trimmed text without navigation.
5. Text field value updates to saved trimmed text.
6. Status message reads `Saved.`.
7. After browser reload, saved greeting remains visible as heading and text field value.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/greeting/test-cases/persistent-greeting-editor.md`.

| # | Given | When | Then |
|---|---|---|---|
| AC-5 | Page shows `Hello, World!` | Visitor enters `Pipeline accepted` and selects `Save` | Heading changes to `Pipeline accepted` without navigation |
| AC-6 | Visitor saved `Pipeline accepted` successfully | Browser reloads page | Heading text is `Pipeline accepted` |
| AC-7 | Visitor saved `Pipeline accepted` successfully | Browser reloads page | Text field value is `Pipeline accepted` |
| AC-8 | Visitor enters `  Trim me  ` | Visitor selects `Save` | Saved heading text is `Trim me` |
| AC-9 | Visitor saves non-empty greeting successfully | Save completes | Status message text is `Saved.` |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Text field is empty or whitespace only when submitted | Nothing is saved; status message reads `Enter a greeting before saving.`; focus returns to input |
| Boundary | Greeting contains long unbroken text | Heading wraps anywhere and page has no horizontal scroll at supported widths |
| Conflict | Two visitors save different greetings | Later successful save becomes current greeting; prior visitor sees latest saved value after reload |
| Upstream failure | Replacement greeting cannot be saved | No error screen is part of the approved design; API error envelope belongs in service contract |
| Permission | Visitor is not signed in | Not applicable: this module has no sign-in and all visitors may save greeting |

**Data touched**

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Input is trimmed before save; empty or whitespace-only values are rejected; latest successful save is current value |

## 5. Screens

The approved design shows one page state. It contains: white page background; centered `main`; one centered section labelled by greeting heading; large black greeting heading; visually hidden label text `Greeting`; text input; blue `Save` button; reserved polite status message area; responsive stacked form at widths `520px` and below. No navigation, loading state, empty state, error screen, hover state, active state, disabled state, or animation is shown.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting editor | Entire approved one-page design (`main > section`) | GREETING-001, GREETING-002 | default |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Accessibility | Text input has programmatic label `Greeting`; section is labelled by `h1`; status message uses polite live region; input and button have visible focus outlines; text/background contrast is at least 4.5:1; button text/background contrast is at least 4.5:1 |
| Responsive | Page works from `320px` viewport width and up with no horizontal page scroll; at widths `520px` and below input and button stack, both remain `44px` high and `100%` wide |
| Persistence | A greeting saved by Visitor remains current after browser reload and service restart, provided database data remains intact |
| Localisation | Product copy is English; exact interface strings are `Hello, World!`, `Greeting`, `Save`, `Enter a greeting before saving.`, and `Saved.` |
| Privacy | No personal data is required or stored by this module; greeting text is public page content |

## 7. Dependencies and assumptions

- **Depends on:** approved design, for exact screen elements, states, layout, colors, spacing, responsive behaviour, and copy.
- **Depends on:** `design/design-system.md`, for reusable visual tokens and component rules.
- **Depends on:** backend service, for reading and saving current greeting.
- **Depends on:** PostgreSQL persistence, for retaining latest saved greeting across reloads and restarts.
- **Assumption:** Greeting text is public content because there is no sign-in and page is public. If false, auth scope must be added before build.

| Open question | Proposed default | Who decides |
|---|---|---|
| — | None. Current scope is decided by stakeholder brief and approved design. | — |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Persistent greeting editor | GREETING-001, GREETING-002 | `test-cases/persistent-greeting-editor.md` |
