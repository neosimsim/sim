# sim
sim is a programmers text editor, where the programmer is me.

Of course you are welcome to use sim and give feedback but without any warranty.

## Goals (or "Why another editor?")
In the next section a brief description on editors that inspired my way of
thinking about text editing and how is given. Nevertheless none of these editor
match my style in total.

I'm a keyboard user. I know that keyboard vs. mouse is controversial but
not to me.  I, at least feel to be more productive using the keyboard
almost exclusively. I feel distracted by looking for the cursor. I
want to write the name of a window or program so the window manager/the
machine can do the search. Then I want to tell the machine what to do,
also without looking for the cursor/button.

A list of must have features to me:

- sam command language
- integrating plugin support using tmux
- present text in some screen (curses, X whatever)
- run on alpine-linux, and OpenBSD

Features which shall not be included are

- mouse support - if you want a mouse just use acme
- syntax highlighting - if you need syntax highlighting use or write some
	screen to highlight text. This is not an editor task.

## Inspirations
### The name
The editor is heavily inspired by ed, sam, acme and vi(m) and so is the name.
I realized that s+im matched my nickname and so a name was found.

What follows are thought on the editors with feature a appreciate to clear my mind
so a see which features my personal editor must have **and** most of all shall not.

### ed
If you starts using ed, especially if you come the IDE (IDE like "integrated" not "integrating", s. [Pike][pike-acme])
kind of world, you way complain that you don't see what you get. If you want to see text of the file you have
to tell ed to print it.

As is started to play around with acme I struggled myself. Then one day I realized: I don't don't care about the text
present in the file when I write code. Every time I have an editor open showing the content while writing I catch myself
searching the file for the place to edit by scrolling up or done. Searching is the one true thing the machine is superior.
Now I just want to tell ed to change the lines matching a criteria, and I don't care where the lines are.

Still I have to confess that seeing the text is comfortable. But ed helped to change the way I think about editing.

### sam
sam comes with a very powerful command language. It is quite similar to ed's but using [structural regular expressions][SRE].
The core difference to ed is that for watch match regular expression sam select the matched text instead of whole line, which is
what ed does.

Combined with the way I think about editor as taught by ed, this makes text editing real fun.

My complain about sam is that GUI sucks.

### acme
acme has it all. The sam command language and integrating "plugin" support. acme has two issue I struggle with.

First it requires Plan9(Ports) and X, which are both not available on all platforms, e.g. alpine-linux (musl).

Second acme is based on mouse usage, which might be good but not to me. Although I admit the way acme or plan9/rio
uses the mouse is quite brilliant.

### vi(m)
The way you can move the cursor around in vi's normal mode is great. It is even used in several other tools like
sh and tmux.

Once upon a time I thought vim would be the only truth. Today I would say most of the features suck, most of verynomagic.

### tmux
How does tmux fits in?

On my thought on acme I mentioned the integrating "plugin" support. The same can be achieved by tmux, which can
be used by keyboard only.

A tmux pane runs a sim viewer bound to a file in the sim file server, watching it's content.

## Roadmap

### v0.1
- Implement the sam command language (except `u`)
- Usage only like `sam -d`

### v0.2
- Define the virtual file server include the concept of a file for *lokal file*

### v0.3
- separate tui program to **watch** the content of a file in the file server, including dot.

	This programm can be used to view the file currently edited in a separate tmux pane

### v1.0
Smooth everything.

## Notes and thoughts
### Features
- Enter on selection calls plumb on selection
- close and attach UI
- UTF8 only
- no plugins
- no syntax highlighting

	To be save against (UI) crashes each file should be kept "open" even if the UI crashes.
	Similar to close and detach on tmux. If the terminal or ssh connection "crashed" tmux
	still runs an can be reattached.

### File server
- <n>

A UI, no matter GUI or TUI just should just watch and present the content of <n>/body

### Keyboard or mouse
I know there are a lot of discussion which is more productive.
I actually don't care. I just have more fun using the keyboard primary.
I don't want to convince anybody. If you prefer the mouse oder the keyboard
just use acme.

