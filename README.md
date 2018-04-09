# sim
sim is a programmers text editor, where the programmer is neosimsim, i.e. me.

## Name
The editor is heavily inspired by vi(m) and sam and so is the name.

And of course the matches my nick :)

Of course the editor is inspired by acme as well, but not the name.

## Warranty
You are welcome to use sim and give feedback but beware that I will still analyse
what fascinates me about acme, sam and vim and might change behaviours of sim as
I get more enlightened.

## Goals
- Write an editor supporting the full sam command language.
- acme-like virtual file server, supporting acme compatible plugin.

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

