package main

import (
	"strings"
)

type Range struct {
	Start int
	End   int
}

// A File is an in memory copy of an external file, e.g. stored on disk.
type File struct {
	Buffer   string
	Range    Range
	FileName string
}

// Text commands

// Insert the text into the file after the range.
// Set dot.
func (f *File) Append(text string) {
	f.Buffer = strings.Join([]string{f.Buffer[:f.Range.End], text, f.Buffer[f.Range.End:]}, "")
	f.Range.Start = f.Range.End
	f.Range.End = f.Range.Start + len(text)
}

// Same as a, but c replaces the text, while i inserts
// before the range.
func (f *File) Change(text string) {
	f.Buffer = strings.Join([]string{f.Buffer[:f.Range.Start], text, f.Buffer[f.Range.End:]}, "")
	f.Range.End = f.Range.Start + len(text)
}

// Delete the text in the range.
// Set dot.
func d() {
}

// Substitute text for the first match to the regular
// expression in the range.  Set dot to the modified
// range.  In text the character & stands for the string
// that matched the expression. Backslash behaves as usual
// unless followed by a digit: \d stands for the string
// that matched the subexpression begun by the d-th left
// parenthesis.  If s is followed immediately by a number
// n, as in s2/x/y/, the n-th match in the range is sub-
// stituted.  If the command is followed by a g, as in
// s/x/y/g, all matches in the range are substituted.
func s() {
}

// Move the range to after a1. Set dot.
func m() {
}

// Copy the range to after a1. Set dot.
func t() {
}

// Display commands

// Print the text in the range.  Set dot.
func (f *File) PrintDot() string {
	return f.Buffer[f.Range.Start:f.Range.End]
}

// =    Print the line address and character address of the
//    range.
func printStatus() {
}

//  =#   Print just the character address of the range.
func printRange() {
}

// File commands

// Set the current file to the first file named in the
// list that sam also has in its menu.  The list may be
// expressed <Plan 9 command in which case the file names
// are taken as words (in the shell sense) generated by
// the Plan 9 command.
func b() {
}

// Same as b, except that file names not in the menu are
// entered there, and all file names in the list are exam-
// ined.
func B() {
}

// Print a menu of files.  The format is:
//                ' or blank indicating the file is modified or clean,
//                - or +     indicating the file is unread or has been
//                           read (in the terminal, * means more than one
//                           window is open),
//                . or blank indicating the current file,
//                a blank,
//                and the file name.
func n() {
}

// Delete the named files from the menu.  If no files are
// named, the current file is deleted.  It is an error to
// D a modified file, but a subsequent D will delete such
// a file.
func D() {
}

// I/O Commands

// Replace the file by the contents of the named external
// file.  Set dot to the beginning of the file.
func e() {
}

// Replace the text in the range by the contents of the
// named external file.  Set dot.
func r() {
}

// Write the range (default 0,$) to the named external
// file.
func w() {
}

// * f filename
//    Set the file name and print the resulting menu entry.
// If the file name is absent from any of these, the current
// file name is used.  e always sets the file name; r and w do
// so if the file has no name.
func f() {
}

// < Plan 9-command
//    Replace the range by the standard output of the Plan 9
//    command.
func pipeIn() {
}

// > Plan 9-command
//    Send the range to the standard input of the Plan 9 com-
//    mand.
func pipeOut() {
}

// | Plan 9-command
//    Send the range to the standard input, and replace it by
//    the standard output, of the Plan 9 command.
func pipe() {
}

// * ! Plan 9-command
//    Run the Plan 9 command.
func exec() {
}

// * cd directory
//    Change working directory.  If no directory is speci-
//    fied, $home is used.
func cd() {
}

// In any of <, >, | or !, if the Plan 9 command is omitted the
// last Plan 9 command (of any type) is substituted.  If sam is
// downloaded (using the mouse and raster display, i.e. not
// using option -d), ! sets standard input to /dev/null, and
// otherwise unassigned output (stdout for ! and >, stderr for
// all) is placed in /tmp/sam.err and the first few lines are
// printed.

// Loops and Conditionals

// x/regexp/ command
//    For each match of the regular expression in the range,
//    run the command with dot set to the match.  Set dot to
//    the last match.  If the regular expression and its
//    slashes are omitted, `/.*\n/' is assumed.  Null string
//    matches potentially occur before every character of the
//    range and at the end of the range.
// y/regexp/ command
//    Like x, but run the command for each substring that
//    lies before, between, or after the matches that would
//    be generated by x.  There is no default regular expres-
//    sion.  Null substrings potentially occur before every
//    character in the range.
// * X/regexp/ command
//    For each file whose menu entry matches the regular
//    expression, make that the current file and run the com-
//    mand.  If the expression is omitted, the command is run
//    in every file.
// * Y/regexp/ command
//    Same as X, but for files that do not match the regular
//    expression, and the expression is required.
// g/regexp/ command
// v/regexp/ command
//    If the range contains (g) or does not contain (v) a
//    match for the expression, set dot to the range and run
//    the command.
// These may be nested arbitrarily deeply, but only one
// instance of either X or Y may appear in a single command.
// An empty command in an x or y defaults to p; an empty com-
// mand in X or Y defaults to f.  g and v do not have defaults.

// Miscellany

// k        Set the current file's mark to the range.  Does not
//    set dot.
// * q      Quit.  It is an error to quit with modified files,
//    but a second q will succeed.
// * u n    Undo the last n (default 1) top-level commands that
//    changed the contents or name of the current file,
//    and any other file whose most recent change was
//    simultaneous with the current file's change.  Suc-
//    cessive u's move further back in time.  The only
//    commands for which u is ineffective are cd, u, q, w
//    and D.  If n is negative, u `redoes,' undoing the
//    undo, going forwards in time again.
// (empty)  If the range is explicit, set dot to the range.  If
//    sam is downloaded, the resulting dot is selected on
//    the screen; otherwise it is printed.  If no address
//    is specified (the command is a newline) dot is
//    extended in either direction to line boundaries and
//    printed.  If dot is thereby unchanged, it is set to
//    .+1 and printed.
