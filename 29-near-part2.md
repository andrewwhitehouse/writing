NEAR continued

[link](https://www.near-sdk.io/zero-to-hero/basics/add-functions-call)

Remember that blockchain is an open ledger, meaning everyone can see the state of smart contracts and transactions taking place.

There is an endpoint that allows you to query the state of the blockchain:

```
$ curl -d '{"jsonrpc": "2.0", "method": "query", "id": "see-state", "params": {"request_type": "view_state", "finality": "final", "account_id": "crossword.awhitehouse.testnet", "prefix_base64": ""}}' -H 'Content-Type: application/json' https://rpc.testnet.near.org
{"jsonrpc":"2.0","result":{"block_hash":"WQ6uiPB1JqBSvPZS1w1QySrSWurhnxNCMrT5nrYkPPq","block_height":93884705,"proof":[],"values":[]},"id":"see-state"}
$
```

There's nothing in there currently because I cleared it down at the end of the last session.

Redeploy:

```
$ near deploy crossword.awhitehouse.testnet --wasmFile res/my_crossword.wasm
Starting deployment. Account id: crossword.awhitehouse.testnet, node: https://rpc.testnet.near.org, helper: https://helper.testnet.near.org, file: res/my_crossword.wasm
Transaction Id 4WwaJqMDDTyyRxMyQ8SXi9ybGZKZtxCswQeSaotspXX9
To see the transaction in the transaction explorer, please open this url in your browser
https://explorer.testnet.near.org/transactions/4WwaJqMDDTyyRxMyQ8SXi9ybGZKZtxCswQeSaotspXX9
Done deploying to crossword.awhitehouse.testnet
$
```

```
$ curl -d '{"jsonrpc": "2.0", "method": "query", "id": "see-state", "params": {"request_type": "view_state", "finality": "final", "account_id": "crossword.awhitehouse.testnet", "prefix_base64": ""}}' -H 'Content-Type: application/json' https://rpc.testnet.near.org
{"jsonrpc":"2.0","result":{"block_hash":"6nN7c67VSPW4D6sEk4djvpnKgnKePZ6ALyzRNQkjUTES","block_height":93885392,"proof":[],"values":[]},"id":"see-state"}Andrews-MBP:rust-template andrewwhitehouse$
```

```
$ near view crossword.awhitehouse.testnet get_puzzle_number
View call: crossword.awhitehouse.testnet.get_puzzle_number()
1
$
```

Set the solution:

```
$ near call crossword.awhitehouse.testnet set_solution "{\"solution\": \"near nomicon ref finance\"}" --accountId awhitehouse.testnet
Scheduling a call: crossword.awhitehouse.testnet.set_solution({"solution": "near nomicon ref finance"})
Doing account.functionCall()
Transaction Id 8sZAarSqNvzXyJTm3QTfjMHSqkytFkzvmsAzay78T3o7
To see the transaction in the transaction explorer, please open this url in your browser
https://explorer.testnet.near.org/transactions/8sZAarSqNvzXyJTm3QTfjMHSqkytFkzvmsAzay78T3o7
''
$
```

Now we see some values ...

```
$ curl -d '{"jsonrpc": "2.0", "method": "query", "id": "see-state", "params": {"request_type": "view_state", "finality": "final", "account_id": "crossword.awhitehouse.testnet", "prefix_base64": ""}}' -H 'Content-Type: application/json' https://rpc.testnet.near.org
{"jsonrpc":"2.0","result":{"block_hash":"6ZKNbW79yWst7DKc5y3yHykq4eSLuvUpbBTH9z5Z24by","block_height":93885558,"proof":[],"values":[{"key":"U1RBVEU=","proof":[],"value":"GAAAAG5lYXIgbm9taWNvbiByZWYgZmluYW5jZQ=="}]},"id":"see-state"}
$
```

```
$ base64 --decode <<< GAAAAG5lYXIgbm9taWNvbiByZWYgZmluYW5jZQ==
near nomicon ref finance
$
```

## Hash the solution, add basic unit tests

[link](https://www.near-sdk.io/zero-to-hero/basics/hashing-and-unit-tests)

The solution that we stored will be public to anyone who looks at the state of the contract. We can hide this by using a one-way function called a hash; we store the hash and then perform the same function on any solution attempt we receive. If the solution matches the guess is correct (it is extremely unlikely that another solution attempt will create the same hash, given the number of possible values, which is 2^256).

It isn't possible to reverse-engineer the hash input (i.e. the solution) from the stored hash, without trying every possible input, which would take an infeasibly long time. If you could check 1 million hashes every second, it would take you around 10^71 seconds. There are approximately 3 x 10^9 seconds in 100 years.

This assumes that the hash function doesn't have any flaws, which is why you should use tried and tested cryptographic libraries for your operations.

The NEAR tutorial links to this [video](https://www.youtube.com/watch?v=PfabikgnD08&feature=youtu.be) which explains more about hashing functions. And here is [MIT Open Courseware](https://www.youtube.com/watch?v=KqqOXndnvic).

## Helper unit test during rapid iteration

We don't have a "main" function in our smart contract to interact with from the command line (through "cargo run") but we do have tests.

Let's add a test.

Update Cargo.toml for dependencies:

```
near-sdk = "4.0.0"
hex = "0.4.3"
```

`src/lib.rs` add this at the end of the file

```
#[cfg(test)]
mod tests {
    use super::*;
    use near_sdk::test_utils::VMContextBuilder;
    use near_sdk::testing_env;

    #[test]
    fn debug_get_hash() {
        // Basic set up for a unit test
        testing_env!(VMContextBuilder::new().build());

        // Using a unit test to rapidly debug and iterate
        let debug_solution = "near nomicon ref finance";
        let debug_hash_bytes = env::sha256(debug_solution.as_bytes());
        let debug_hash_string = hex::encode(debug_hash_bytes);
        println!("Let's debug: {:?}", debug_hash_string);
    }
}
```

```
$ cargo test
. . .
running 1 test
test tests::debug_get_hash ... ok

test result: ok. 1 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s

$
```

Let's add another unit test.

```
// part of writing unit tests is setting up a mock context
// provide a `predecessor` here, it'll modify the default context
fn get_context(predecessor: AccountId) -> VMContextBuilder {
    let mut builder = VMContextBuilder::new();
    builder.predecessor_account_id(predecessor);
    builder
}

#[test]
fn check_guess_solution() {
    // Get Alice as an account ID
    let alice = AccountId::new_unchecked("alice.testnet".to_string());
    // Set up the testing context and unit test environment
    let context = get_context(alice);
    testing_env!(context.build());

    // Set up contract object and call the new method
    let mut contract = Contract::new(
        "69c2feb084439956193f4c21936025f14a5a5a78979d67ae34762e18a7206a0f".to_string(),
    );
    let mut guess_result = contract.guess_solution("wrong answer here".to_string());
    assert!(!guess_result, "Expected a failure from the wrong guess");
    assert_eq!(get_logs(), ["Try again."], "Expected a failure log.");
    guess_result = contract.guess_solution("near nomicon ref finance".to_string());
    assert!(guess_result, "Expected the correct answer to return true.");
    assert_eq!(
        get_logs(),
        ["Try again.", "You guessed right!"],
        "Expected a successful log after the previous failed log."
    );
}
```

Actually to get this working, I had to modify it, compared to the [online docs](https://www.near-sdk.io/zero-to-hero/basics/hashing-and-unit-tests):

```
#[near_bindgen]
impl Contract {
    pub fn get_puzzle_number(&self) -> u8 {
        PUZZLE_NUMBER
    }

    pub fn set_solution(&mut self, solution: String) {
        self.crossword_solution = solution;
    }

    pub fn guess_solution(&mut self, solution: String) -> bool {
        let hashed_attempt = hex::encode(env::sha256(solution.as_bytes()));
        if hashed_attempt == self.crossword_solution {
            env::log_str("You guessed right!")
        } else {
            env::log_str("Try again.")
        }
        hashed_attempt == self.crossword_solution
    }
}
```

```
#[test]
    fn check_guess_solution() {
        // Get Alice as an account ID
        let alice = AccountId::new_unchecked("alice.testnet".to_string());
        // Set up the testing context and unit test environment
        let context = get_context(alice);
        testing_env!(context.build());

        // Set up contract object and call the new method
        let mut contract = Contract {
            crossword_solution: "69c2feb084439956193f4c21936025f14a5a5a78979d67ae34762e18a7206a0f".to_string()
        };
        let mut guess_result = contract.guess_solution("wrong answer here".to_string());
        assert!(!guess_result, "Expected a failure from the wrong guess");
        assert_eq!(get_logs(), ["Try again."], "Expected a failure log.");
        guess_result = contract.guess_solution("near nomicon ref finance".to_string());
        assert!(guess_result, "Expected the correct answer to return true.");
        assert_eq!(
            get_logs(),
            ["Try again.", "You guessed right!"],
            "Expected a successful log after the previous failed log."
        );
    }
```

which was rather disappointing.

