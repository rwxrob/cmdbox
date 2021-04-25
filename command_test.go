package cmdbox

import (
	"os"
	"path/filepath"
	"testing"
)

func mockGit() (string, error) {
	dir, err := os.MkdirTemp("", "tmp-git")
	if err != nil {
		return "", err
	}
	err = os.Chdir(dir)
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(filepath.Join(dir, ".git", "refs", "tags"), 0700)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func mockTag(dir, name string) (*os.File, error) {
	filename := filepath.Join(dir, ".git", "refs", "tags", name)
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func cleanTags(tags []*os.File) {
	for _, tag := range tags {
		os.Remove(tag.Name())
	}
}

func TestVersion(t *testing.T) {
	dir, err := mockGit()
	if err != nil {
		t.Fatal("Error setting up test git")
	}
	defer os.RemoveAll(dir)

	testCases := []struct {
		Tags     []string
		Expected string
	}{
		{
			Tags:     []string{},
			Expected: "v0.0.1",
		},
		{
			Tags:     []string{"v0.0.1", "2.0.1", "v1.3.1"},
			Expected: "v1.3.1",
		},
		{
			Tags:     []string{"v0.0.1", "2.0.1", "v1.3.1", "HelloWorld", "v1.3.2"},
			Expected: "v1.3.2",
		},
		{
			Tags:     []string{"v1.0.0", "v0.0.1", "v0.4.2"},
			Expected: "v1.0.0",
		},
		{
			Tags:     []string{"1.0.0", "first-tag", "v1"},
			Expected: "v0.0.1",
		},
	}
	for _, test := range testCases {
		var tags []*os.File
		for _, tag := range test.Tags {
			f, err := mockTag(dir, tag)
			if err != nil {
				t.Fatalf("Error happened creating tag %s\n", tag)
			}
			tags = append(tags, f)
		}
		if v := version(); v != test.Expected {
			t.Errorf("Expected: %s\nActual: %s\n", test.Expected, v)
		}
		cleanTags(tags)
	}
}

func TestVersionNoTagsFolder(t *testing.T) {
	dir, err := os.MkdirTemp("", "no-tags")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if v := version(); v != "v0.0.1" {
		t.Errorf("Expected: v0.0.1\nActual: %s\n", v)
	}
}
