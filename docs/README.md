
# Elesto

[![Go Reference](https://pkg.go.dev/badge/github.com/elesto-dao/elesto.svg)](https://pkg.go.dev/github.com/elesto-dao/elesto)
[![build](https://github.com/elesto-dao/elesto/actions/workflows/quality.yaml/badge.svg?branch=main)](https://github.com/elesto-dao/elesto/actions/workflows/quality.yaml)
[![codecov](https://codecov.io/gh/elesto-dao/elesto/branch/main/graph/badge.svg?token=NLT5ZWM460)](https://codecov.io/gh/elesto-dao/elesto)
[![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/elesto-dao/elesto)](https://libraries.io/go/github.com%2Felesto-dao%2Felesto)
[![DeepSource](https://deepsource.io/gh/elesto-dao/elesto.svg/?label=active+issues&show_trend=true&token=BRR7kVLyskz5-N1etTDRay5J)](https://deepsource.io/gh/elesto-dao/elesto/?ref=repository-badge)



### Summary

Elesto is a protocol designed to build a secure and resilient identity framework based on self-sovereign identity.

Elesto was born from the research project [Cosmos Cash](https://github.com/allinbits/cosmos-cash): a regulatory compliant protocol that offers the same guarantees as traditional banking systems. Features that enable these guarantees are Know Your Customer (KYC), anti-money laundering (AML) tracking,
Financial Action Task Force (FATF) travel rule, and identity management. 

Elesto uses a novel approach to identity management by leveraging W3C specifications for decentralized identifiers and verifiable credentials.

### Research paper

For more information on the research behind the Elesto protocol, see the Cosmos Cash research paper:

[Cosmos Cash: Investigation into EU regulations affecting E-Money tokens](https://drive.google.com/file/d/1zmEyA8kA0uAIRGDKxYElOKvjtz4f_Ep5/view)

### Architecture

The Elesto approach leverages open standards to reach its goals and offers an open model compatible with
third-party projects that use the open standards. In particular, the Elesto project roadmap includes:

- Self-sovereign identity ([SSI](./Reference/GLOSSARY.md#self-sovereign-identity-ssi))
- Decentralized identifier ([DID](./Reference/GLOSSARY.md#decentralized-identifier-did))
- Verifiable credentials ([VC](./Reference/GLOSSARY.md#verifiable-credential-vc))
- Zero-knowledge proofs ([ZKP](./Reference/GLOSSARY.md#zero-knowledge-proof-zkp))

For a detailed architecture description and design choices, visit the [ADR](./Explanation/ADR) section.


--- 

Do you have questions or want to get in touch? Please send us an email at *cosmos-cash@tendermint.com*.
