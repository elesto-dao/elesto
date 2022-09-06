---
title: About Architecture Decision Records (ADRs)
---


# Architecture Decision Records (ADRs)

This section includes all high-level architecture decisions for the Elesto protocol.

### Definitions

Within the context of an ADR, we define the following:

- Architectural decision records (ADRs) are documents that describe a software design choice that addresses a functional or non-functional requirement that is architecturally significant.

- An architectural decision record (ADR) document describes a software design choice that addresses a functional or non-functional requirement that is architecturally significant. The collection of ADRs created and maintained in this project constitutes its decision log. 

- An architecturally significant requirement (ASR) is a requirement that has a measurable effect on a software system’s architecture and quality.

All these records are within the topic of Architectural Knowledge Management (AKM).

You can read more about the ADR concept in the [Documenting architecture decisions, the Reverb way](https://product.reverb.com/documenting-architecture-decisions-the-reverb-way-a3563bb24bd0#.78xhdix6t) blog post.

## Rationale

ADRs are intended to be the primary mechanism for proposing new feature designs and processes, collecting community input on an issue, and documenting the design decisions.

An ADR provides:

- Context on the relevant goals and the current state
- Proposed changes to achieve the goals
- Summary of pros and cons
- References
- Changelog

Note the distinction between an ADR and a specification. The ADR provides the context, intuition, reasoning, and justification for a change in architecture or the architecture of something new. The specification is a summary of everything as it stands today.

If recorded decisions turned out to be lacking the required substance, the process is to convene a discussion, record the new decisions here, and then modify the code to match.

## Creating new ADRs

See [ADR Creation Process](PROCESS.md).

#### Use RFC 2119 keywords

When writing ADRs, follow the best practices that apply to writing [RFCs](https://www.ietf.org/standards/rfcs/).

Keywords are used to signify the requirements in the specification and are often capitalized: 

- "MUST"
- "MUST NOT"
- "REQUIRED"
- "SHALL"
- "SHALL NOT"
- "SHOULD"
- "SHOULD NOT"
- "RECOMMENDED"
- "MAY"
- "OPTIONAL"

Keywords are to be interpreted as described in [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119).

## ADR Table of Contents

### Accepted

- [ADR 001](adr-001-docs-structure.md)
### Proposed

