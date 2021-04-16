package cmdbox

import (
	"testing"

	"github.com/rwxrob/cmdbox/util"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		Name    string
		SubCmds []string
	}{
		{Name: "calculator", SubCmds: []string{"add", "sub", "mul", "div"}},
		{Name: "todo", SubCmds: []string{}},
	}
	for _, test := range testCases {
		c := New(test.Name, test.SubCmds...)
		if c.Name != test.Name {
			t.Fatalf("Expected Name: %s\nActual: %s\n", test.Name, c.Name)
		}
		if len(c.Subcommands()) != len(test.SubCmds) {
			t.Fatalf("Expected %d subcommands\nActual: %d\n", len(test.SubCmds), len(c.Subcommands()))
		}
		for i, sub := range test.SubCmds {
			if c.Subcommands()[i] != sub {
				t.Fatalf("Expected subcommand name: %s\nActual: %s\n", sub, c.Subcommands()[i])
			}
		}
	}
}

func TestNewCommandHasName(t *testing.T) {
	testCases := []struct {
		Name    string
		SubCmds []string
	}{
		{Name: "", SubCmds: []string{}},
		{Name: "", SubCmds: []string{"start", "stop"}},
	}
	for _, test := range testCases {
		if c := New(test.Name, test.SubCmds...); c != nil {
			t.Fatal("Command.Name is undefined")
		}
	}
}

func TestHidden(t *testing.T) {
	testCases := []struct {
		Name     string
		IsHidden bool
	}{
		{Name: "format", IsHidden: false},
		{Name: "_format", IsHidden: true},
		{Name: "_doc", IsHidden: true},
		{Name: "open", IsHidden: false},
	}
	for _, test := range testCases {
		c := New(test.Name)
		if c.Hidden() != test.IsHidden {
			t.Fatalf("Expected %s to be hidden\n", test.Name)
		}
	}
}

func Testvsubcommands(t *testing.T) {
	testCases := []struct {
		Name    string
		SubCmds []string
	}{
		{
			Name:    "calculator",
			SubCmds: []string{},
		},
		{
			Name:    "calculator",
			SubCmds: []string{"add", "sub", "mul", "div"},
		},
	}
	for _, test := range testCases {
		c := New(test.Name, test.SubCmds...)
		if len(c.vsubcommands()) != len(test.SubCmds) {
			t.Fatalf("Expected %d commands\nActual: %d\n", len(test.SubCmds), len(c.vsubcommands()))
		}
		for i, sub := range c.vsubcommands() {
			if test.SubCmds[i] != sub.Name {
				t.Fatalf("Expected: %s\nActual: %s\n", test.SubCmds[i], sub.Name)
			}
		}
	}
}

func TestSprintUsage(t *testing.T) {
	testCases := []struct {
		Name     string
		Usage    util.Stringer
		SubCmds  []string
		Expected string
	}{
		{
			Name:    "todo",
			Usage:   "",
			SubCmds: []string{},
			Expected: `**todo**
`,
		},
		{
			Name:     "todo",
			Usage:    nil,
			SubCmds:  []string{},
			Expected: `**todo**`,
		},
		{
			Name:    "calculator",
			Usage:   "",
			SubCmds: []string{"add", "sub", "mul", "div"},
			Expected: `**calculator**
**calculator** **add**
**calculator** **sub**
**calculator** **mul**
**calculator** **div**
`,
		},
		{
			Name:    "calculator",
			Usage:   nil,
			SubCmds: []string{"add", "sub", "mul", "div"},
			Expected: `**calculator** **add**
**calculator** **sub**
**calculator** **mul**
**calculator** **div**
`,
		},
		{
			Name:    "crud",
			Usage:   "[create|read|update|delete]",
			SubCmds: []string{"create", "read", "update", "delete"},
			Expected: `**crud** [create|read|update|delete]
**crud** **create**
**crud** **read**
**crud** **update**
**crud** **delete**
`,
		},
	}
	for _, test := range testCases {
		c := New(test.Name, test.SubCmds...)
		c.Usage = test.Usage
		if len(test.SubCmds) > 0 {
			for _, subcmd := range test.SubCmds {
				_ = New(subcmd)
			}
		}
		su := c.SprintUsage()
		if su != test.Expected {
			t.Fatalf("Expected: %s\nActual: %s\n", test.Expected, su)
		}
	}
}

