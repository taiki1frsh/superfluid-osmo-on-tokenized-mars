# Osmosis Superfluid Staking **on Tokenized Mars**

****This is just POC content****

Make $OSMO that is deposited on Mars stakable via Superfluid Staking by leveraging super flexible design capacity of CosmWasm and Cosmos SDK.
(And beyond, to other lending protocols on Osmosis)

## Concept

The simplest way to describe the goal is to **enable **deposited OSMO **on** Mars to** be** staked.


Looking back at the revolutionary idea of Superfluid Staking as an App Chain, it was quite simple at its core, seeing the potential liquidity that exists locked on-chain but is not being used efficiently as a kind of staking state. In the case of Osmosis, this was $OSMO that was provided as liquidity in constant product AMM pools. As time passed, the Osmosis ecosystem diversified, and so did the ways of providing liquidity. A prime example of this is the Mars Protocol.

In Mars Protocol, there are always more than a certain amount of $OSMO tokens being lent but not used as collateral. This is due to the actions of users who aim to earn interest through lending, as well as the actions of passive borrowers who only borrow a small number of tokens against their collateral.

The nature of these tokens is very similar to that of $OSMO, which is provided as liquidity in constant product AMM pools but is practically unexchangeable due to being in a price range where no trades can occur. Based on this, I believe that it is possible and beneficial to create a staking state for some of the $OSMO collateral in the Mars Protocol through a mechanism similar to Superfluid Staking.

But, beyond that, this attempt is not limited to the Mars Protocol. The goal is to create a mechanism that can be applied to any lending protocol on Osmosis. In other words, it is a mechanism that can be applied to any token that is provided as collateral in a lending protocol on Osmosis.

### User flow

1. Deposit $OSMO to Mars Red-Bank
1. Get mOSMO, wihtout borrowing any tokens
1. Send tx contain `Bond` msg of the `superfluid-msomo` contract to stake OSMO using unutilized mOSMO (OSMO collateral on Red-Bank)
1. The amount of mOSMO which you staked is freezed and you can't use it for borrowing and transfer it
1. Send tx contains `ClaimReward` to claim the reward of staking (this can be changed to the automatic distribution)
1. Send tx contains `Unbond` to exit staking 
1. After the unbonding period, you can get your mOSMO back by sending tx contains `Claim` msg

## Core Designs

Casually describing, I believe `superfluid mars` can be achieved by the following features:

1. Tokenize bonds to know the amount of Osmosis being lent and to improve composability.

- Issuance of mOSMO (CW20)
- (Issuance of debt tokens (CW20))

2. (Low priority) Enable collateral management based on the presence or absence of bonds. For all transactions involving claims, use bond tokens.

- Mint of mToken when depositing
- Locking of claim tokens when borrowing
- Burning of bond tokens when redeeming
- Burning (moving) of bond tokens when liquidating

3. Lock bond tokens in the Superfluid staking Contract, manage the rental usage rate of Osmo, and return the amount to Superfluid staking up to a certain proportion while not in use.

- Calculate the usage rate of Osmo in the Contract
- Calculate the maximum staking amount based on Params
- Hit a msg to update the staking amount
- Enable redemption after 14 days by unlocking

4. A Superfluid staking module for generic use cases of bond tokens.

- Keep the basic functionality the same as the existing LP share token-based Superfluid staking
- Possibly, design the module so that it can support the other bond tokens from Mars

### Bonding

Unlike Superfluid Staking with LP Share Tokens, the value of mOSMO is always 1:1 with OSMO, so the Bonding Amount of mOSMO specified by the user and the amount of OSMO actually bonded through the Superfluid module are the same. The lockup of mOSMO is controlled by the `superfluid-mosmo-staking` contract. While user-specified Validators are not implemented in this case, it is believed that there would be no significant barriers to addressing this through the management of an intermediary account when integrating with the current x/superfluid module.

### Unbonding

