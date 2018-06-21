# Commontypes

I've found I'm always manually creating these types in various projects, so rather than continuing to do them by hand every single time, I'm going to dump them up here and import them.

Have included (Unm|M)arshalText and (Unm|M)arshalJSON for all types for convenient loading from configuration files/json blobs.

## Types

### [network.go](network.go)

Network and Networks - used for creating ip whitelists/blacklists

### [url.go](url.go)

Not as often used as Network/Networks but still useful

### [duration.go](duration.go)

time.Duration wrapper - Everyone has written ths one before

### [keyfile.go](keyfile.go)

loading key files as string or referenced file://, read only - really

## TODO

 * Documentation
 * Write tests
 * Add more

## Copyright and License

Copyright 2016 Shannon Wynter

Licensed under [The MIT License](LICENSE.md)