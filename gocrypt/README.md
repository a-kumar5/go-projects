gocrypt - AES-256-GCM encryption CLI

Features
- AES-256-GCM authenticated encryption
- Key generation (32-byte raw key)
- Passphrase-based key derivation with scrypt (salt stored in a file)
- Stream-friendly stdin/stdout support

Install
- cd gocrypt
- go build ./cmd/crypt
  - produces ./crypt

Usage
  ./crypt gen-key -out key.bin

  echo -n "secret" | ./crypt enc -pass -salt mysalt.bin > out.b64
  cat out.b64 | ./crypt dec -pass -salt mysalt.bin

  echo -n "hello" > msg.txt
  ./crypt enc -in msg.txt -out msg.enc -key key.bin
  ./crypt dec -in msg.enc -out msg.dec -key key.bin

Format
- Encryption outputs base64 of: nonce || ciphertext
- Nonce is 12 bytes, random per message

Security notes
- Keep the key (key.bin) secret with 0600 permissions.
- Use a unique salt per context when deriving from the same passphrase.
- For very large files, consider chunked AEAD with associated data and explicit framing.
