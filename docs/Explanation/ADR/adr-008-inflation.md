# ADR 008: Inflation

## Changelog

- 2022-08-20: Initial draft
- 2022-08-30: Draft 2 (inflation rewards allocation, minor grammar edits)
- 2022-09-06: Draft 3 - replace years with epochs

## Status

DRAFT

## Abstract

Inflation is the process by which a currency like the dollar or euro loses value over time, causing the price of goods to rise. Elesto, like other cryptocurrencies, is designed to experience predictable and low rates of inflation. This document details the inflation model for the Elesto chain.


## Context

-

## Decision

Elesto inflation is inspired by Bitcoin inflation model. The Bitcoin supply is limited and known, and the creation of new bitcoin will taper off over time in a predictable way. (There will only ever be 21 million bitcoin, and every four years the amount of bitcoin that is mined is reduced by half.) 

The Elesto initial supply is to be **200,000,000** (two hundred million) tokens and will reach its maximum value at **1,000,000,000** (one billion) in a period of **10 epochs** (roughly 10 years).

The rewards per block are further split into three categories: 10% to the community pool, 10% to the development team, and 80% to the staking rewards.

We introduce the concept of epoch: an epoch is **6,307,200** blocks. The number of blocks of an eopch is calculated as the number of blocks in a calendar year assuming a block rate of 1 block every 5 seconds.

!!! Note
    The actual token configuration uses the `u` (micro) unit to express decimal, therefore the actual value for supply must be multplied for `10^6`


| Epoch n.  [1]  | Current epoch inflation [2] | Epoch inflation supply amount [3]  | Total supply estimate (EOY) [4] | Block inflation amount [5] | **Total supply actual amount (EOY)** [6] |
| ---------- | ---------------- | --------------------- | --------------------------- | --------------- | ------------------------- |
| 1          | 1                | **200,000,000**,000,000   | **400,000,000**,000,000         | **31**,709,792      | **400,000,000**,102,400       |
| 2          | 0.5              | **200,000,000**,000,000   | **600,000,000**,000,000         | **31**,709,792      | **600,000,000**,102,400       |
| 3          | 0.25             | **150,000,000**,000,000   | **750,000,000**,000,000         | **23**,782,344      | **750,000,000**,076,800       |
| 4          | 0.125            | **93,750,000**,000,000    | **843,750,000**,000,000         | **14**,863,965      | **843,750,000**,048,000       |
| 5          | 0.0625           | **52,734,375**,000,000    | **896,484,375**,000,000         | **8**,360,980       | **896,484,373**,056,000       |
| 6          | 0.03125          | **28,015,136**,718,750    | **924,499,511**,718,750         | **4**,441,771       | **924,499,513**,051,200       |
| 7          | 0.0200           | **18,489,990**,234,375    | **942,989,501**,953,125         | **2**,931,569       | **942,989,503**,715,550       |
| 8          | 0.0200           | **18,859,790**,039,063    | **961,849,291**,992,188         | **2**,990,200       | **961,849,291**,393,125       |
| 9          | 0.0200           | **19,236,985**,839,844    | **981,086,277**,832,031         | **3**,050,004       | **981,086,277**,220,988       |
| 10         | 0.019278348      | **18,913,722**,682,071    | **1,000,000,000**,514,100       | **2**,998,751       | **1,000,000,000**,139,230     |

where the columns are:

#### 1. Epoch n.

The epoch from the chain starts, 1 is during the first epoch, 2 the second epoch, and so on. 

#### 2. Current epoch inflation % 

The current epoch inflation is a percentage over the current supply that should be minted for the next epoch. The percentage is between 0-1. The last epoch percentage (0.019278348) is an adjustment over 0.2 to reach the desired value of 1 billion. 

#### 3. Epoch inflation supply amount

The epoch supply inflation is the theoretical amount that will be minted by the end of the current epoch. It is calculated with the formula:

```
Epoch supply inflation amount = FLOOR(Current epoch inflation [2] * Total supply at beginning of the epoch)
```

For the first epoch, the total supply at beginning of the epoch is the initial supply of two hundred million (200,000,000).

#### 4. Total supply estimate (EOY) 

The total supply estimate is the total supply at the end of the current epoch. It is the cumulated sum of each epoch supply.

#### 5. Block inflation amount

The block inflation is the amount to be minted on each block to reach the expected **Epoch inflation supply**. It is calcualted with the formula: 

```
Block inflation amount = ROUND(Epoch inflation supply amount [3] / Blocks per epoch)
```



#### 6. Total supply actual amount (EOY)

Due to rounding errors, the "Total supply estimate (EOY) [4]" is slightly different from the "Total supply actual amount (EOY) [6]." This column is the actual supply that will be obeserved on chain and is calcualted using the formula:

```
Total supply actual amount = Block inflation amount [5] * Blocks per epoch
```



The value for block inflation per epoch is hard coded in the node code in the mint module [ABCI](../../x/mint/abci.go#L21). 




??? Example "Visualizing the inflation curve"

    To visualize the distribution paste the following data to [Vega](https://vega.github.io/editor/#/)

    ```
    {
      "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
      "description": "Supply change over time.",
      "data": {"values": [
          {"epoch": "1", "block": 6307200, "supply": 200000000},
          {"epoch": "2", "block": 12614400, "supply": 600000000},
          {"epoch": "3", "block": 18921600, "supply": 750000000},
          {"epoch": "4", "block": 25228800, "supply": 843750000},
          {"epoch": "5", "block": 31536000, "supply": 896484373},
          {"epoch": "6", "block": 37843200, "supply": 924499513},
          {"epoch": "7", "block": 44150400, "supply": 942989503},
          {"epoch": "8", "block": 50457600, "supply": 961849291},
          {"epoch": "9", "block": 56764800, "supply": 981086277},
          {"epoch": "10", "block": 63072000, "supply": 1000000000},
          {"epoch": "11", "block": 63072000, "supply": 1000000000}

        ]
      },
      "width": 800,
      "height": 600,
      "mark": {
        "type": "line"
      },
      "encoding": {
        "z": {"field": "block", "type": "quantitative"},
        "x": {"field": "epoch", "type": "quantitative"},
        "y": {"field": "supply", "type": "quantitative"}
      }
    }
    ```





## Privacy Considerations

N/A

## Security Considerations

N/A

## Consequences
  
By leveraging the public verifiable credentials, the Elesto node offers native support for revocation lists. Revocation lists are stored as credentials in the node state, within the credential module keeper. 

### Backward Compatibility

N/A

### Positive

N/A

### Negative

N/A

### Neutral

N/A

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section summarizes issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.

Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

## Test Cases [optional]

N/A

## References

<<<<<<< HEAD
[Bitcoin Whitepaper](https://bitcoin.org/bitcoin.pdf)
=======
[Bitcoin Whitepaper](https://bitcoin.org/bitcoin.pdf)
>>>>>>> 66c85d2 (docs: amend ADR #8, introduce epoch concept)
