# Design System — Hello World Acceptance

> Source of truth: approved `index.html`.
> Every value below is extracted from it. Changing a value here without
> changing approved design is a defect.

Last updated: 2026-09-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background, input background, button text |
| `--color-surface` | `#FFFFFF` | Form field surface |
| `--color-surface-raised` | `#FFFFFF` | Not separately drawn; same as page surface |
| `--color-border` | `#000000` | Input border |
| `--color-text` | `#000000` | Heading, input text, status text |
| `--color-text-muted` | `#000000` | Not separately drawn; status text uses default text color |
| `--color-primary` | `#2563EB` | Save button background and border, focus ring |
| `--color-primary-text` | `#FFFFFF` | Save button text |
| `--color-success` | `#000000` | Success status message text |
| `--color-warning` | `#000000` | Validation status message text |
| `--color-danger` | `#000000` | Not drawn; use default text if needed before design changes |
| `--color-focus` | `#2563EB` | Input and button focus outline |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA |
| `--color-border` | `--color-bg` | `21:1` | UI border |
| `--color-primary-text` | `--color-primary` | `5.17:1` | AA |
| `--color-focus` | `--color-bg` | `5.17:1` | UI focus |

### 1.2 Spacing

Base unit: `6px`. Every margin, padding, and gap in product uses one of these, except known deviations listed in section 4.

| Token | Value |
|---|---|
| `--space-0` | `0` |
| `--space-2` | `12px` |
| `--space-3` | `18px` |
| `--space-4` | `24px` |

Actual spacing uses:

| Value | Used for |
|---|---|
| `0` | Reset margins, input/button vertical padding |
| `12px` | Form gap, input horizontal padding, status negative top margin |
| `18px` | Button horizontal padding |
| `24px` | Body padding, section gap |

### 1.3 Typography

Font families (include fallback stack and how font is loaded):

- Body: `Arial, Helvetica, sans-serif`; system/local font stack, no external font loading.
- Headings: `Arial, Helvetica, sans-serif`; same inherited family.
- Mono: none used.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-sm` | `14px` | normal/browser default | inherited `400` | Status message |
| `--text-base` | browser default `16px` | normal/browser default | `400` | Input text |
| `--text-base-strong` | browser default `16px` | normal/browser default | `700` | Save button |
| `--text-display` | `clamp(44px, 9vw, 80px)` | `1` | `700` | h1 greeting |

Heading levels are used in order. Only `h1` appears.

Weight and letter-spacing tokens:

| Token | Value | Used for |
|---|---|---|
| `--font-weight-body` | `400` | Running text, input, status |
| `--font-weight-medium` | `700` | Button label |
| `--font-weight-heading` | `700` | h1 greeting |
| `--tracking-tight` | `-0.04em` | h1 greeting |
| `--tracking-normal` | `normal` | Input, button, status |

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-sm` | `6px` | Input and button |
| `--radius-md` | `6px` | Same control radius; no larger radius drawn |
| `--radius-lg` | `6px` | No modal/card drawn; keep same until design changes |
| `--radius-full` | not used | No pill/avatar drawn |
| `--border-width` | `1px` | Input and button border |
| `--outline-width` | `3px` | Focus-visible outline |
| `--outline-offset` | `3px` | Focus-visible outline offset |
| `--shadow-sm` | none | No shadows drawn |
| `--shadow-md` | none | No dropdown/popover drawn |
| `--shadow-lg` | none | No modal drawn |
| `--duration-fast` | `0ms` | No hover/focus transitions drawn |
| `--duration-base` | `0ms` | No panel motion drawn |
| `--easing` | none | No transitions drawn |

Motion: design has no animation or transition. `prefers-reduced-motion` needs no override because no movement exists.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `base` | `0` | `main: min(100%, 560px)`; `form: min(100%, 420px)` | 1 | `24px` body padding |
| `compact-form` | max width `520px` | same containers | Form fields stack vertically | `12px` form gap |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |
| Sticky header | not used |
| Dropdown | not used |
| Modal backdrop | not used |
| Modal | not used |
| Toast | not used |

## 2. Components

### 2.1 GreetingHeading

**Purpose** — Display current persisted greeting as primary page content. Do not use for secondary labels or navigation.

**Anatomy** — `[greeting text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--text-display`, `--font-weight-heading`, `--tracking-tight`, `--color-text` | Single greeting display |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Responsive display | Content height from `line-height: 1` | none | `--text-display` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Black centered greeting text wraps anywhere | `--color-text`, `--text-display` |

**Accessibility** — Render as `h1` with stable ID referenced by section `aria-labelledby`. Preserve text content as current greeting. Heading order starts at `h1`.

### 2.2 GreetingInput

**Purpose** — Let visitor enter replacement greeting. Use only for single-line greeting text.

