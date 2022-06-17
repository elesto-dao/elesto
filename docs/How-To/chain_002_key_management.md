---
title: Manage Keys
---


This section covers how to manage keys for the Elesto network.


## Basic Key Management

Create, import, export, and delete keys using the Elesto CLI.

### Create a new key

To generate a new key pair, run this command:

```
elestod keys add $KEY_NAME
```

where:

- `$KEY_NAME` is a human readable name for your key

??? Example "Example: generate Alice's key"

    ```shell
    ➜ elestod keys add alice
    ```

    ```shell
    - name: alice
    type: local
    address: elesto1pp7tyzj80hrys3aae043lerkxkd0h3e8mf7khg
    pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AnogArgAOO1CDC87MOGeJ7mcRqIWa0rOiRbDmm9X3ddi"}'
    mnemonic: ""


    **Important** write this mnemonic phrase in a safe place.
    It is the only way to recover your account if you ever forget your password.

    dice argue will silent team drink rate print lift pair copy method rather spy jungle way tribe panther outdoor reject agree employ rain poverty
    ```

The key comes with a "mnemonic phrase", which is serialized into a human-readable 24-word mnemonic. User can recover their associated addresses with the mnemonic phrase.

!!! Danger
    It is important that you keep the mnemonic for address **secure**, as there is **no way** to recover it. You would not be able to recover and access the funds in the wallet if you forget the mnemonic phrase. Do not share your mnemonic key with anyone!!



## List your keys 

Your wallet can host multiple keys. To list the keys available in your wallet, run this command:

```
elestod keys list
```

??? Example "Example: list keys"

    ```shell
    ➜ elestod keys list
    ```

    ```shell
    - name: alice
    type: local
    address: elesto1pp7tyzj80hrys3aae043lerkxkd0h3e8mf7khg
    pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AnogArgAOO1CDC87MOGeJ7mcRqIWa0rOiRbDmm9X3ddi"}'
    mnemonic: ""
    ```



### Delete a key

To remove a key from your wallet use the command:

```
elestod keys delete $KEY_NAME
```

??? Example "Example: delete Alice's keys"
    
    ```shell
    ➜ elestod keys delete alice
    ```

    ```shell
    Key reference will be deleted. Continue? [y/N]: y
    Key deleted forever (uh oh!)
    ```


### Restore existing key by seed phrase

```
elestod keys add $KEY_NAME --recover
```

You can restore an existing key with the mnemonic.

??? Example "Example: restore Alice's keys"

    ```shell
    ➜ elestod keys add alice --recover
    ```

    ```shell
    > Enter your bip39 mnemonic
    ## Enter your 24-word mnemonic here ##

    - name: alice
    type: local
    address: elesto1pp7tyzj80hrys3aae043lerkxkd0h3e8mf7khg
    pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AnogArgAOO1CDC87MOGeJ7mcRqIWa0rOiRbDmm9X3ddi"}'
    mnemonic: ""

    ```