Users can essentially exit at any time by executing the Unbond Msg. When the Unbond Msg is executed, the lockup period for unbonding defined by Osmosis is naturally required. Once that period has ended and it can be confirmed, users will be able to move their mOSMO freely through the Claim Msg (tentative).

### Slashing

One of the most complex aspects to consider when staking is slashing. There are several possible ways to execute slashing in Superfluid Staking for collateral tokens in the lending market. However, it is essential to perform the calculation of the amount to be slashed through a common module.

1. Specify the Red-Bank contract address that holds the $OSMO to be burned in advance, and execute the burn for that address during slashing. Additionally, for burning mOSMO, execute a Msg that forcibly burns the necessary amount from the superfluid-mosmo module. In this case, the basic logic is performed entirely within the module.
1. Detect slashing and control both by triggering a slashing Msg through the contract. Perform calculations similar to those in the x/slashing module to determine the slashing amount, and execute slashing by transferring the corresponding total slashing amount of $OSMO from a predetermined slashing insurance pool held by Mars Protocol to the Community Pool. If the amount is insufficient, the remaining slashing will not be performed.

etc...

Inherently, due to the nature of Superfluid, there is no strong necessity to be strict against slashing. Considering this, it is reasonable to incorporate a flexible design.

In this instance, although we will not be implementing the actual system, we will be mindful of peripheral design that envisions the first policy.

### Rebalance

If the maximum staking amount calculated from the Utilization Rate is exceeded, it will be necessary to reduce the staking amount according to the share ratio. This occurs on an epoch-by-epoch basis. Rebalancing is assumed to take place in two stages.

In the first stage, additional staking is no longer allowed. Then, if the allowable Utilization Rate increases by a few more percentage points, the actual staking amount for each account that is already staking is reduced according to their respective shares. This is done by triggering a Rebalance Msg implemented in the x/superfluid-mosmo module by the contract, and the actual processing takes place within that msg.

#### Claim the staking reward

The detailed logic for reward distribution is controlled and executed in the x/superfluid-mosmo module. This is to maintain consistency with the existing x/superfluid module and to integrate and handle a certain type of debt token as a Superfluid bondToken standard in a generic way. Furthermore, since Staking, Slashing, and Bonding are all managed within the module, it is more concise and secure to manage the fine-grained control of reward distribution within the module as well.

- Determine the reward distribution ratio according to the locked mOSMO balance of each Superfluid mOsmo participant.
- This needs to be done by the contract.
- Since the amount of rewards varies for each delegating validator, it is necessary to determine the ratio according to each reward amount.

### Tokenization of the lending functionalities

By tokenizing the lending protocol, external integration with the lending market's state becomes significantly easier, and by utilizing the flexible design of tokens, it becomes possible to provide composability with external markets for collateral, debt, interest, and other financial aspects. As a result, a more efficient lending protocol can be realized.

This time, we particularly want to enable staking of some idle collateral assets through the tokenization of collateral.

## Technical Design

### Tokenized Red-Bank

To enable the staking of unused OSMO collateral through Superfluid, I would like to first propose the tokenization of the lending functionality. This can be easily imagined with the precedents set by Aave, Compound, and others. In simple terms, it involves the issuance of synthetic tokens representing collateral amounts and synthetic tokens representing debt, in order to express interest rates and borrowing conditions through these tokens. In this implementation, we will consider only the implementation of collateral tokens, which are particularly related to Superfluid Staking (though we are aware of the interactions with debt tokens, they will not be implemented in this case).

