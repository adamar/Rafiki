Rafiki
=========

Rafiki is a simple SSL cert storage system written in Golang.

![rafiki](https://raw.githubusercontent.com/adamar/rafiki/master/doc/rafiki.gif)


Usage
--------------

Import a CSR file
```sh
./Rafiki csr import --file=/loc/of/file.csr
```

List CSRs
```sh
./Rafiki csr list
```

Export CSR to a file
```sh
./Rafiki csr export --file=/loc/of/newfile.csr
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
openssl req -in domain.com.csr -text -noout

Show CSR Public Key
openssl req -in domain.com.csr -noout -pubkey


ToDo
-----------
- Write Tests
- Add more error checking
- Add key delete function


Useful Resources
------------

http://redkestrel.co.uk/articles/CSR-FAQ/

* [CSR FAQ] - Certificate Signing Request FAQ



[CSR FAQ]:http://redkestrel.co.uk/articles/CSR-FAQ/


