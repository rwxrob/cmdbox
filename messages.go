package cmdbox

// Messages contains every human readable message used by the cmdbox
// package and can be used by CmdBox module and composite developers to
// support dynamic locale output at runtime. By convention, any human
// readable natural language messages should be stored here rather than
// embedded thus preserving the ability to dynamically adapt to the
// language and locale. Since, by design, any CmdBox package can change
// anything in the Messages map module, creators should prefix their
// Message keys with the Command.Name and a single space to prevent
// unwanted accidental conflict. See JSON and Load for more.
var Messages = map[string]string{
	"InvalidName":   "invalid name (lower case words only)",
	"Unimplemented": "unimplemented: %v",
}
