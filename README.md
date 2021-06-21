# diceware

Simple command line interface password generator based at [this](//diceware.dmuth.org).

## Motivation

I created it, because CLI interface is more suitable for generating passwords for me than browser.

## Installation

Install go tool chain with:

    $ brew install go

on macOS or:

    $ sudo dnf -y install go

on Fedora. I can bet that you can install go compiler on other linux distros with you favourite package manager.

Create directory for your go binaries and libraries in your home directory.

    $ mkdir $HOME/go

Then add this environmental variables to your `.zshrc` or `.bashrc`.

```sh
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
```

Now you can clone this repository and build __diceware__ with following commands.

    $ git clone https://github.com/thinkofher/diceware.git
    $ cd diceware
    $ go build ./cmd/diceware

If you want to install __diceware__, just execute below command.

    $ go install ./cmd/diceware

Now __diceware__ binary is in your `$GOPATH/bin` directory, which is also added to your `$PATH`
environmental variable, which means you can execute `diceware` like your other unix programs
in your shell environment.

Have fun!

## Usage

It is simple as:

    $ diceware
    QuackLassoMatadorAppealingMobilize

If you want to have more or less words in your password, just pass number of words to `-words` flag.

    $ diceware -words 8
    ClericalCrankUnmanagedUnthawedEndnoteGigglingUncouthNativity

You can generate more than one password with `-passwords` flag.

    $ diceware -words 8 -passwords 3
    RescuerStatisticLungHelperGarlicBruteClothesOutdoors
    CranberryExpansionHuddlingStopperPrecinctExtraditeAnimatingCarless
    XeroxIguanaBaffleRockstarBossMumpsUnfailingRecognize

You can use `-w` and `-p` shorthands for `-words` and `-passwords` flags and `-help` flag to output
above information in more simpler way.

That's all folks! Thank you for your attention.

## Security

`diceware` uses [crypto/rand](https://golang.org/pkg/crypto/rand/) package for generating random
numbers, which uses _"global, shared instance of a cryptographically secure random number generator"_ of your OS.

So we can assume it's safe.
