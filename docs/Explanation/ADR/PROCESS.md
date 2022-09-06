---
title: ADR creation process
---
# ADR creation process

Architectural decision records (ADRs) describe a software design choice that addresses a functional or non-functional requirement that is architecturally significant.
## ADR life cycle

ADR creation is an **iterative** process. Instead of solving all decisions in a single ADR pull request, we MUST initially understand the problem and collect feedback by having conversations in a GitHub Issue.

1. Every ADR proposal SHOULD start with a [new GitHub issue](https://github.com/elesto-dao/elesto/issues/new) or be a result of existing issues. The issue must contain a brief proposal summary.

2. After the motivation is validated, create a new ADR document by copying the [adr-template.md](adr-template.md) file with the following filename pattern: `adr-next_number-title.md`.

   1. Create a draft pull request to propose a new ADR and curate the conversations in the PR.

   1. Make sure the context and solution are clear and well documented.

   1. Add an entry to the proposed section in the [ADR Table of contents](#adr-table-of-contents) list in the [README](README.md) file.

   1. Create a pull request and curate the conversations and review comments.


3. An ADR solution does not have to arrive to the `main` branch with an _accepted_ status in a single PR. If the motivation is clear and the solution is sound, we SHOULD be able to merge PRs iteratively and keep a _proposed_ status. It is preferable to have an iterative approach rather than long, not-merged pull requests.

4. If a _proposed_ ADR is merged, then be sure to document the outstanding issues in the **Further discussion** section of the ADR document notes and in a new GitHub issue.

5. The PR SHOULD always be merged. In the case of a faulty ADR, we still prefer to merge it with a _rejected_ status. The only time the ADR SHOULD NOT be merged is if the author abandons it.

6. Merged ADRs SHOULD NOT be pruned.

### ADR status

Status has two components:

```
{CONSENSUS STATUS} {IMPLEMENTATION STATUS}
```

IMPLEMENTATION STATUS is either `Implemented` or `Not Implemented`.

#### Consensus status

```
DRAFT -> PROPOSED -> LAST CALL yyyy-mm-dd -> ACCEPTED | REJECTED -> SUPERSEEDED by ADR-xxx
                  \        |
                   \       |
                    v      v
                     ABANDONED
```

+ `DRAFT`: [optional] An ADR that is a work in progress and is not ready for a general review. This draft status is to present an early work and get an early feedback in a Draft Pull Request form.

+ `PROPOSED`: An ADR covering a full solution architecture and still in the review - project stakeholders have not reached an agreement yet.

+ `LAST CALL <date for the last call>`: [optional] A clear notify that we are close to accepting updates. Changing status to `LAST CALL` means that social consensus (of Elesto maintainers) has been reached, and we still want to give it time to let the community react or analyze.

+ `ACCEPTED`: ADR that represents a currently implemented or an agreed-to-be-implemented architecture design.

+ `REJECTED`: ADR can go from PROPOSED or ACCEPTED to REJECTED if the consensus among project stakeholders decide to reject.

+ `SUPERSEDED by ADR-xxx`: ADR that has been superseded by a new ADR.

+ `ABANDONED`: The ADR is no longer pursued by the original authors.

## Language used in ADR

+ Write the context and background in the present tense.

+ Avoid using a first, personal form.
