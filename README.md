# 🍱 Go cmdbox Composite Commander

***💥 DEPRECATED: This project long been superseded by [bonzai](https://github.com/rwxrob/bonzai). This repo is only kept for historical reference.***

*A composite commander for modern human-computer text interactions.*

> "It's like a modular, multicall BusyBox builder for Go with built in
> completion and embedded documentation support."

> "The utility here is that cmdbox lets you maintain your own personal
> 'toolbox' with built in auto-complete that you can assemble from
> various Go modules. Individual commands are isolated and unaware of
> each other and possibly maintained by other people." (tadasv123)

*cmdbox* is a lightweight commander for creating commands composed of
modular subcommands that can be imported as packages with portable
completion and embedded documentation compiled into a single, minimal,
multicall binary.

## Installation

Normally you would simply `import "github.com/rwxrob/cmdbox"`.

## `cmdbox.DEBUG = true`/`CMDBOX_DEBUG=1` Mode

Developers can set the `CMDBOX_DEBUG` environment variable to anything
or the `cmdbox.DEBUG` package boolean to `true` to activate a verbose
log written to standard error output designed to make cmdbox command
modules easier to develop and debug. This is particularly important
since traditional Go unit test methods are somewhat more difficult when
dealing with things that happen during `init()` time.

Also consider using `cmdbox.Init()` during testing to reset the package
variable state between synchronous tests.

## Advantages

* Build the command-line interface to your application completely
  independently from your core API library

* Automatic and extendible bash tab completion built in

* Modular, interchangeable, readable subcommands

* Zero dependencies on anything but the Go 1.17+ standard library

* Never any error written to anything but stderr (`log.Printf`)

* Commands can be used standalone or as multicall-enabled subcommands of
  a composite binary (think BusyBox)

* Rich ecosystem of importable `cmdbox-*` commands

# Terminology

A ***command module*** is an importable Go module hosted in a Git
repository containing both an `init` function and a *standalone* command
package with a `main` function allowing it to be used by itself or
composed into another command as a subcommand.

When referring to ***standalone*** in this context we mean a *command
module* that has been compiled into a single command from its command
package with a `main` function. 

## How It Works

The magic of *cmdbox* comes from the clean separation of *Commands* into
their own files --- or even separate packages imported as "side-effects"
(`_`) with their own Git repos --- and composing them into a single
multicall binary that contains an internal register of these Commands.
This is accomplished through simple, judicious use of the `init()`
function. This approach incurs no more performance hit than would
already be required for any other such solution and is 100% safe for
concurrency.

By the time the command-line application's `main()` function is called
all it has to do is `cmdbox.Execute()` the top-level command. A typical
`main.go` needs no more than three lines:

```go
package main
import "github.com/rwxrob/cmdbox"
func main() { cmdbox.Execute() }
```

The name of the main command may be passed to `Execute("mymain")` or can
be left off and inferred from the name of the binary instead.

The bulk of the code is kept in one or more other modular Command files
such as this sample (the package name usually does not matter, `cmd` is
used by convention):

```go

package cmd

import (
    "fmt"
    "github.com/rwxrob/cmdbox"
)

func init() {
    x := cmdbox.Add("greet")
    x.Summary = `print a polite greeting`
    x.Usage = `[NAME]`
    x.Copyright = `Copyright 2021 Rob Muhlestein`
    x.License = `Apache-2`
    x.Version = `v1.0.0`
    x.AddHelp()
    // ...
    x.Method = func(args ...string) error {
        name = "you"
        if len(args) > 0 {
            name = args[0]
        }
        fmt.Printf("Well, hello there, %v!\n",name)
        return nil
    }
}

```

This code can now be used *both* as a standalone `foo` program or
imported by other composites with `import
"github.com/rwxrob/cmdbox-greet"` in one of its own Command files (or
just the main file).

This modularity allows Commands to freely be exchanged and migrated
between different projects and repos, or just factored out entirely and
imported freely by anything. If there are naming conflicts, the
cmdbox.Rename provides for easy, explicit renames as needed.

The `x.AddHelp()` helper method will add a default `h|help` command that
includes a summary of the legal information.

## Motivation

This package scratches several "personal itches" that have come up from
the needs of modern human-computer interaction and annoyances with the
popular alternatives:

1. No commanders exist for writing simple interfaces (no dashes)
1. Separating documentation from a program is a modern anti-pattern
1. No commanders exist with internal tab completion
1. Sourcing shell completion is a bloated, modern anti-pattern
1. `kubectl` requires sourcing 12637 lines in `.bashrc` for completion
1. Applications with cmdbox require 1 line (`complete -C foo foo`)
1. Completion resolved by the application itself is more powerful
1. Tab completion is the simplest way to provide good help to the user
1. Tab completion resolution is faster and easier in Go
1. Getopt-style options have always been bad UI/HCI
1. Modern command-line interfaces depend more on stateful context
1. Modern command-line programs need to be easily distributed

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

Today the Go programming language has revolutionized applications
development and respectfully embraced modern use of resources. Static
linking, multicall binaries, in which an entire embedded filesystem can
be included, and cross-compilation, built in concurrency, reflection and
documentation are all evidence of how times have changed. It follows
then, that our command interface designs and documentation approaches
should be equally modernized. There is now plenty of room to compose
several commands and subcommands into a single, composite, multicall
binary executable. In fact, this has become idiomatic for Go's own tool
set and popularized by the container-world's love for BusyBox.
Distribution and portability are more important than memory and storage
size these days.

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

The *cmdbox* approach takes these HCI considerations very seriously
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
  years. Developers can now create "cool" additions by implementing
  their own `cmdbox-*` packages and leaving the core `cmdbox` package
  alone. Such an ecosystem is far more beneficial to the community as a
  whole.

* Any shell or system that sets the `COMP_LINE` environment and allows
  the targeting of the same command being called to complete itself (ex:
  `complete -C foo foo`) is supported. This obviously includes only bash
  at the moment, but could easily be added to other existing and future
  shells. This established mechanism for communicating *completion
  context* is not only mature but simple and trivial to implement. 

* Aliases for Commands can be used for alternative languages as well
  providing multi-lingual speakable command line interface
  possibilities.

* Centralizing messages in a `map[string]string` allows them to be
  changed easily by those wishing to support more than the
  English language in their composite programs.

* Dashes on the command line have always been a bad idea, even more so
  now. The world has suffered enough from dubious design decisions made
  decades ago in the `getops` world when string sizes were limited and
  usernames no more than eight characters. Thankfully, we can now
  dismiss these bad practices. As Unicode has standardized, and
  subcommand composites such at `git`, `kubectl`, `docker`, `gh` and
  others like them have become standard practice --- along with the rise
  of conversational interfaces that transform any command line into a
  speakable user interface --- it is time we started giving our programs
  support for written natural language grammars. Not only will our code
  and command lines be more readable but more accessible allowing the
  tried-and-true command line interface model to extend to chat
  services, conversational assistants, and voice-enabled devices. In
  fact, cmdbox is ideal for creating modular (but monolithic) chat bots
  of all kinds.

* Originally, a custom, Markdown-derived syntax was supported for the
  text of all Command fields. This has been removed for now and will be
  reconsidered at a later date with proposals from the community. For
  now, the convention is to use nothing but regular text with no attempt
  to provide any type of formatting markup. It is critical that the
  cmdbox community agree on a markup standard, if any. While cmdbox
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
  that encapsulate them with `Call()` or the commands themselves can be
  front-end shells to high-level library API calls.

* Using `structs` instead of `interfaces` makes more sense given the
  serialization goals to enable quick and easy to read and write
  documentation embedded in the source.

* Use of aliases (`d|del|rm|delete`) allows accessibility even if tab
  completion is not (yet) supported on a particular platform. This
  provides minimal multi-language support possibilities as well.

* At one point the internal CommandMap (`map[string]*Command`) index was
  private to discourage tight coupling between commands and prevent
  "innovation" that could fork away simplicity as a core tenet of the
  cmdbox project. But it was decided that this inflexibility came at too
  great a cost to potential needs of command creators and `cmdbox.Reg()`
  was added. 

* Naming of `cmdbox-*` module repos allows for easy discovery. Using the
  conventional `cmd`/`cmd.go` package/file name (which is ignored at
  `init` time) allows consistency for generators and cataloging.

* YAML support was dropped in v0.7.7 after discovering that the YAML
  project team has chosen to break compatibility between v2 and v3 (a
  source of frustration for the Kubernetes project). Dropping YAML
  support allows developer teams to make their own YAML decisions and
  removes the only external package dependency cmdbox had up to that
  point. JSON with two space indent will be the standard for marshaling
  with RawJSON being available for those who wish JSON on a single line.
  The main JSON parsing functions now return an error as well with
  MustJSON/MustRawJSON variations that return an empty string if there
  is any error and log the error to standard error.

* All errors have been moved to output to standard error instead of
  stdout so they never conflict with using cmdbox apps as filters.

## How Does Completion Work?

Please [Bash Programmable
Completion](https://duck.com/lite?kae=t&q=Bash%20Programmable%20Completion) and pay specific attention to the `complete -C`
option. 

For Bash all you need to add to your `~/.bashrc` file is the following: 

```bash

complete -C foo foo

```

This will cause the shell to run `foo` and set the `COMP_LINE`
environment variable every time you tap the tab key once or twice. This
allows a cmdbox composite program to detect completion context and only
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
other innovations are bringing the focus (rightly) back to speakable
user interfaces of all kinds. 

## Does this have anything to do with BusyBox?

*cmdbox* originally started as `cmdtab` and was used for more than two
years before Rob realized [BusyBox](https://busybox.net) was even a
thing. The name was changed when it became clear just how close to
BusyBox cmdbox is in approach and purpose, even if the design and
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

Nothing wrong with that if you are okay with dashes and such. cmdbox is
designed for more.

## What about Cobra?

*Cobra is the Electron of command-line interface libraries.*

While Cobra is popular and Steve and the rest of the contributors are
wonderful, talented people. Cobra has codified several anti-patterns
into rather large applications that have spread like cancer wasting
resources and breaking `.*rc` files with wild abandon. For example,
`kubectl` requires the sourcing of 12637 lines of bash code every time a
new shell spawns, and that is only one Cobra application. There are many
others.

Considering that Cobra is used for most cloud-native related
applications it is not hyperbole to say that tens of thousands of people
are evaluating 50,000+ lines of code every time they start a new Bash
shell. 

While there are ways around this, the simplest has been completely
overlooked by the Cobra project: `complete -C foo foo`. The project now
maintains so much technical debt because of these hastily made
early design decisions in the race to be the first Go Commander that
something radically new is needed.

Beyond that, Cobra missed an amazing opportunity to create a modular
ecosystem of composable commands by leveraging `init()` side-effect
imports (`_`). cmdbox takes full and appropriate advantage of these
amazing Go features. Cobra's syntax for `init()` is anything *but*
clean. cmdbox `init()` bodies are sexy, simple, and self-documenting. In
fact, they usually *contain* the documentation itself in readable,
simple text that is rendered much like an embedded "man" page.

## Terminology

* **command** - `foo` or `mycmd foo`
* **tool** - another name for *command*
* **module** - a Go module containing cmdbox
* **composite** - final binary composed of multiple commands

## Conventions

* Prefix repos with `cmdbox-` for easy discovery.
* Use `x` for Command pointer `func init() {x := cmdbox.Add("name")}`
* Reuse `x` when adding multiple commands in same `init()`
* Use `x` for Command method receivers `func (x *Command) ...`

## Real World Examples

* <https://github.com/rwxrob/cmdbox-pomo>
* <https://github.com/rwxrob/cmdbox-config>

## Legal 

Copyright 2021 Robert S. Muhlestein (<mailto:rob@rwx.gg>)  
Licensed under Apache-2.0

"cmdbox" and "cmdbox" are legal trademarks of Robert S. Muhlestein but
can be used freely to refer to the cmdbox project
<https://github.com/rwxrob/cmdbox> without limitation. To avoid
potential developer confusion, intentionally using these trademarks to
refer to other projects --- free or proprietary --- is prohibited.
