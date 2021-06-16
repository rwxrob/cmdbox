/* Package cmdbox is a multicall, modular commander with embedded tab
completion and locale-driven documentation, that prioritizes modern,
speakable human-computer interactions from the command line.

    package mycmd

    import "github.com/rwxrob/cmdbox"

    func init() {
      x := New("foo","subcmd")
      x.Summary = `foo is a thing`
      // ...
    }

First Argument is Command Name

The first argument is *always* the main Command.Name by which it will be
called.  This uniquely identifies the Command and becomes the key used
by Call to lookup the Command from the internal Register (Reg) for
completion and execution.  By convention, these must be speakable,
complete words with absolutely no punctuation whatsoever.  (For
performance reasons, no validation is performed on the Name.)

The Command.Name may contain two complete words separated by a single
space. This is to avoid collisions and facilitate default tab
completion. It also removes indirection when called from Execute.

    x := New("foo help")
    x.Summary = `output help information for foo`

Using two-word names is common when packaging subcommands with commands
in such a way as to disambiguate which subcommand is wanted --
particularly when common words are used.

    x := New("foo help")

(Note, however, that the default cmdbox-help package is able to detect
the command for which help is needed through the caller passed to Call.
This example assumes wanting to override that, if used.)

Remaining Arguments are Subcommands

Any variadic arguments that follow will be directly passed to
Command.Add(). This provides a succinct summary of how the command may be
called. The first in the list will be assigned to Command.Default, but
can be overridden with a direct assignment later.

    x := New("foo", "help")
    // x.Default == "help"

Each argument passed in the list may be in signature form (rather than
just a name) meaning it may have one or more aliases prefixed and
bar-delimited which are added to the Commands map:

    x := New("foo", "h|help")
    // x.Default == "help"
    // x.Commands == {"h":"help","help":"help"}

When the Command is called by Call the Commands map is used to delegate
the call to a matching Command in the internal Register (Reg) if and
only if the Command itself does not have a Command.Method defined. See
Call for more about this delegation on how it finds key name matches in
the internal Register (Reg).

Command Method Has Priority

All but top-level Commands will usually assign a Command.Method to
handle the work of the Command. By convention the arguments should be
named "args" and no name given to the error returned:

    x.Method = func(args []string) error {
        fmt.Println("would do something")
        return nil
    }

If a Command has a Method, then Call will pass all arguments as-is
allowing the Method to decide if they just arguments or keywords for
actions to be handled within that Command.Method (usually within
a switch/case block). The Method may still cmdbox.Call() to delegate to
other other Commands in Register (but avoid unnecessary coupling between
Commands when possible. See Call for more.)

No Command Method Will Delegate

If the Command does not have a Method of its own, then the list of
arguments passed to New is assumed to be the signatures for other
Commands in the Register that must eventually be populated by other
Command init() functions including subcommands of the given Command.

Note that New does no validation of any potential command in the
internal Register (Reg) because the state of the Register cannot be
predicted at init() time. Not all Commands may yet have been registered
before any other cmdbox.New is called. This means runtime testing is
required to check for errant calls to unregistered Commands (which
otherwise produce a relatively harmless "Unimplemented" error.)

Duplicate Names Append Underscore

Although every convenience has been designed to avoid naming conflicts
when importing Commands into the internal Register (Reg) duplicates are
still a possibility. Rather than override those previously added any
identical duplicate will simply have an underscore added to the name.
Since the processing of init functions is guaranteed to happen in
consistent fashion for any given composition this allows rare naming
conflicts to be resolved in the main init before calling Execute when
needed.

*/
package cmdbox