**Anatomy** — `[visually hidden label] [text input]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-surface`, `--color-text`, `--color-border`, `--radius-sm`, `--border-width` | Editable greeting text |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `44px` | `0 12px` | `--text-base` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White field, black text, black 1px border | `--color-surface`, `--color-text`, `--color-border` |
| Focus visible | Blue 3px outline offset by 3px | `--color-focus`, `--outline-width`, `--outline-offset` |

**Accessibility** — Native text input with real label. Label may be visually hidden but must remain associated by `for`/`id`. Minimum hit target height is `44px`. On widths ≤ 520px, input width is `100%` and `flex: none` keeps it `44px` tall.

### 2.3 SaveButton

**Purpose** — Submit updated greeting. Use for primary save action only.

**Anatomy** — `[label]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Primary | `--color-primary`, `--color-primary-text`, `--radius-sm`, `--border-width`, `--font-weight-medium` | Save greeting |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `44px` | `0 18px` | `--text-base-strong` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Blue fill and border, white bold label, pointer cursor | `--color-primary`, `--color-primary-text` |
| Focus visible | Blue 3px outline offset by 3px | `--color-focus`, `--outline-width`, `--outline-offset` |

**Accessibility** — Native `button type="submit"`. Minimum hit target height is `44px`. Keyboard activation uses browser defaults: Enter/Space. Label uses sentence/title case exactly as design: `Save`.

### 2.4 GreetingForm

**Purpose** — Group greeting input and Save button into one submission unit.

**Anatomy** — `[GreetingInput] [SaveButton]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Inline | `--space-2` gap, max width `420px` | Widths above `520px` |
| Stacked | `--space-2` gap, full-width controls | Widths `520px` and below |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | Content height | none | inherited |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Inline row with `12px` gap | `--space-2` |
| Compact | Column layout with `12px` gap; input and button both `44px` high and `100%` wide | `--space-2` |

**Accessibility** — Use native `form` submit behavior. `novalidate` lets scripted validation message control wording.

### 2.5 FormMessage

**Purpose** — Report validation or save result without moving focus unless validation fails.

**Anatomy** — `[message text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Status | `--text-sm`, `--color-text` | Validation and saved messages |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `min-height: 20px` | none | `--text-sm` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Empty | Blank line reserves `20px` height | `--text-sm` |
| Validation | Text reads `Enter a greeting before saving.` | `--text-sm`, `--color-text` |
| Saved | Text reads `Saved.` | `--text-sm`, `--color-text` |

**Accessibility** — Use `role="status"` and `aria-live="polite"`. On validation failure, move focus to input.

### 2.6 GreetingPageLayout

**Purpose** — Center one greeting section on blank page. Do not add navigation or extra regions.

**Anatomy** — `[main] [section] [GreetingHeading] [GreetingForm] [FormMessage]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Centered single-section | `--color-bg`, `--space-4`, max width `560px` | Whole app page |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Viewport centered | `min-height: 100vh` | `24px` body padding | inherited |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White background, centered single column, `24px` section gap | `--color-bg`, `--space-4` |

**Accessibility** — Main content uses `<main>`. Section labels itself with greeting heading. No navigation landmarks because design has none.

## 3. Content and formatting

- Voice and tone: plain, direct, minimal.
- Date, time, number, and currency formats: none used.
- Capitalization rule: heading uses greeting text as stored; button label uses title case as shown (`Save`); validation/success messages use sentence case.
- Empty-state and error-message wording pattern: empty status reserves space with no text; validation message tells user exact correction: `Enter a greeting before saving.`; success confirmation is short: `Saved.`.

## 4. Known deviations

Places where approved design does not follow its own rules or anti-patterns in `references/ai-defaults.md`. Record, do not silently fix.

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Spacing scale | `1px`, `3px`, `6px`, `14px`, `20px`, `44px`, `420px`, `520px`, `560px` are layout/control/type values outside 6px spacing scale | Values are border, outline, radius, type, hit target, and container/breakpoint dimensions from approved design, not spacing tokens | Keep as non-spacing tokens; do not reuse as arbitrary spacing |
| Component states | No hover, active, disabled, loading, or API error visuals are drawn | Approved mockup shows default/focus and scripted status messages only | Add states only after approved design change |
| Radius scale | Same `6px` radius used for all drawn controls; no 3-step radius scale | Minimal one-page app draws only input and button | Extend radius scale when new component types are designed |
| Empty status | Empty `role="status"` area renders blank but reserves height | Approved mockup intentionally starts with no message | Keep blank initial status unless content design changes |

AI-default checks avoided by approved design:

- No purple/indigo default palette; accent is `#2563EB` from stakeholder brief.
- No gradients.
- No maximum rounding; controls use `6px`, no pills.
- No heavy shadows.
- Layout matches one real task: read and update greeting.
- No emoji iconography.
- No filler copy; real strings are `Hello, World!`, `Save`, `Enter a greeting before saving.`, and `Saved.`.
- Focus states are visible for input and button.
- No text over images.
- No hover-only affordances.

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2026-09-14 | Initial design system extracted from approved `index.html` | This PR |
