Rafiki
=========

![rafiki](https://raw.githubusercontent.com/adamar/rafiki/master/doc/rafiki.gif)


Rafiki is a CLI tool for securely storing SSL and RSA files in a local SQLite3 Database. Imported files are first encrypted using GPG and then stored
in the database along with an identifying key (ie. CommonName for CSRs, etc..) 

The database will be created when Rafiki is run for the first time and can be re-located and referenced by Rafiki using the --db flag. 

Note: The term 'key' is used throughout to refer to any/all types of files for simplicity's sake.


Installation
--------------

Ensure that your go bin is setup correctly [GO-BIN]

then run 

```sh
go install github.com/adamar/rafiki
```


Usage
--------------

#### Import a key
```sh
./rafiki import --file=/loc/of/file
```

![rafiki-import](https://raw.githubusercontent.com/adamar/rafiki/master/doc/rafiki-import.gif)

#### List keys
```sh
./rafiki list
```

![rafiki-list](https://raw.githubusercontent.com/adamar/rafiki/master/doc/rafiki-list.gif)

#### Export a key (using the original filename)
```sh
./rafiki export
```

![rafiki-export](https://raw.githubusercontent.com/adamar/rafiki/master/doc/rafiki-export.gif)



Dependencies
-------------

- SQLite v3+
- Go 1.3+



Key Types Supported
-------------

Key Type | Identifier | Supported
-------- | ------ | :---------:
SSL Certificate | Common Name | Yes
SSL Certificate Signing Request | Common Name | Yes
SSL RSA Private Key | MD5 Fingerprint | Yes
SSL ECDSA Private Key | MD5 Fingerprint | Yes
SSH RSA Private Key | - | No
SSH RSA Public Key | MD5 Fingerprint | Yes
SSH DSA Private Key | - | No
SSH DSA Public Key | MD5 Fingerprint | Yes
SSH ECDSA Private Key | - | No
SSH ECDSA Public Key | MD5 Fingerprint | Yes
PGP Private Key | - | No
PGP Public Key | Public Fingerprint | Yes



Development Misc
-------------

#### Useful Testing Tools

Print Public Key Fingerprint
```sh
ssh-keygen -lf /path/to/key.pub
```

Print CSR Info
```sh
openssl req -in domain.com.csr -text -noout
```

Show CSR Public Key
```sh
openssl req -in domain.com.csr -noout -pubkey
```

Show an RSA Key's SHA1 thumbprint
```sh
openssl rsa -noout -modulus -in your-private.key | openssl sha1
```

Show an RSA Key's MD5 thumbprint
```sh
openssl rsa -noout -modulus -in your-private.key | openssl md5
```



To Do
-----------
- ReWrite of internals, the program flow is a mess at the moment
- Write more tests
- Add more error checking
- Better text layout
- Print out file details on import & export
- Add (unautheticated) option to profile a key
- Add sub command to "List" option to filter on key type
- Add API Key file type definition
- Flatten file structure
- Move File checking from import to its own function



Useful Resources
------------

* [CSR FAQ] - Certificate Signing Request FAQ



[CSR FAQ]:http://redkestrel.co.uk/articles/CSR-FAQ/
[GO-BIN]:https://github.com/golang/go/wiki/GOPATH#directory-layout

