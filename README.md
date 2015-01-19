Rafiki
=========

![rafiki](https://raw.githubusercontent.com/adamar/rafiki/master/doc/rafiki.gif)


Rafiki is a CLI tool for securely storing SSL and RSA files in a local SQLite3 Database. Imported files are first encrypted using GPG and then stored
in the database along with an identifying key (ie. CommonName for CSRs, etc..) 

The database will be created when Rafiki is run for the first time and can be re-located and referenced by Rafiki using the --db flag. 

Note: The term 'key' is used throughout to refer to any/all types of files for simplicity's sake.

Usage
--------------

Import a key
```sh
./Rafiki import --file=/loc/of/file.csr
```

List keys
```sh
./Rafiki list
```

Export a key (using the original filename)
```sh
./Rafiki key export
```


Dependencies
-------------

- SQLite v3+
- Go 1.3+



Development Misc
-------------

Useful Testing Tools

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



ToDo
-----------
- Write more tests
- Add more error checking
- Better text layout
- Print out file details on import & export
- Add (unautheticated) option to profile a key
- Add sub command to "List" option to filter on key type


Useful Resources
------------

http://redkestrel.co.uk/articles/CSR-FAQ/

* [CSR FAQ] - Certificate Signing Request FAQ



[CSR FAQ]:http://redkestrel.co.uk/articles/CSR-FAQ/