func TestSprintCommandSummaries(t *testing.T) {
	testCases := []struct {
		Name      string
		SubCmds   []string
		Summaries []string
		Expected  string
	}{
		{
			Name:      "print",
			SubCmds:   []string{},
			Summaries: []string{},
			Expected:  "",
		},
		{
			Name:      "pomo",
			SubCmds:   []string{"start", "stop"},
			Summaries: []string{"Starts the pomodoro", "Stops the pomodoro"},
			Expected: `**start**      Starts the pomodoro
**stop**       Stops the pomodoro
`,
		},
	}
	for _, test := range testCases {
		c := New(test.Name, test.SubCmds...)
		if len(test.SubCmds) > 0 {
			for i, subcmd := range test.SubCmds {
				sc := New(subcmd)
				sc.Summary = test.Summaries[i]
			}
		}
		scs := c.SprintCommandSummaries()
		if scs != test.Expected {
			t.Fatalf("Expected: %s\nActual: %s\n", test.Expected, scs)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		Name     string
		Usage    util.Stringer
		Author   util.Stringer
		License  util.Stringer
		SubCmds  []string
		Expected string
	}{
		{
			Name:     "todo",
			Usage:    "",
			Author:   "",
			License:  "",
			SubCmds:  []string{},
			Expected: `{"Name":"todo"}`,
		},
		{
			Name:     "print",
			Usage:    "",
			Author:   "John Doe",
			License:  "",
			SubCmds:  []string{},
			Expected: `{"Author":"John Doe","Name":"print"}`,
		},
		{
			Name:     "pomo",
			Usage:    "[start|stop]",
			Author:   "Rob",
			License:  "",
			SubCmds:  []string{},
			Expected: `{"Author":"Rob","Name":"pomo","Usage":"[start|stop]"}`,
		},
		{
			Name:     "express",
			Usage:    "",
			Author:   "",
			License:  "MIT",
			SubCmds:  []string{},
			Expected: `{"License":"MIT","Name":"express"}`,
		},
		{
			Name:     "pomo",
			Usage:    "",
			Author:   "",
			License:  "",
			SubCmds:  []string{"start", "stop"},
			Expected: `{"Name":"pomo","Subcommands":["start","stop"]}`,
		},
	}
	for _, test := range testCases {
		c := New(test.Name, test.SubCmds...)
		c.Usage = test.Usage
		c.Author = test.Author
		c.License = test.License
		mc, err := c.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v\n", err)
		}
		if string(mc) != test.Expected {
			t.Fatalf("Expected: %s\nActual: %s\n", test.Expected, string(mc))
		}
	}
}

func TestVersionLine(t *testing.T) {
	testCases := []struct {
		Name      string
		Version   util.Stringer
		Copyright util.Stringer
		License   util.Stringer
		Expected  string
	}{
		{
			Name:      "print",
			Version:   nil,
			Copyright: nil,
			License:   nil,
			Expected:  "",
		},
		{
			Name:      "print",
			Version:   "",
			Copyright: "",
			License:   "",
			Expected:  "",
		},
		{
			Name:      "print",
			Version:   "v1.0.0",
			Copyright: nil,
			License:   "MIT",
			Expected:  "print v1.0.0 (MIT)",
		},
		{
			Name:      "print",
			Version:   "v1.0.0",
			Copyright: "",
			License:   "",
			Expected:  "print v1.0.0",
		},
		{
			Name:      "print",
			Version:   "v1.0.0",
			Copyright: "copyright",
			License:   "MIT",
			Expected:  "print v1.0.0 copyright (MIT)",
		},
	}
	for _, test := range testCases {
		c := New(test.Name)
		c.Version = test.Version
		c.Copyright = test.Copyright
		c.License = test.License
		vl := c.VersionLine()
		if vl != test.Expected {
			t.Fatalf("Expected: %s\nActual: %s\n", test.Expected, vl)
		}
	}
}

