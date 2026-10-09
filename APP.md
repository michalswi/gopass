# App prompt: gopass

Describe **gopass** as a Go command-line utility for encrypting and decrypting files with a 32-byte secret key. Focus on the behavior implemented in this repository, not a password manager, key vault, or graphical app.

## Purpose and features

- Generate a usable secret with `-g` or `--generate`. The program obtains 32 cryptographically random bytes, base64-encodes the first 24 bytes, and prints the resulting 32-character string after its banner. The printed key has 192 bits of randomness despite its 32-byte length; use that string as the 32-byte ASCII key for encryption and decryption.
- Encrypt a file with `-e` or `--encrypt`, followed by the key and input path. Write the result to `encrypted.dat` by default, or use `-ef <output-path>` to choose another destination.
- Decrypt a file with `-d` or `--decrypt`, followed by the same key and an encrypted input path. Write the result to `decrypted.dat` by default, or use `-df <output-path>` to choose another destination.
- Show help with `-h` or `--help`, and print the version with `-v` or `--version`. Display a colored banner and status messages for generation, encryption, and decryption.

## Cryptography and file format

Use AES-256-GCM authenticated encryption with an exactly 32-byte key. Generate a fresh random 12-byte nonce for each encryption. Store binary output as `nonce || ciphertext || authentication tag`; decryption checks the tag and fails for a wrong key or altered/truncated input. Read input files as bytes and write output files with mode `0644`. Report read, write, and cryptographic errors rather than claiming success. File contents are written to disk, **not** to standard output.

## Example CLI workflow

```sh
gopass -g
key='mZLbpOORzppLSOirl6KBnqUGQ56jJI1f'
gopass -e "$key" data.txt
gopass -d "$key" encrypted.dat
gopass -e "$key" data.txt -ef /tmp/encrypted.dat
gopass -d "$key" /tmp/encrypted.dat -df /tmp/decrypted.dat
```

The example key above is public; use your own generated key for real data. The current default decrypt path reads `encrypted.dat` even if another input path is supplied without `-df`; use the explicit `-df` form when decrypting from a different location. Treat the key as sensitive: passing it as a command-line argument can expose it in shell history or process listings. Generated secrets are printed, not saved to a key file.

## Project context

This is a single Go executable (`github.com/michalswi/gopass`) targeting Go 1.25.5. It uses the Go standard library for cryptography and `github.com/michalswi/color` for terminal colors. The repository documents installation via `go install github.com/michalswi/gopass@latest`; its Makefile has macOS ARM64, Linux AMD64, and Linux ARM64 build targets.
