# zcrypt

[![CircleCI](https://circleci.com/gh/IkezoeMakoto/zcrypt.svg?style=svg)](https://circleci.com/gh/IkezoeMakoto/zcrypt)

# What
This is golang CLI tool which can encrypt and decrypt with aes-256-cbc.
You can get the encrypted file by specifying the file you want to encrypt and the encryption key.

# How to get
## Using go get
```
$ go get -u github.com/IkezoeMakoto/zcrypt
```
or
```
wget https://github.com/IkezoeMakoto/zcrypt/releases/download/v1.0.5/zcrypt_darwin_amd64
```

# Usage
```
zcrypt (enc|dec) -in input_path -out output_path -key encrypt_key_path
```
- **(enc|dec)** - you can select encrypt mode or decrypt mode
- **-in input_path** - target input file. For encrypted mode, plain text.
- **-out output_path** - target output file. For encrypted mode, encrypted text.
- **-key encrypt_key_path** - target encrypt key file. **It must be 16 bits or 32 bits**

## Encrypt
```
zcrypt enc -in {secret.file} -out {encrypted.file} -key encrypt_key_path
```

## Decrypt
```
zcrypt dec -in {encrypted.file} -out {decrypted.file} -key encrypt_key_path
```
