# Architecture Decision Record

In order to keep track of technical decisions that are made in the project, all major decisions are recorded using an Architecture Decision Record (ADR).

This helps future maintainers of the project understand the reasoning behind technical decisions that were made by providing context about why the decision was made.

The whole document should be one or two pages long. We will write each ADR as if it is a conversation with a future developer. This requires good writing style, with full sentences organised into paragraphs. Bullets are acceptable for visual style and clarity.

## Location and File Naming

- Each design decision should be captured in a file called `/docs/arch/adr-xxx.md` (with `xxx` representing the next number)
- Design decisions that are no longer relevant will be renamed to `/docs/arch/adr-xxx-superceded.md` and updated to reflect why they are no longer applicable
- All ADRs must follow the template below

## Template

The following template should be used when creating a new ADR

```markdown
# ADR XXX: Example ADR Template
<!-- These documents have names that are short noun phrases. For example, "ADR 001: Deployment on Ruby on Rails 3.0.10" or "ADR 009: LDAP for Multitenant Integration" -->

## Context
<!-- This section describes the forces at play, including technological, political, social, and project local. These forces are probably in tension, and should be called out as such. The language in this section is value-neutral. It is simply describing facts. -->

## Decision
<!-- This section describes our response to these forces. It is stated in full sentences, with active voice and listed as bullet points for clarity. For example " - We will use conventional commits for all commit messages" -->

## Status
<!-- A decision may be "PROPOSED" if the project stakeholders haven't agreed with it yet, or "ACCEPTED" once it is agreed. If a later ADR changes or reverses a decision, it may be marked as "DEPRECATED" or "SUPERSEDED" with a reference to its replacement. -->

## Consequences
<!-- This section describes the resulting context, after applying the decision. All consequences should be listed here, not just the "positive" ones. A particular decision may have positive, negative, and neutral consequences, but all of them affect the team and project in the future. -->
```
