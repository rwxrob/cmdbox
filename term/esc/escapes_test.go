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

package esc

// NOTE: These tests must be run with -v and verified by human eyes on
// a fully capable terminal for best test validation.

import (
	"testing"
	"time"
)

func Test_ClearScreen(t *testing.T) {
	t.Log(ClearScreen)
	t.Log("Something")
	time.Sleep(2 * time.Second)
	t.Log(CS)
}

func Test_Bold(t *testing.T) {
	t.Log(Bold + "bold" + Reset)
	t.Log(B + "bold" + X)
}

func Test_Italic(t *testing.T) {
	t.Log(Italic + "italic" + Reset)
	t.Log(I + "italic" + X)
}

func Test_BoldItalic(t *testing.T) {
	t.Log(BoldItalic + "bold_italic" + Reset)
	t.Log(BI + "bold_italic" + X)
}
