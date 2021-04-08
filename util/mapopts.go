package util

import "strings"

// MapOpts takes an array of getopt like argument strings, examines them for
// options, and returns the options as a map and the remaining args as an
// array. To be clear:
//
//     * single dash options cannot have a value (-g)
//     * single dash options can be combined (-tgz)
//     * double dash options can have value (--template template.html)
//     * double dash options can use equals (--template=template.html)
//     * options with no value will have empty string
//     * values with quotes will contain quotes (--template='template.html')
//
// Note that there are many other getopt libraries. This is a minimal
// implementation of the most common getopt library. The most notable
// distinction is that single dash options cannot ever have values.
//
// IMPORTANT: While MapOpts is useful and encouraged when called from within
// a Command.Method implementation requiring dashed options, very often --- in
// our emerging conversational user interface world --- it is discouraged to
// use any dashed options at all. Instead create natural language options and
// arguments that can be spoken into an interface rather than typed and
// detected by context instead of looking for the dash, which is extremely
// unfriendly to voice interfaces of any kind. Nevertheless the choice is
// up to you.
func MapOpts(stuff []string) (opts map[string]string, args []string) {
	opts = make(map[string]string)
	args = []string{}
	for n := 0; n < len(stuff); n++ {
		thing := stuff[n]
		ln := len(thing)
		switch {
		case ln > 2 && thing[0:2] == "--":
			x := strings.Split(thing[2:], "=")
			if len(x) == 1 {
				if len(stuff) > n && len(stuff[n+1]) > 0 && stuff[n+1][0] != '-' {
					opts[x[0]] = stuff[n+1]
					n++
					continue
				}
				opts[x[0]] = ""
			}
			opts[x[0]] = x[1]
		case ln > 1 && thing[0] == '-' && thing[1] != '-':
			for _, o := range thing[1:] {
				opts[string(o)] = ""
			}
		default:
			args = append(args, thing)
		}
	}
	return
}
