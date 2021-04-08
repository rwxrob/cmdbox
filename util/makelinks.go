package util

// MakeSymLinks will destructively create one symbolic link within the
// target directory for each VisibleCommand sorted in lexographical
// order. If any error is encountered link creation is halted and the
// error is returned. Existing files/links are not considered an error
// and will be forcibly overwritten. Use caution. The directory passed
// as the only argument must exist and have compatible permissions. This
// function is only supported on Unix-flavored operating systems.
func MakeSymLinks(from, dir string) error {
	// TODO
	return nil
}

// MakeHardLinks will destructively create one hard link within the
// target directory for each VisibleCommand sorted in lexographical
// order. If any error is encountered link creation is halted and the
// error is returned. Existing files/links are not considered an error
// and will be forcibly overwritten. Use caution. The directory passed
// as the only argument must exist and have compatible permissions. This
// function is only supported on Unix-flavored operating systems.
func MakeHardLinks(from, dir string) error {
	// TODO
	return nil
}

// MakeCopies will destructively create one identical copy of the
// executable currently being executed within the target directory for
// each VisibleCommand sorted in lexographical order. If any error is
// encountered copying is halted and the error returned. Existing
// files/links are not considered an error and will be forcibly
// overwritten. Use caution. The directory passed as the only argument
// must exist and have compatible permissions. This function works on
// all operating systems but should only be used as a last resort when
// neither MakeSymLinks() nor MakeHardLinks() are available.
func MakeCopies(from, dir string) error {
	// TODO
	return nil
}
