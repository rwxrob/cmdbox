/*
Copyright 2021 Robert S. Muhlestein.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fmt

import (
	"testing"
)

func TestEmphasize(t *testing.T) {
	want := []string{
		italic + "Italic" + reset,
		bold + "Bold" + reset,
		bolditalic + "BoldItalic" + reset,
		underline + "BRACKETED" + reset,
	}
	args := []string{"*Italic*", "**Bold**", "***BoldItalic***", "<bracketed>"}
	for i, arg := range args {
		t.Logf("testing: %v\n", arg)
		got := Emphasize(arg)
		if got != want[i] {
			t.Errorf("\nwant: %q\ngot:  %q\n", want[i], got)
		}
	}
}

func TestEmph(t *testing.T) {
	text := `
    Something *easy* to write here that can be indented however you like
		and wrapped and have each line indented and with <code>:

        This will not be messed with.
        Nor this.

    So it's a lot like a **simple** version of Markdown that only supports
    what is likely going to be used in stuff similar to man pages.

    Let's try a hard  
    return.`

	want := "     Something " + italic + "easy" + reset + " to write here that can be indented however\n     you like and wrapped and have each line indented and with\n     " + underline + "CODE" + reset + ":\n     \n         This will not be messed with.\n         Nor this.\n     \n     So it's a lot like a " + bold + "simple" + reset + " version of Markdown that only\n     supports what is likely going to be used in stuff similar to\n     man pages.\n     \n     Let's try a hard\n     return."

	got := Emph(text, 5, 70)
	t.Log("\n" + got)
	if want != got {
		t.Errorf("\nwant:\n%q\ngot:\n%q\n", want, got)
	}
}

func TestPlain(t *testing.T) {
	text := `
    Something *easy* to write here that can be indented however you like
		and wrapped and have each line indented and with <code>:

        This will not be messed with.
        Nor this.

    So it's a lot like a **simple** version of Markdown that only supports
    what is likely going to be used in stuff similar to man pages.

    Let's try a hard  
    return.`

	want := "     Something *easy* to write here that can be indented however\n     you like and wrapped and have each line indented and with\n     <code>:\n     \n         This will not be messed with.\n         Nor this.\n     \n     So it's a lot like a **simple** version of Markdown that only\n     supports what is likely going to be used in stuff similar to\n     man pages.\n     \n     Let's try a hard\n     return."

	got := Plain(text, 5, 70)
	t.Log("\n" + got)
	if want != got {
		t.Errorf("\nwant:\n%q\ngot:\n%q\n", want, got)
	}
}
