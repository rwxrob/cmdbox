/* Package cmdboxdoc provides parsing and formatting for the CmdBoxDoc
markup language used to document all CmdBox modules. See cmdboxdoc.pegn
for the formal specification in PEGN (pegn.dev).

* Initial and trailing blank lines are removed.

* Initial indentation is removed - the number of spaces preceeding
  the first word of the first line are ignored in every line (including
  raw text blocks).

* Raw text ignored - any line beginning with at least four spaces
  (after initial indentation is removed) will be kept as it is
  exactly (code examples, etc.) but in practice should never really
  exceed 80 characters (including the spaces). This is different than
  Godoc which only requires a single space.

* Blocks are unwrapped - any non-blank (without three or less initial
  spaces) will be trimmed and joined to the preceding line
  recursively (unless hard break).

* Hard breaks kept - like Markdown any line that ends with two or
  more spaces will automatically force a line return.

* Identifiers such as parameter names, URL links, and anything else
  within angle brackets (ex: <url>) will trigger Identifier formatting
  in both text blocks and usage sections.

* Emph, Strong, and EmphStrong inline emphasis using one, two, or
  three stars respectively will be observed and cannot be intermixed or
  intraword.  Each opener must be preceded by a UNICODE space (or
  nothing) and followed by a non-space rune. Each closer must be
  preceded by a non-space rune and followed by a UNICODE space (or
  nothing).

*/
package cmdboxdoc
