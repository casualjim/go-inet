## RELASE NOTES

The API is now almost stable.

# go-inet
[![Build Status](https://img.shields.io/travis/gaissmai/go-inet.svg)](https://travis-ci.org/gaissmai/go-inet)
[![Go Report Card](https://goreportcard.com/badge/github.com/gaissmai/go-inet)](https://goreportcard.com/report/github.com/gaissmai/go-inet)
[![Coverage Status](https://coveralls.io/repos/github/gaissmai/go-inet/badge.svg?branch=master)](https://coveralls.io/github/gaissmai/go-inet?branch=master)

A Go library for reading, formatting, sorting and converting IP-addresses and IP-blocks.

## go-inet/inet

Package inet represents IP-addresses and IP-Blocks as comparable types.

The IP addresses and blocks from this package can be used as keys in maps, freely copied and fast sorted
without prior conversion from/to IPv4/IPv6.

Some missing utility functions in the standard library for IP-addresses and IP-blocks are provided.

Blocks are IP-networks or IP-ranges, e.g.

    192.168.0.1/24              // CIDR
    ::1/128                     // CIDR
    10.0.0.3-10.0.17.134        // IP range, no CIDR
    2001:db8::1-2001:db8::f6    // IP range, no CIDR

## go-inet/tree

Tree is an implementation of a CIDR or Block tree for fast IP lookup with longest-prefix-match.
It is *NOT* a standard patricia-trie, this isn't possible for general blocks not represented as CIDRs.

The tree can be visualized as:

```
 ▼
 ├─ 10.0.0.0/9
 │  ├─ 10.0.0.0/11
 │  │  ├─ 10.0.0.0-10.0.0.29
 │  │  ├─ 10.0.16.0/20
 │  │  └─ 10.0.32.0/20
 │  └─ 10.32.0.0/11
 │     ├─ 10.32.8.0/22
 │     ├─ 10.32.12.0-10.32.13.77
 │     └─ 10.32.16.0/22
 ├─ 2001:db8:900::/48
 │  ├─ 2001:db8:900::/49
 │  │  ├─ 2001:db8:900::/52
```

## Documentation

[![GoDoc](https://godoc.org/github.com/gaissmai/go-inet?status.svg)](https://godoc.org/github.com/gaissmai/go-inet)

Full `go doc` style documentation for the project can be viewed online without
installing this package by using the excellent GoDoc site here:
http://godoc.org/github.com/gaissmai/go-inet


## Installation

```bash
$ go get -u github.com/gaissmai/go-inet/...
```
You can also view the documentation locally once the package is installed with
the `godoc` tool by running

```bash
$ godoc -http=:6060
```
and pointing your browser to
http://localhost:6060/pkg/github.com/gaissmai/go-inet

## License

go-inet is licensed under the MIT License.

