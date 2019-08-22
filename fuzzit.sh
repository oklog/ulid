#!/bin/bash
set -xe

NAME=ulid
TYPE=$1

function fuzz {
    TARGET=$NAME-$1
    FUNCTION=Fuzz$2
    go-fuzz-build -libfuzzer -func $FUNCTION -o $NAME.a .
    clang -fsanitize=fuzzer $NAME.a -o $NAME
    ./fuzzit create job --type $TYPE $TARGET $NAME
}

# Setup
export GO111MODULE="off"
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
go get -d -v -u ./...
wget -q -O fuzzit https://github.com/fuzzitdev/fuzzit/releases/download/v2.4.29/fuzzit_Linux_x86_64
chmod a+x fuzzit

# Fuzz
fuzz new New
fuzz parse Parse
fuzz parse-strict ParseStrict
