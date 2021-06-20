# 🍱 Go CmdBox Composite Commander

[![GoDoc](https://godoc.org/github.com/rwxrob/cmdbox?status.svg)](https://godoc.org/github.com/rwxrob/cmdbox)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rwxrob/cmdbox)](https://goreportcard.com/report/github.com/rwxrob/cmdbox)
[![Coverage](https://gocover.io/_badge/github.com/rwxrob/cmdbox)](https://gocover.io/github.com/rwxrob/cmdbox)

*A composite commander for modern human-computer text interactions.*

> "It's like a modular, multicall BusyBox builder for Go with built in
> completion and multi-lingual, embedded documentation support."

*CmdBox* is a lightweight commander package focused on creating
light-weight, self-contained, human-friendly terminal command-line
interfaces composed of modular subcommands that can be imported as
packages with portable completion and embedded, dynamic documentation
compiled into a single, minimal, multicall binary.

## Installation

Normally you would simply `import "github.com/rwxrob/cmdbox"`.

## Advantages

* Build the command-line interface to your application completely
  independently from your core API library.

* Automatic and extendible Bash tab completion built in.

* Modular, interchangeable, readable subcommands.

* Commands can be used standalone or as multicall-enabled subcommands of
  a composite binary (think BusyBox).

* Rich ecosystem of importable `cmdbox-*` commands.

# Terminology

A ***command module*** is an importable Go module hosted in a Git
repository containing both an `init` function and a *standalone* command
package with a `main` function allowing it to be used by itself or
composed into another command as a subcommand.

When referring to ***standalone*** in this context we mean a *command
module* that has been compiled into a single command from its command
package with a `main` function. 

## How It Works

The magic of *CmdBox* comes from the clean separation of *Commands*
into their own files --- or even separate packages imported as `_` with
their own Git repos --- and composing them into a single multicall
binary that contains an internal register of these Commands. This is
accomplished through simple, judicious use of the `init()` function in
each file. This approach incurs no more performance hit than would
already be required for any other such solution and is 100% safe for
concurrency.

By the time the command-line application's `main()` function is called
all it has to do is `cmdbox.Execute()` to execute the top-level
command. A typical `main.go` needs no more than three lines:

```go
package main
import "github.com/rwxrob/cmdbox"
func main() { cmdbox.Execute() }
```

The bulk of the code is kept in one or more other modular Command files
such as this sample (the package name almost never matters):

```go

package cmd

import (
    "fmt"
    "github.com/rwxrob/cmdbox"
    _ "github.com/rwxrob/cmdbox-help"
    _ "github.com/rwxrob/cmdbox-version"
)

func init() {
    x := cmdbox.Add("foo","h|help","v|version")
    x.Summary = `foo the things`
    x.Usage = `[h|help|v|version]`
    x.Copyright = `Copyright 2021 Rob Muhlestein`
    x.License = `Apache-2`
    x.Version = `v1.0.0`
    // ...
    x.Method = func(args []string) error {
        fmt.Println("would foo all the things")
        return nil
    }
}

```

This code can now be used *both* as a standalone `foo` program or
imported by other composites with `import
"github.com/rwxrob/cmdbox-foo"` in one of its own Command files (or just
the main file).

This modularity allows Commands to freely be exchanged and migrated
between different projects and repos, or just factored out entirely and
imported freely by anything. If there are naming conflicts, the
cmdbox.Rename provides for easy, explicit renames as needed.

## Motivation

This package scratches several "personal itches" that have come up from
the needs of modern human-computer interaction and annoyances with the
popular alternatives:

1. No commanders exist for writing simple, speakable interfaces
1. No commanders exist that know anything about tab completion internally
1. Tab completion is the simplest way to provide good help to the user
1. Getopt-style options are ancient, bad HCI
1. More users are voicing their commands rather than typing them
1. More users are using keyboard input with their thumbs
1. Modern command-line programs need to be easily distributed
1. Separating docs from the program is a modern anti-pattern
1. Tab completion resolution is faster and easier in Go
1. Alternatives act like they have never seen `complete -C foo foo`
1. Sourcing shell completion is a bloated, modern anti-pattern
1. Why create a package when a single executable will do?
1. Embedded, multi-lingual documentation is non-existent today

### Times Have Changed

The world has outgrown the original UNIX model for both command design
and documentation, presumably due to the resource constraints of the
day. These constraints forced shared library dependencies and commands
were designed to be as minimal as possible doing "only one thing really
well." Combining several actions into a single command (think `git
clone` or `go tool dist list`) or embedding full command documentation
into the executable would have been considered monstrously wasteful at
that time. The result was lots of small commands dependent on specific
shared libraries and separate systems both for documentation (`man`
pages) and tab completion (`/etc/bash_complete`). This forced software
delivery to invent the package management systems we see today. Shipping
everything in a single executable was simply unthinkable.

Today Go has revolutionized applications development and respectfully
embraced modern use of resources. Static linking, multicall binaries,
cross-compilation, built in concurrency, reflection and documentation
are all evidence of how time have changed. It follows then, that our
command interface designs and documentation approaches should be equally
modernized. There is now plenty of room to compose several commands and
subcommands into a single, composite, multicall binary executable. In
fact, this has become idiomatic for Go's own tool set and popularized by
the container-world's love for BusyBox. Distribution and portability are
more important than memory and storage size these days.

### Sometimes You Just Need Tab

Usually completion context is triggered simply by tapping tab once or
twice from the shell when typing in the command and its different
subcommand arguments. Tab completion is supported as a standard in most
all modern shells. Tab completion is often all that is needed, even
usage would be overkill. A user simply wants to remember what the
command name is.

Not only is tab completion desirable, it has become the preferred method
of containing complexity. The `git` application, for example, has very
precise rules governing its tab completion based on the subcommand being
called. Pressing tab from `git branch` will only show you the current
branches.

### Keep the Docs with the Command

Clean, standardized, extensive documentation needs to travel with the
command it documents rather than being separate from it. There is no
telling what operating systems will emerge nor how (and if) they will
implement system documentation of any kind. Indeed, with the use of
modern web, containerization, and git hosting documentation the need for
any system-level documentation system is sharply diminishing. Better to
embed it with the command itself when needed. When building container
`FROM scratch` binaries this enables documentation without bloating the
container up to contain the entire `man` page system. The same is true
for tab completion itself.

### Modern Human-Computer Interaction

Voice and human-friendly textual command-line interfaces are becoming
more and more popular as more humans interact with command line
interfaces in chat applications, messaging and more. This influence has
permeated to system command design and deprecated the design decision to
use traditional getopt-like switches and options, which cannot be voiced
or tab-completed easily. (Consider how difficult it is to type or voice
a simple dash into most text applications.)

Simplicity on the command-line has been an ignored requirement for some
time. 

Why?

Because only amazing technical people are using the command line?

Perhaps. But that is a very bad reason. Interfaces need not be overly
complex or difficult to memorize just because most users are technical.

The *CmdBox* approach takes these HCI considerations very seriously
placing priority on how humans use modern command-line interfaces and
developing them accordingly. Human interaction is the single most
important design decision for any application, including one that runs
on the command line --- *especially* in the modern world of
conversational assistants and chat bot plugins.

## Design Decisions

This is a summary of the design decisions made over the course of the
project (in no particular order). It is provided in the hopes of
addressing other design concerns anyone reviewing this package might
have before choosing to use it.

* Several "cool" things have been dropped. This has resulted in a code
  base that seems deceptively simple, even trivial. That is by design,
  and came after many hours of refactoring different approaches over the
  years. Developers can now create "cool" additions by implementing their
  own `cmdbox-*` packages and leaving the core `cmdbox` package alone.
  Such an ecosystem is far more beneficial to the community as a whole.

* Any shell or system that sets the `COMP_LINE` environment and allows
  the targeting of the same command being called to complete itself (ex:
  `complete -C foo foo`) is supported. This obviously includes only Bash
  at the moment, but could easily be added to other existing and
  future shells. This established mechanism for communicating *completion
  context* is not only mature but simple and trivial to implement.
  The upcoming `cmdbox-ash` CmdBox command module will provide a POSIX
  compliant shell with CmdBox compatible tab completion (as inspired by
  BusyBox).

* Aliases for Commands can be used for alternative languages as well
  providing multi-lingual speakable command line interface
  possibilities.

* Centralizing messages in a `map[string]string` allows them to be
  changed easily by those wishing to support more than the
  English language in their composite programs.

* Dashes on the command line have always been a bad idea, even more so
  now. The world has suffered enough from dubious design decisions made
  decades ago in the `getops` world when string sizes were limited and
  usernames no more than eight characters. Beginning anything but
  negative numbers with dashes on the command line has *always* been an
  HCI UX anti-pattern that we tolerated when we had to, expanded on when
  we could (with "long" options), and can now gratefully dismiss. As
  Unicode has standardized, and subcommand composites such at `git`,
  `kubectl`, `docker`, `gh` and others like them have become standard
  practice --- along with the rise of conversational interfaces that
  transform any command line into a speakable user interface --- it is
  time we started giving our programs support for written natural
  language grammars. Not only will our code and command lines be more
  readable but more accessible allowing the tried-and-true command line
  interface model to extend to chat services, conversational assistants,
  and voice-enabled devices. 

* At one point the base `cmdbox` package contained a "builtin" `help`
  and `version` subcommand but they were both factored into the
  `cmdbox-help` package instead. This reduces potential bloat for
  composites that do not need them and better provides for independent
  multi-lingual support. 

* Originally a custom, Markdown-derived syntax was supported for the
  text of all Command fields. This has been removed for now and will be
  reconsidered at a later date with proposals from the community. For
  now, the convention is to use nothing but regular text with no attempt
  to provide any type of formatting markup. It is critical that the
  CmdBox community agree on a markup standard, if any. While CmdBox
  plans to improve on the austerity of Go doc (allowing for color,
  emphasis, and such) this has to be considered as carefully as the base
  `cmdbox` API itself.

* Internationalization was a big design consideration from the
  beginning. All code containing the English defaults is separate from
  the rest allowing build tools to generate language specific versions.
  The idea of putting all supported languages into monolithic files with
  locale templates was considered but quickly dismissed given the
  potential impact to executable size. Localized builds are a common
  industry practice. The English naming of the default builtin commands
  is at least as appropriate as these standard keywords are included in
  many other contexts `help`, `version`, and `usage` are very
  ubiquitous. If needed, these can be aliased easily by adding commands
  that encapsulate them with `Call()`.

* Using `structs` instead of `interfaces` makes more sense given the
  serialization goals to enable quick and easy to read and write documentation embedded in the source.

* Use of aliases (`d|del|rm|delete`) allows accessibility even if tab
  completion is not (yet) supported on a particular platform. This
  provides multi-language support possibilities as well.

* At one point the internal `map[string]*Command` index was private to
  discourage tight coupling between commands and prevent "innovation"
  that could fork away simplicity as a core tenet of the CmdBox project.
  But it was decided that this inflexibility came at too great a cost to
  potential needs of command creators and `cmdbox.Reg()` was added. 

* Naming of `cmdtab-*` module repos allows for easy discovery. Using the
  conventional `cmd` package name (which is ignored at `init` time)
  allows consistency for generators and cataloging.

## How Does Completion Work?

Please [Bash Programmable Completion](https://duck.com/lite?kae=t&q=Bash
Programmable Completion) and pay specific attention to the `complete -C`
option. 

For Bash all you need to add to your `~/.bashrc` file is the following: 

```bash

complete -C foo foo

```

This will cause the shell to run `foo` and set the `COMP_LINE`
environment variable every time you tap the tab key once or twice. This
allows a CmdBox composite program to detect completion context and only
print the words that should be possible for the last word of the command
when the tab key was pressed. (Also see Go documentation of the
Command.Complete method.)

## Machine Learning in Simple Terminal Commands?

Yep. Allowing a completion function allows incredible interesting
possibilities with completion that uses full human language semantics
rather than just prefixing the most likely candidates. Even a full
machine learning code module could be added allowing any possible speech
completion. Such considerations seem very absent from the HCI
conversation regarding terminal command line usage in the past, but
times are changing. Phone, OpenCV, natural language processing and many
other innovations are bringing the focus (rightly) back to speakable user
interfaces of all kinds. 

## Does this have anything to do with BusyBox?

*CmdBox* originally started as `cmdtab` and was used for more than two
years before Rob realized [BusyBox](https://busybox.net) was even a
thing. The name was changed when it became clear just how close to
BusyBox CmdBox is in approach and purpose, even if the design and
implementation are completely different.

Here are the similarities:

* Multicall binaries (hard and symbolic links)
* Light-weight
* Can be used as a shell (of sorts)
* Facilitate small containers 

Here are the major differences:

* Written in Go
* Modular and composable
* Rich ecosystem of importable commands
* Embedded documentation 
* Support for standalone commands
* No symbolic links required

## Why Not Just `import flag`?

Nothing wrong with that if you are okay with dashes and such. CmdBox is
designed for more.

## What about Cobra?

While Cobra is popular and Steve and the rest of the contributors are
wonderful, talented people. Cobra has codified several anti-patterns
into rather large applications that have spread like cancer. 

Consider that the output of `kubectl completion bash` is more than
14,000 lines long and that the project recommends "sourcing" or
"evaling" all of them from your `.bashrc` file. That is just a single
project that uses Cobra. Considering that Cobra is used for most
cloud-native related applications it is not hyperbole to say that tens
of thousands of people are evaluating 50,000+ lines of code every time
they start a new Bash shell. 

While there are ways around this, the simplest has been completely
overlooked by the Cobra project: `complete -C foo foo`. The project now
maintains so much technical debt because of these hastily made
early design decisions in the race to be the first Go Commander that
something radically new is needed.

Beyond that, Cobra missed an amazing opportunity to create a modular
ecosystem of composable commands by leveraging `init()` side-effect
imports (`_`). CmdBox takes full and appropriate advantage of these
amazing Go features. Cobra's syntax for `init()` is anything *but*
clean. CmdBox `init()` bodies are sexy, simple, and self-documenting. In
fact, they usually *contain* the documentation itself in readable,
simple text that is rendered much like an embedded "man" page if the
`cmdtab-help` module is imported, or left out entirely for ultra-light
container-friendly composite commands, cousins of BusyBox.

## Terminology

* **command** - `foo` or `mycmd foo`
* **tool** - another name for *command*
* **module** - a Go module containing CmdBox
* **composite** - final binary composed of multiple commands

## Conventions

* Prefix repos with `cmdbox-` for easy discovery.
* Use `x` for Command pointer `func init() {x := cmdbox.Add("name")}`
* Use `x` for Command method receivers `func (x *Command) ...`

## Real World Examples

* <https://github.com/rwxrob/kn>
* <https://github.com/rwxrob/cmdbox-pomo>
* <https://github.com/rwxrob/cmdbox-config>
* <https://github.com/rwxrob/cmdbox-twitch>
* <https://github.com/rwxrob/auth-go/cmd/auth>

## Legal 

Copyright 2021 Robert S. Muhlestein (<mailto:rob@rwx.gg>)  
Licensed under Apache-2.0

"CmdBox" and "cmdbox" are legal trademarks of Robert S. Muhlestein but
can be used freely to refer to this specific project without limitation.
To avoid potential developer confusion, intentionally using these
trademarks to refer to other projects --- free or proprietary --- is
prohibited.
