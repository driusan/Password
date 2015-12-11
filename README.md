# Go Password Package

This is a Go implementation of the PHP 5.5+ simplified password hashing API
described [here](http://php.net/manual/en/book.password.php).

The goal is to implement a simple API for strong password hashing with secure
defaults which can be upgraded as stronger algorithms become available without
modifying the code using the library.

Currently, it's a small wrapper around `x/crypto/bcrypt`. This implementation
is backwords compatible with PHP's implementation (at least Verify and Hash 
are, NeedsRehash might disagree about the security of a hash since PHP 
uses a cost of 10 by default while `x/crypto/bcrypt` uses a default of 12) and
can be used for porting existing PHP applications to Go.

A simple program using this library to verify a login would look like this:

```go
var usersupplied Password.PlainTextPassword; // Get this from the user input
var hash         Password.PasswordHash;      // Get this from your database

if verified, _ := Password.Verify(usersupplied, hash); verified == true {
    // Check if the hash needs upgrading. This must be done after verifying
    // that the password matches the old hash.
    if Password.NeedsRehash(hash) {
        newhash, _ := Password.Hash(usersupplied);
        // save newhash to the database.
    }
    // The password is validated, do stuff.
} else {
    // The password supplied does not match the hash, do something else
}
```

The following things are different from the PHP specification:

1. PlainTextPassword and PasswordHash are typed strings instead of arbitrary strings.
2. The `password_` prefix is removed since Go requires the `Password.` package
   name to reference the functions in this library.
3. Function names are changed from underscore separated to MixedCaps convention
   to be more idiomatic for Go .
