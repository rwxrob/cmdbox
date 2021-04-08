# 🍱 CmdBox Go Commander

![WIP](https://img.shields.io/badge/status-wip-red)
[![Go Report
Card](https://goreportcard.com/badge/github.com/rwxrob/cmdbox)](https://goreportcard.com/report/github.com/rwxrob/cmdbox)
[![GoDoc](https://godoc.org/github.com/rwxrob/cmdbox?status.svg)](https://godoc.org/github.com/rwxrob/cmdbox)

*A commander for modern command-line human-computer interactions.*

*CmdBox* is a lightweight commander package focused on creating
light-weight, self-contained, human-friendly terminal command-line
interfaces composed of modular subcommands that can be imported as
packages with portable completion and embedded, dynamic documentation,
and symbolic link detection. 

## Installation

Normally you would simply `import "github.com/rwxrob/cmdbox"`.

## Advantages

* Build the command-line interface to your application completely
  independently from your core API library.

* Automatic and extendible Bash tab completion built in.

* Modular, interchangeable, readable subcommands.

* Easiest possible command documentation with support for emphasis with
  simple Markdown formatting.

* Commands can be used standalone or as subcommands of a monolith (think
  BusyBox).

# Terminology

A ***command module*** is an importable Go module hosted in a Git
repository containing both an `init` function and a *standalone* command
package with a `main` function allowing it to be used by itself or
composed into another command as a subcommand.

When referring to ***standalone*** in this context we mean a *command
module* that has been compiled into a single command from its command
package with a `main` function. 

## How It Works

The magic of *CmdBox* comes from the clean separation of subcommands
into their own files --- or even separate packages imported as `_` with
their own Git repos --- and composing them into a single command index
within the main executable program. This is accomplished through simple,
judicious use of the `init()` function in each file. This approach
incurs no more performance hit than would already be required for any
other such solution and is 100% concurrency safe.

By the time the command-line application's `main()` function is called
all it has to do is `cmdbox.Execute("mycmd")` to execute the top-level
command:

1. Detects and responds to programmable shell completion context
1. Optionally adds the `help`, `usage`, and `version` builtins
1. Delegates to any subcommands, or
1. Runs its own top-level method

Because of the loose coupling, and top-down design a full subcommand
node tree can easily be displayed and documented, indeed the builtin
commands do exactly that. When care is given to maintain this loose
coupling between subcommands a file can be modularly added or moved to
any other executable directory or repo creating very clean sharing and
composition possibilities.

## Motivation

This package scratches several "personal itches" that have come up from
the needs of modern human-computer interaction and annoyances with the
popular alternatives:

1. No commanders exist for writing simple commands with human-friendly subcommands
1. No commanders exist that know anything about tab completion internally
1. Tab completion is the simplest way to provide good help to the user
1. Getopt-style options are ancient bad HCI
1. More users are voicing their commands rather than typing them
1. More users are using keyboard input with their thumbs
1. Modern command-line programs need to be easily distributed
1. Distributing executables with separate docs complicates understanding the program
1. Modern monolith approaches allow embedded documentation and better distribution
1. Tab completion code is faster and easier Go, not shell
1. Alternatives act like they have never seen `complete -C foo foo`
1. Sourcing shell completion in `bashrc` (and company) is a bloated anti-pattern
1. Operating system package managers are overkill for distributing most command executables
1. Documentation only needs to cover what the user is specifically interested in
1. Documentation can be provided dynamically when embedded within the command it documents

### Times Have Changed

The world has outgrown the original UNIX model for both command design
and documentation, presumably due to the resource constraints of the
day. These constraints forced shared library dependencies and commands
were designed to be as minimal as possible doing "only one thing really
well." Combining several actions into a single command (think `git
clone` or `go tool dist list`) or embedding full command documentation
into the executable would have been considered monstrously wasteful at
that time. The result was lots of small commands dependent on specific
shared libraries and a separate documentation system (`man` pages).

Today Go has revolutionized applications development and respectfully
embraced modern use of resources. Static linking, cross-compilation,
built in concurrency, reflection and documentation are all evidence of
how time have changed.  It follows then, that our command interface
designs and documentation approaches should equally modernize.

There is now plenty of room to compose several commands and subcommands
into a single monolithic executable. In fact, this has become idiomatic
for Go's own tool set. Distribution and portability are more important
than memory and storage size these days.

### Sometimes You Just Need Tab

Usually completion context is triggered simply by tapping tab once or
twice from the shell when typing in the command and its different
subcommand arguments. Tab completion is supported as a standard in most
all modern shells. Tab completion is often all that is needed, even
usage would be overkill. A user simply wants to remember what the
command name is.

### Keep the Docs with the Command

Clean, standardized, extensive documentation needs to travel with the
command it documents rather than being separate from it. There is no
telling what operating systems will emerge nor how (and if) they will
implement system documentation of any kind. Indeed, with the use of
modern web and git hosting documentation the need for any system-level
documentation system is sharply diminishing. Better to embed it with the
command itself when needed.

Embedded documentation has the ability to dynamically sense the
environment and context for its use thereby creating more useful,
specific help to the user. This means that rather than dumping every
single usage option into a massive man page separate from the command
(think `gpg`), such documentation can be broken up and displayed
programmatically by the command itself.

Most modern terminals support color (ANSI escapes) and UTF-8 (yes even
Windows) and regularly display more than 80 columns of text.

### Modern Human-Computer Interaction

Voice and human-friendly textual command-line interfaces are becoming
more and more popular as more humans interact with command line
interfaces in chat applications, messaging and more. This influence has
permeated to system command design and deprecated the design decision to
use traditional getopt-like switches and options, which cannot be voiced
or tab-completed easily. (Consider how difficult it is to type or voice
a simple dash into most text applications.)

Simplicity on the command-line has been an ignored requirement for some
time. Why? Because only amazing technical people are using the command
line? Perhaps. But that is a very bad reason. Interfaces need not be
overly complex or difficult to memorize just because most users are
technical.

The *CmdBox* approach takes these HCI considerations very seriously
placing priority on how humans use modern command-line interfaces and
developing them accordingly. Human interaction is the single most
important design decision for any application, including one that runs
on the command line.

## Design Decisions

This is a summary of the design decisions made roughly in the order they
were made. It is provided in the hopes of addressing other design
concerns anyone reviewing this package might have before choosing to use
it.

* To this day command options plague developers and users by different
  ways of dealing with single or double dashes, the equals sign, single
  letter options and more. Most are also not friendly to the use of
  UTF-8 runes in the upper ranges.

* Even through `getopt` can be problematic (having had to code for it
  and around it for two decades and still has nightmares about `gpg`'s
  interface) `MapOpts` is available for those who insist on using them
  giving the command author the choice. Parsing traditional getopt-type
  options and switches was originally dropped since the goals of this
  project are to deprecate such designs in favor of modern HCI
  approaches to command-line user interfaces. But, choice prevails over
  opinion in this case. Still, please don't use `-f <file>` when the
  tab-completable `file <file>` could be used instead. Even a single
  dash `-` is nearly impossible to voice-to-text command-line interface.

* Semantic emphasis `*bold*`, `**italic**`, `***bold-italic***` are the
  only inline formatting options allowed. This seems prudent given the
  fundamental requirement for readability as well as the complication of
  other inline formatting and escapes.  These three inline formats have
  been supported by all system documentation formats from the beginning
  and play nicely with Go's back-ticked raw-strings. Users can change
  the colors of these on terminals that use term(cap|info)-aware man
  pages and less/more (i.e.  `LESS_TERMCAP_??` or `COMPCMD_??`) without
  any need to provide theming in the package itself.

* The decision to not use other Markdown for formatting was easy given
  the minimal three formats allowed. The decision to keep documentation
  100% compatible with CommonMark was a no-brainer decision. 

* Pager application (i.e. less, more) detection seems practical since
  most everyone would want such when help output exceeds the page height
  (when such can be determined). There is a TODO to add fall-back
  builtin pager support.

* Allowing more than 80 columns and never hard-wrapping lines allows the
  user to decide the preferred width of documentation.  Lines that are
  hard-wrapped are unreadable when the column count drops to less than
  that of the hard-wrapping. With TMUX small pane widths are common.
  (This is one thing Go documentation seriously failed to consider.)

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

* Using `structs` instead of interfaces makes more sense given the goals
  to enable quick and easy to read and write documentation embedded in
  the source.

* Including a fair amount of output formatting and printing code seemed
  appropriate given that one of the three main goals of the project was
  to produce consistent command output formatting. It had been suggested
  to put such into another package instead but this package is so small
  that ultimately turned out to be overkill. Besides, the formatting
  used by `cmdbox` is highly specific to making output look good on a
  terminal as it relates to command documentation.

* Rather than hard-code a dependency on Bash completion, every effort
  has been made to decouple completion from any specific shell
  completion API (despite the many references to Bash, which dominates
  currently). As long as any shell completion implementation sets an
  environment variable containing the full line to be completed package
  `cmdbox` will always be able to sense completion context and perform.
  This puts the completion logic safely embedded into the command that
  needs the completion and exposes as little dependency on shell
  completion methods as possible.

* Use of aliases (`d|del|delete`) allows accessibility even if tab
  completion is not (yet) supported on a particular platform.

* Inferring the main (top-level) command to use was considered to be the
  first argument (`os.Args[0]`) for some time but further research
  revealed it can never be relied upon fully. It also initially seemed
  clever to change the behavior of an executable simply by changing its
  name but this was quickly dismissed when clearer thinking prevailed
  concerning this and other security concerns. Therefore, the first
  command must be passed as an argument to `cmdbox.Execute("mytopcmd")`. It
  is simple enough, however, to compile other executables with different
  main commands. Indeed, only the argument to need change to do so.

* Rather than add the complication of wrapping lines in a block and
  indenting the proper number of spaces for a given terminal width
  (which could be resized), the choice to keep all blocks beginning from
  the leftmost column was made. This makes the best use of screen space,
  is consistent with the help documentation from Go itself, and allows
  the output to be sized to any width without complication. Instead,
  emphasis has been given to the headers of the specific blocks. When
  indentation is truly needed raw-text (initial four spaces) can be
  used.

* Full DocOpt parsing of usage was considered and dismissed. DocOpt
  solves a different problem with different priorities (which do no
  include consideration of getopt-style command lines interfaces as an
  anti-pattern in terms of HCI). In fact, the only thing that should
  every really been in a usage string are lists of subcommands or
  arguments called out by name with `<something>` both of which are
  automatically formatted when detected. In other words, no usage should
  ever include `-` or `--` prefixed anything. *Everything* should be
  voice-able. This makes even the most complicated commands extremely
  easy to understand quickly. 

* Support for all-caps keywords (ex: `[OPTIONS]`) in usage strings was
  considered and dismissed because angle-bracket notation already covers
  that and having single suggested usage format is simpler. It is also
  far easier for the formatting parser to determine what should be
  italicized using angle-brackets.

* Decided against any paged output of any special hidden builtin to
  allow combination with other shell scripts when quick customization is
  needed. Paged output remains for the main `help`, `usage`, and
  `version` builtins however.

* Removed all color and formatting from the special hidden builtin
  output since is will mostly likely be used by shell scripts that
  further parse it, for example to email all the authors.

* `Emphasize()`, `Wrap()`, and `Indent()` have been made public in addition to
  `Format()` for convenient use by command authors, but the values for
  italic, bold, and bolditalic have not been. This is to preserve the
  look and feel of commands that use `Emphasize()` and put the power to
  control appearance into the hands of users instead of developers. 

* At one point the internal `map[string]*Command` index was private to
  discourage tight coupling between commands. But it was decided that
  this inflexibility came at too great a cost to potential needs of
  command creators in the future who might want to inspect the Index
  directly themselves without necessarily doing anything to is, say to
  use some sort of prefix convention. `Has()` was kept as a convenience.
  This does not change the fact that subcommands should be as independent
  and uncoupled as possible. In the extreme case when subcommands need
  to communicate they should use system environment variables as would
  any other two commands normally.


## How Does Completion Work?

Reading about [Bash Programmable Completion](https://duck.com/lite?kae=t&q=Bash Programmable Completion) is probably a good idea.

For Bash all you need to add to your `~.bashrc` or `~/.bash_completion` file is the following:

```bash

complete -C mycmd mycmd

```

This will cause the shell to run `mycmd` and set the `COMP_LINE`
environment variable every time you tap the tab key once or twice. This
allows *Complete Commander* to detect completion context and only print
the words that should be possible for the last word of the command when
the tab key was pressed.

The `cmdbox` package then sees completion context it resolves it by calling
Complete(). See the package docs for more on the specific algorithm
used, but generally it does the following:

* Prints the output of any completion function found, or
* Recursively looks up any matching Subcommand names and prints them

## Machine Learning in Simple Terminal Commands?

Yep. Allowing a completion function allows incredible interesting
possibilities with completion that uses full human language semantics
rather than just prefixing the most likely candidates. Even a full
machine learning code module could be added allowing any possible
speech. Such considerations seem very absent from the HCI conversation
regarding terminal command line usage.

## Automatic (Builtin) Subcommands

Unless specifically disabled with `OmitBuiltins` (although the size to
add these builtins is trivial), the following internal subcommands are
added to any executable that is created using the `cmdbox` package:

* `help [<subcmd>]`
* `usage`
* `version`

These subcommands are so common they have become something of an
unspoken standard for modern commands. 

Several other utility subcommands are also builtin for help integrating
with shell environments, producing documentation in other formats, and
more:

```

_authors       list names and authors
_bash_complete print line to add for bash completion
_builtins      list all cmd package builtins names and summaries
_cmdversion    print the cmd package version
_complete      force completion context
_copyrights    list names and copyrights
_descriptions  list names and descriptions
_examples      list names and examples
_gits          list names and git source repos
_help_json     dump help documentation as JSON
_index         list all names and summaries from cmd package index
_issues        list names and issue reporting URLs
_licenses      list names and licenses
_names         list names, main first
_summaries     list names and summaries
_usages        list names and usages
_versions      list names and versions

```

### Help Documentation

Help documentation is inspired by the look and design of UNIX `man`
pages so as to feel comfortable to those using such documentation for
commands for decades. It contains all the details for the given command
or `<subcmd>`. Rather than dump all the documentation into a single
page, however, details of subcommands can be displayed separately.

### Usage Documentation

Usage is when double-tab doesn't provide enough hints about how the
command is to be used. Usage is meant to provide only minimal usage
output while 'help' provides the full detailed information for the
command.  If an additional argument is provided the detailed help for
that specific subcommand will be provided. If the subcommand does not
exist it will be ignored and the main help information shown instead.

### Version Documentation

Having a `version` subcommand in particular is well-defined as being the
place to put all legal and authorship information in addition to just
the version. Such is required by most all free software and open-source
licenses.

## Does this have anything to do with BusyBox?

*CmdBox* originally started as `cmdtab` and was used for more than two
years before Rob realized [BusyBox](https://busybox.net) was even a
thing. The name was changed when it became clear just how close to
BusyBox CmdBox is in approach and purpose, even if the design and
implementation are completely different.

Here are the similarities:

* Recognize symbolic links
* Light-weight
* Can be used as a shell (of sorts)
* Facilitates small containers 

Here are the major differences:

* Written in Go
* Modular and composable
* Embedded documentation 
* Support for standalone commands
* No symbolic links required

## What about Cobra?

While Cobra is popular and Steve and the rest of the contributors are
wonderful, talented people. Cobra has codified several anti-patterns
into rather large applications that have spread like cancer. The most
diseased of these is the use of several hundred lines of sourced,
insecure, poorly written tab-completion shell code forced on anyone who
wants completion for *any* command that uses Cobra. This could have been
accomplished with the simpler, safer, and faster `complete -C foo foo`
variation with *zero* shell sourcing but it's too late now for Cobra to
turn back. Something new is needed.

Beyond that, Cobra missed an amazing opportunity to create a modular
system leveraging `init()` side-effect imports (`_`). CmdBox takes full
and appropriate advantage of these amazing Go features. Cobra's syntax
for `init()` is anything *but* clean. CmdBox `init()` bodies are sexy,
simple, and self-documenting. In fact, they usually *contain* the
documentation itself in readable, simple Markdown that is rendered much
like an embedded "man" page.

Cobra rendered documentation is ugly, unclear, and misses the
opportunity to allow the judicious use of color and local pager
applications when detected. Cobra applications therefore render
difficult to read, unformatted help documentation.

## Terminology

* **tool** - a CmdBox command

## Conventions

By convention CmdBox commands are called *tools* (as in that would go
into a *tool box*) and are always prefixed with `cmdbox-` for easy
discovery on GitHub, etc.

## Real World Examples

* <https://github.com/rwxrob/kn>
* <https://github.com/rwxrob/cmdbox-pomo>
* <https://github.com/rwxrob/cmdbox-config>
* <https://github.com/rwxrob/cmdbox-twitch>
* <https://github.com/rwxrob/auth-go/cmd/auth>