[Tokenized Red-Bank](https://github.com/taiki1frsh/superfluid-osmo-on-tokenized-mars/blob/main/tokenized-red-bank/contracts/red-bank/README.md)

### mToken

Introduce mToken as a collateral token for Mars Protocol. This corresponds 1:1 with tokens deposited in Red-Bank, and the total issuance of mTokens is equal to the total amount of tokens deposited in Red-Bank.   

Its features include:

- mToken is a CW20
- mToken corresponds 1:1 with tokens deposited in Red-Bank
- mToken holds UnusableAmount data, which restricts free transfers without requiring lockup to a specific contract
- By using UnusableAmount, it is possible to properly distribute interest by realizing the lockup of mOSMO in the Superfluid contract without transferring it to the contract
- UnusableAmount can only be controlled by the contract address that holds the data
- mToken also uses UnusableAmount to handle lockup of collateral due to borrowing
- Collateral redemption can only be performed by the mToken holder at the time of redemption, regardless of the original mToken issuer
- mToken can be basically transferred freely, and can also be controlled by contracts through Allowance as normal CW20

[mToken Contract](https://github.com/taiki1frsh/superfluid-osmo-on-tokenized-mars/blob/main/tokenized-red-bank/contracts/m-token/README.md)

### Superfluid mOSMO Module

We will introduce `x/superfluid-mosmo` as a module in the Osmosis chain. However, the original intention is to integrate it as a module of `x/superfluid`, and instead of implementing specific features for mOSMO, we aim to provide a versatile implementation targeting certain debt tokens like mOSMO. In order to properly achieve this, a deep understanding of the `x/superfluid` module and a thorough comprehension of all the necessary features for staking debt tokens are required. Therefore, at this point, we decided to implement it as a separate module rather than blindly integrating it.

Primarily, this module is responsible for handling the information management and processing execution that the `x/superfluid` module takes care of. This is because the management of states related to the underlying debt tokens is maintained by the contract, making it unsuitable for controlling such information management in the module.

Therefore, the specific functions that this module will handle are as follows:

- Store the staking address, amount, and specified Validator through a message (Msg)
- (It is still undecided whether or not to pass the LockId at this time)
- Associate that information with the Intermediary account created for each Validator address
- Stake the newly issued OSMO through the Intermediary account
- The distribution of rewards should be carried out in each Epoch in a similar manner to `x/superfluid` (ideally, it should be aligned as closely as possible to this method)
- When the Rebalance Msg is called, un-stake and burn the necessary amount of OSMO based on each participant's share.

[x/superfluid-mosmo module](https://github.com/taiki1frsh/superfluid-osmo-on-tokenized-mars/blob/main/superfluid-mosmo-osmosis-custom/x/superfluid-mosmo/README.md)

### Superfluid mOSMO Contract

In this contract, we provide functions that correspond to the control part in the x/superfluid module where Superfluid Staking is controlled and executed by LP share tokens. Specifically, this contract performs the pseudo lockup of the unused portion of mOSMO as collateral using UnusableAmount, and calculates the maximum amount of mOSMO that can be staked in Superfluid Staking based on the usage rate of OSMO in Red-Bank.

It is not possible and should be avoided to introduce all unused mOSMO into Superfluid Staking. This is because if the usage rate of mOSMO becomes 100% due to Superfluid Staking, the usage rate of OSMO in Red-Bank will also become 100%, making debt repayment in Red-Bank impossible.

Therefore, it is desirable to introduce a certain proportion of the unused mOSMO into Superfluid Staking. This proportion is managed and updated as Config data held by this contract. The distribution of staking rewards is maintained according to the locked mOSMO amount associated with each address managed by this contract. Distribution occurs only through data during Epochs, either by releasing it through a Claim Msg or by actually sending tokens.

Unbonding is realized by releasing the UnusableAmount of mOSMO after a period similar to the normal OSMO unbonding period.

[Superfluid-mOsmo-Staking Contract](https://github.com/taiki1frsh/superfluid-osmo-on-tokenized-mars/blob/main/superfluid-mosmo-staking/README.md)

### osmosis-rust

[customized osmosis-rust](https://github.com/taiki1frsh/superfluid-osmo-on-tokenized-mars/blob/main/custom-osmosis-rust/README.md)
