# grypto :closed_lock_with_key::robot::man_technologist:

*Understanding cryptographic algorithms by implementing them in Go*

## Introduction :pencil:

**grypto** is a collection of cryptographic algorithms implemented in go.

It also contains a simple CLI that can be used to test and demonstrate the different algorithms implemented in the grypto library.

It was implemented by Tim Ebert ([@timebertt](https://github.com/timebertt)) as a practical exercise to understand the fundamental mathematical concepts behind cryptographic algorithms that were discussed in a lecture on cryptography in his Computer Science Master studies. :mortar_board:

If you want to get a deeper understanding of the basics of cryptography, maybe you will find this collection helpful. Enjoy! :nerd_face::book:

## :warning: WARNING! :warning:

**Please use this only for learning purposes!**

The grypto library and CLI were build only to demonstrate and understand the basics of different cryptographic algorithms. There is no guarantee on correctness, security and quality of the implementation. The implementation might not be compatible with proper implementations of the different algorithms and might be vulnerable to attacks. Please use the respective official implementations of the Go standard library (see https://golang.org/pkg/crypto/) in your Go applications.

## Using the CLI :computer:

You can try out grypto by downloading the CLI:
```
$ go get -u github.com/timebertt/grypto/grypto
```

Encrypt some secret messages:
```
$ grypto caesar encrypt -K 3 -I "Caesar Cipher is old but not very secure"
Fdhvdu Flskhu lv rog exw qrw yhub vhfxuh
```

Have fun! :tada:

## Algorithms implemented :gear:

- [Caesar Cipher](/caesar) (`grypto caesar`)

More to come! :rocket:

## References :books:

Please see the following references further and deeper explanations on the implemented cryptographic algorithms:

- Johannes Buchmann: [Einführung in die Kryptographie](https://doi.org/10.1007/978-3-642-39775-2), [Introduction to Cryptography](https://doi.org/10.1007/978-1-4419-9003-7)
- Christof Paar, Jan Pelzl: [Kryptografie verständlich](https://doi.org/10.1007/978-3-662-49297-0), [Understanding Cryptography](https://doi.org/10.1007/978-3-642-04101-3), http://crypto-textbook.com/
