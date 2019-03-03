crypt
===

Simple utility to generate and verify bcrypt hashes on the fly. 

# Installation

Go is required

```
go install github.com/iwittkau/crypt
```

# Usage

```
crypt [-h]
```

If no hash is given a new hash will be generated from the password you enter.

If the hash flag is set the hash will compared to the password entered.


Set a hash like this (password = "test"):

```bash
crypt -h='$2a$04$XYTflRUXi7dgWJodhsYpQO.BgPjdd9Hkjgx1GTvoqO9qCNrCH/fi6'
```


# Security

`crypt` uses `golang.org/x/crypto/ssh/terminal`'s `ReadPassword` to read the password. This avoids reading the password from `stdin` as `string`.  
Reading the password as a string can expose passwords because they reside in memory as such.