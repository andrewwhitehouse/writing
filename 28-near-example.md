**NEAR Protocol Notes**

[_There are a number of programmable blockchain contenders swirling about currently. The "main" one is Ethereum but the gas fees are ridiculous, it's not clear how the move from a different "proof" mechanism is going to work, and whether or not Ether is treated as a security with all the associated regulation is an open question. There is plenty of unruly behaviour going on. I'm not seeing any particularly useful things being done: swapping tokens, and creating collectible images seem to be the main activities. I'm more interested in how tokens can be used to incentivise behaviour that makes an application valuable; addressing the cold-start problem that affects networks when they're starting out._

_So I'm exploring the different blockchains, and this may lead to a thread through the book. The NEAR procotol uses Rust, like Solana, but has a different approach to scaling, which may be significant if and when adoption grows._ 

_These are written in note form. Not quite book-ready yet, but I'm still making sense of it myself._]

Building a sample appliation on NEAR

There are a few blockchain that I think have potential.

For Ethereum, the gas fees are currently too high. So you need to use a level 2 blockchain like Arbitrum.

There is also Polygon, which has its own MATIC token. I haven't looked at that in detail yet.

I'm not ready to use that yet.

So I'm going to look at NEAR and Solana instead.

This is me following the process to get an example application working with NEAR. (I did setup Solana in the past and it was fiddly.)

[Basic setup skeleton](https://www.near-sdk.io/zero-to-hero/basics/set-up-skeleton)

```
$ npm install -g near-cli
npm WARN config global `--global`, `--local` are deprecated. Use `--location=global` instead.
npm WARN deprecated @ledgerhq/hw-transport-u2f@5.36.0-deprecated: @ledgerhq/hw-transport-u2f is deprecated. Please use @ledgerhq/hw-transport-webusb or @ledgerhq/hw-transport-webhid. https://github.com/LedgerHQ/ledgerjs/blob/master/docs/migrate_webusb.md

added 304 packages, and audited 305 packages in 44s

25 packages are looking for funding
  run `npm fund` for details

4 moderate severity vulnerabilities

To address all issues (including breaking changes), run:
  npm audit fix --force

Run `npm audit` for details.
$ 
```

It seemed to work.

```
$ near
Usage: near <command> [options]

Commands:
  near create-account <accountId>                                                                        create a new developer account (subaccount of the masterAccount, ex: app.alice.test)
  near state <accountId>                                                                                 view account state

. . .
```

I already have awhitehouse.testnet set up.

https://examples.near.org/rust-counter

panic profit flower mobile thing use example soft sing pizza clump purity

awhitehouse.testnet

Now I'm logged in


```
$ near login

Please authorize NEAR CLI on at least one of your accounts.

If your browser doesn't automatically open, please visit this URL
https://wallet.testnet.near.org/login/?referrer=NEAR+CLI&public_key=ed25519%3A76wh6QA8ZCLYLY9Y2hkWEVvHdEpLKaGXfUxEAkJPw4op&success_url=http%3A%2F%2F127.0.0.1%3A5000
Please authorize at least one account at the URL above.

Which account did you authorize for use with NEAR CLI?
Enter it here (if not redirected automatically):
Logged in as [ awhitehouse.testnet ] with public key [ ed25519:76wh6Q... ] successfully
$
```

```
$ near keys awhitehouse.testnet
Keys for account awhitehouse.testnet
[
  {
    access_key: { nonce: 86198558000004, permission: 'FullAccess' },
    public_key: 'ed25519:3UsvjSDjYnQ361mUGHY8m84PdjtQjA7ShN9bju5cjkBy'
  },
  {
    access_key: { nonce: 86158837000000, permission: 'FullAccess' },
    public_key: 'ed25519:3hdhYJA29cHHBkYih51J1x2rLW1cxRD2h1u61Dc7VUy7'
  },
  {
    access_key: { nonce: 86158744000004, permission: 'FullAccess' },
    public_key: 'ed25519:4datLDvthG1e2ePQQVruTUxTmwJQLFGyRwHyhGtuvCnd'
  },
  {
    access_key: { nonce: 93802546000000, permission: 'FullAccess' },
    public_key: 'ed25519:76wh6QA8ZCLYLY9Y2hkWEVvHdEpLKaGXfUxEAkJPw4op'
  },
  {
    access_key: { nonce: 86158738000002, permission: 'FullAccess' },
    public_key: 'ed25519:9u1QXyrPRFhBcGbVfM46jgqvVGq1KL6kbQZs9u9HBURo'
  },
  {
    access_key: { nonce: 86198444000001, permission: 'FullAccess' },
    public_key: 'ed25519:A1dyYPPp9duzvHLyogvxWumG5AbNKQSvD6rxBxGYcHaD'
  },
  {
    access_key: { nonce: 93802465000001, permission: 'FullAccess' },
    public_key: 'ed25519:AdDrUvREPy2Bp4L23XLtv5msDMgNX327GE61dSAmPmYk'
  },
  {
    access_key: { nonce: 86158785000000, permission: 'FullAccess' },
    public_key: 'ed25519:C3EPPhDRNLNWZbAEA18ZcKndBhaq8mt4AFHcZGpGMKon'
  }
]
Andrews-MBP:reflections andrewwhitehouse$
```

I already have Rust setup so am skipping the next section.

```
$ rustup target add wasm32-unknown-unknown
info: component 'rust-std' for target 'wasm32-unknown-unknown' is up to date
$ 
```

Clone the template

```
$ git clone https://github.com/near-examples/rust-template
Cloning into 'rust-template'...
remote: Enumerating objects: 55, done.
remote: Counting objects: 100% (2/2), done.
remote: Total 55 (delta 1), reused 1 (delta 1), pack-reused 53
Unpacking objects: 100% (55/55), 13.47 KiB | 229.00 KiB/s, done.
$
```

```
$ cargo test
  Downloaded libc v0.2.99
  Downloaded serde_json v1.0.66
  Downloaded proc-macro2 v1.0.28
  Downloaded typenum v1.13.0
  Downloaded near-sys v0.1.0
  Downloaded near-sdk v4.0.0-pre.4
  Downloaded near-sdk-macros v4.0.0-pre.4
  Downloaded indexmap v1.7.0
  Downloaded cpufeatures v0.1.5
  Downloaded sha2 v0.9.5
  Downloaded derive_more v0.99.16
  Downloaded 11 crates (961.1 KB) in 2.50s
   Compiling proc-macro2 v1.0.28
   Compiling autocfg v1.0.1
   Compiling unicode-xid v0.2.2
   Compiling syn v1.0.57
   Compiling serde_derive v1.0.118
   Compiling serde v1.0.118
   Compiling typenum v1.13.0
   Compiling version_check v0.9.3
   Compiling ryu v1.0.5
   Compiling serde_json v1.0.66
   Compiling hashbrown v0.11.2
   Compiling memchr v2.4.1
   Compiling itoa v0.4.7
   Compiling block-padding v0.2.1
   Compiling ahash v0.4.7
   Compiling opaque-debug v0.3.0
   Compiling hex v0.4.3
   Compiling regex-syntax v0.6.25
   Compiling convert_case v0.4.0
   Compiling cfg-if v1.0.0
   Compiling cpufeatures v0.1.5
   Compiling libc v0.2.99
   Compiling lazy_static v1.4.0
   Compiling keccak v0.1.0
   Compiling wee_alloc v0.4.5
   Compiling bs58 v0.4.0
   Compiling base64 v0.13.0
   Compiling memory_units v0.4.0
   Compiling cfg-if v0.1.10
   Compiling byteorder v1.4.3
   Compiling Inflector v0.11.4
   Compiling near-sys v0.1.0
   Compiling generic-array v0.14.4
   Compiling indexmap v1.7.0
   Compiling num-traits v0.2.14
   Compiling num-integer v0.1.44
   Compiling num-bigint v0.3.2
   Compiling num-rational v0.3.2
   Compiling hashbrown v0.9.1
   Compiling aho-corasick v0.7.18
   Compiling quote v1.0.9
   Compiling regex v1.5.4
   Compiling digest v0.9.0
   Compiling block-buffer v0.9.0
   Compiling sha2 v0.9.5
   Compiling sha3 v0.9.1
   Compiling near-runtime-utils v4.0.0-pre.1
   Compiling borsh-derive-internal v0.8.2
   Compiling borsh-schema-derive-internal v0.8.2
   Compiling derive_more v0.99.16
   Compiling near-sdk-macros v4.0.0-pre.4
   Compiling toml v0.5.8
   Compiling proc-macro-crate v0.1.5
   Compiling near-rpc-error-core v0.1.0
   Compiling near-rpc-error-macro v0.1.0
   Compiling borsh-derive v0.8.2
   Compiling borsh v0.8.2
   Compiling near-primitives-core v0.4.0
   Compiling near-vm-errors v4.0.0-pre.1
   Compiling near-vm-logic v4.0.0-pre.1
   Compiling near-sdk v4.0.0-pre.4
   Compiling my-crossword v0.1.0 (/Users/andrewwhitehouse/code/active-projects/rust-template)
warning: unused import: `super::*`
  --> src/lib.rs:25:9
   |
25 |     use super::*;
   |         ^^^^^^^^
   |
   = note: `#[warn(unused_imports)]` on by default
help: consider adding a `#[cfg(test)]` to the containing module
  --> src/lib.rs:24:1
   |
24 | mod tests {
   | ^^^^^^^^^

warning: unused import: `get_logs`
  --> src/lib.rs:26:32
   |
26 |     use near_sdk::test_utils::{get_logs, VMContextBuilder};
   |                                ^^^^^^^^
   |
help: consider adding a `#[cfg(test)]` to the containing module
  --> src/lib.rs:24:1
   |
24 | mod tests {
   | ^^^^^^^^^

warning: unused import: `testing_env`
  --> src/lib.rs:27:20
   |
27 |     use near_sdk::{testing_env, AccountId};
   |                    ^^^^^^^^^^^
   |
help: consider adding a `#[cfg(test)]` to the containing module
  --> src/lib.rs:24:1
   |
24 | mod tests {
   | ^^^^^^^^^

warning: function is never used: `get_context`
  --> src/lib.rs:31:8
   |
31 |     fn get_context(predecessor: AccountId) -> VMContextBuilder {
   |        ^^^^^^^^^^^
   |
   = note: `#[warn(dead_code)]` on by default

warning: `my-crossword` (lib test) generated 4 warnings
    Finished test [unoptimized + debuginfo] target(s) in 3m 12s
     Running unittests (target/debug/deps/my_crossword-77d530044811f886)

running 0 tests

test result: ok. 0 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s

You have mail in /var/mail/andrewwhitehouse
Andrews-MBP:rust-template andrewwhitehouse$ 
```

No tests yet.

[Altering the smart contract](https://www.near-sdk.io/zero-to-hero/basics/add-functions-call)

> This section will modify the smart contract skeleton from the previous section. This tutorial will start by writing a contract in a somewhat useless way in order to learn the basics. Once we've got a solid understanding, we'll iterate until we have a crossword puzzle.

I like the approach.

Update the puzzle

```
const PUZZLE_NUMBER: u8 = 1;

#[near_bindgen]
#[derive(Default, BorshDeserialize, BorshSerialize)]
pub struct Contract {
    crossword_solution: String
}

#[near_bindgen]
impl Contract {
    pub fn get_puzzle_number(&self) -> u8 {
        PUZZLE_NUMBER
    }

    pub fn set_solution(&mut self, solution: String) {
        self.crossword_solution = solution;
    }

    pub fn guess_solution(&mut self, solution: String) {
        if solution == self.crossword_solution {
            env::log_str("You guessed right!")
        } else {
            env::log_str("Try again.")
        }
    }
}
```

```
$ ./build.sh
```

Create a subaccount

```
$ near create-account crossword.awhitehouse.testnet --masterAccount awhitehouse.testnet --initialBalance 1
Saving key to '/Users/andrewwhitehouse/.near-credentials/testnet/crossword.awhitehouse.testnet.json'
Account crossword.awhitehouse.testnet for network "testnet" was created.
$
```

Now that we have a key pair for our subaccount, we can deploy the contract to testnet and interact with it!

```
$ near state crossword.awhitehouse.testnet
Account crossword.awhitehouse.testnet
{
  amount: '1000000000000000000000000',
  block_hash: 'DgT2AYqtVeqnfC2pmPg9VefpBkAfnYtGvihX8j32LKXq',
  block_height: 93803952,
  code_hash: '11111111111111111111111111111111',
  locked: '0',
  storage_paid_at: 0,
  storage_usage: 182,
  formattedAmount: '1'
}
$ 
```

Note the code_hash here is all ones. This indicates that there is no contract deployed to this account.

```
$ near deploy crossword.awhitehouse.testnet --wasmFile res/my_crossword.wasm
Starting deployment. Account id: crossword.awhitehouse.testnet, node: https://rpc.testnet.near.org, helper: https://helper.testnet.near.org, file: res/my_crossword.wasm
Transaction Id 2TJehwkVuRnLEiy1xpwi1NMwKsgRQHVH3SfFK3ikmchW
To see the transaction in the transaction explorer, please open this url in your browser
https://explorer.testnet.near.org/transactions/2TJehwkVuRnLEiy1xpwi1NMwKsgRQHVH3SfFK3ikmchW
Done deploying to crossword.awhitehouse.testnet
$ 
```

```
near state crossword.friend.testnet
```

After deployment the code hash has changed:

```
$ near state crossword.awhitehouse.testnet
Account crossword.awhitehouse.testnet
{
  amount: '999246714091318300000000',
  block_hash: 'EvS9xfJG9NoppUfC3vfyd5JVQwkKbZV1zjJtVdSKRTCr',
  block_height: 93804095,
  code_hash: '8nnhpFmmCN25ft5iNPVwtk5igZd2vbtBNdbfxKD2xGBV',
  locked: '0',
  storage_paid_at: 0,
  storage_usage: 97501,
  formattedAmount: '0.9992467140913183'
}
Andrews-MBP:rust-template andrewwhitehouse$
```

```
$ near view crossword.awhitehouse.testnet get_puzzle_number
View call: crossword.awhitehouse.testnet.get_puzzle_number()
1
Andrews-MBP:rust-template andrewwhitehouse$
```

Set a solution:

```
$ near call crossword.awhitehouse.testnet set_solution '{"solution": "near nomicon ref finance"}' --accountId awhitehouse.testnet
Scheduling a call: crossword.awhitehouse.testnet.set_solution({"solution": "near nomicon ref finance"})
Doing account.functionCall()
Transaction Id D241jMtB17w5SmRcxRajTPyCjXNRURBGVx4PQZ3KJjLF
To see the transaction in the transaction explorer, please open this url in your browser
https://explorer.testnet.near.org/transactions/D241jMtB17w5SmRcxRajTPyCjXNRURBGVx4PQZ3KJjLF
''
Andrews-MBP:rust-template andrewwhitehouse$ 
```

```
$ near call crossword.awhitehouse.testnet guess_solution '{"solution": "near nomicon ref finance"}' --accountId awhitehouse.testnet
Scheduling a call: crossword.awhitehouse.testnet.guess_solution({"solution": "near nomicon ref finance"})
Doing account.functionCall()
Receipt: 8PyBbPLzpEPXbKM8VW11M7djDy4U5mW92pSbVA5Y8CW2
	Log [crossword.awhitehouse.testnet]: You guessed right!
Transaction Id 61315nTU9iWmVbCRVa2S7xUTq5yLZA3DM4j1jVL2KPHv
To see the transaction in the transaction explorer, please open this url in your browser
https://explorer.testnet.near.org/transactions/61315nTU9iWmVbCRVa2S7xUTq5yLZA3DM4j1jVL2KPHv
''
Andrews-MBP:rust-template andrewwhitehouse$ 
```

Now we clear down the subaccount.

```
near delete crossword.awhitehouse.testnet awhitehouse.testnet
near create-account crossword.awhitehouse.testnet --masterAccount awhitehouse.testnet
```

Next steps: <https://www.near-sdk.io/zero-to-hero/basics/hashing-and-unit-tests>

**Reflection**: This tools actually did what the tutorial said they would. That is quite encouraging. Generally the NEAR team seem to put a good amount of effort into making decent docs. Solana seems to be further ahead in adoption in some ways, but platform builders need to make their tools friendly to the people who are going to build on the platform.

@Beaver 


