# Testdata

These are test data used in tests 

To set up the test account, launch the chain with the `make start-dev` command then import the test key:

```sh
elestod keys add test -i
```

and when asked use the following mnemonic (and empty password):
```
coil animal waste sound canvas weekend struggle skirt donor boil around bounce grant right silent year subway boost banana unlock powder riot spawn nerve
```

The account address associated with the imported mnemonic will be `elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr` 


Once the account has been imported, transfer some tokens from the `validator` account to the `test` account 

```sh
elestod tx bank send validator elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr 1000000stake --chain-id elesto -y -b block
```

the account details are:

```
Name       test
Address    elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr
PubKeyHex  02c95ff8bb85ce61a646e5e2587c1d2b4603f4e113802c846e0f0f2328bded47f5
PubKeyLen  33
PubKeyType secp256k1
Type       local
Path       
PubKeyB64  Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1

{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aslf+LuFzmGmRuXiWHwdK0YD9OETgCyEbg8PIyi97Uf1"}
```

## Dummy credential

The dummy credential is a credential used for testing: to publish the credential run 

```sh
 elestod tx credential publish-credential-definition \
   dummy \
   Dummy \
   dummy.schema.json \
   dummy.vocab.json \
  --public \
  --from test \
  --gas auto \
  --chain-id elesto -y -b block
```

A credential based on the schema has been already generated:

- `dummy.credential.json`
- `dummy.credential.signed.json`
- `dummy.credential.signed.cosmosadr036.json`

but in case you want to refresh it, run the command:

```sh
elestod query credential prepare-credential dummy --export dummy.credential.json
```

and to sign it

```sh
elestod tx credential issue-public-credential dummy dummy.credential.json --export dummy.credential.signed.json --sign-only --from test --chain-id elesto
```