func TestHas(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected bool
	}{
		{Input: "pause", Expected: false},
		{Input: "add", Expected: true},
		{Input: "div", Expected: true},
		{Input: "reset", Expected: false},
		{Input: "sub", Expected: true},
		{Input: "new", Expected: false},
	}
	for _, test := range testCases {
		c := New("calculator", "add", "sub", "mul", "div")
		if has := c.Has(test.Input); has != test.Expected {
			t.Fatalf("Has should not return true for %s\n", test.Input)
		}
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		Input    []string
		Expected []string
	}{
		{Input: []string{}, Expected: []string{}},
		{Input: []string{"add"}, Expected: []string{"add"}},
		{Input: []string{"add|sub", "mul"}, Expected: []string{"add", "sub", "mul"}},
		{Input: []string{"add|sub|mul|div"}, Expected: []string{"add", "sub", "mul", "div"}},
	}
	for _, test := range testCases {
		c := New("test")
		c.Add(test.Input...)
		subs := c.Subcommands()
		if len(subs) != len(test.Expected) {
			t.Fatalf("Expected commands to have %d subcommands\nActual: %d\n", len(test.Expected), len(subs))
		}
		for i, sub := range subs {
			if sub != test.Expected[i] {
				t.Fatalf("Expected subcommand: %s\nActual: %s\n", test.Expected[i], sub)
			}
		}
	}
}

func TestSubcommands(t *testing.T) {
	testCases := []struct {
		SubCmds []string
	}{
		{SubCmds: []string{}},
		{SubCmds: []string{"add"}},
		{SubCmds: []string{"add", "sub"}},
		{SubCmds: []string{"add", "sub", "mul"}},
		{SubCmds: []string{"add", "sub", "mul", "div"}},
	}
	for _, test := range testCases {
		c := New("test", test.SubCmds...)
		subs := c.Subcommands()
		if len(subs) != len(test.SubCmds) {
			t.Fatalf("Expected commands to have %d subcommands\nActual: %d\n", len(test.SubCmds), len(subs))
		}
		for i, sub := range subs {
			if sub != test.SubCmds[i] {
				t.Fatalf("Expected subcommand: %s\nActual: %s\n", test.SubCmds[i], sub)
			}
		}
	}
}

func TestSuncommandUsage(t *testing.T) {
	testCases := []struct {
		SubCmds   []string
		SubUsages []string
	}{
		{
			SubCmds:   []string{},
			SubUsages: []string{},
		},
		{
			SubCmds:   []string{"start", "stop"},
			SubUsages: []string{"Starts the command", "Stops the command"},
		},
	}
	for _, test := range testCases {
		c := New("test", test.SubCmds...)
		if len(test.SubCmds) > 0 {
			for i, usage := range test.SubUsages {
				sc := New(test.SubCmds[i])
				sc.Usage = usage
			}
		}
		subs := c.SubcommandUsage()
		if len(subs) != len(test.SubUsages) {
			t.Fatalf("Expected commands to have %d subcommands\nActual: %d\n", len(test.SubUsages), len(subs))
		}
		for i, sub := range subs {
			if sub != test.SubUsages[i] {
				t.Fatalf("Expected subcommand: %s\nActual: %s\n", test.SubUsages[i], sub)
			}
		}
	}
}

func TestTitle(t *testing.T) {
	testCases := []struct {
		Name     string
		Summary  string
		Expected string
	}{
		{Name: "random", Summary: "", Expected: "random"},
		{Name: "random", Summary: "A short summary", Expected: "random - A short summary"},
	}
	for _, test := range testCases {
		c := New("random")
		c.Summary = test.Summary
		title := c.Title()
		if title != test.Expected {
			t.Fatalf("Expected: %s\nActual: %s\n", test.Expected, title)
		}
	}
}
