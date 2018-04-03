# plan9 mit der Tastatur

Welche Tools sollen ersetzt werden?

- rio/mux
- acme
- sam

## Was stört (mich) an den Tools
In Summe sind alle Tools auf Bedienung mit einer (3-Button) Maus
ausgelegt.

### rio
Ich finde die Idee spannend, dass Unterprogramme das gleiche Fenster benutzen.

### acme
Die Idee von acme ist gut, jedoch passt das UI-Konzept nicht zu mir.
eine Art Fusion von acme und tmux wäre ideal.

### sam
Ein guter Editor. Eigentlich macht sam seine eigene *Command language* aus.
Viel mehr liefert sam nicht, abgesehen von einem nicht zu mir passenden UI-Konzept.

## Warum reicht mein aktuelles Setup nicht?
### Setup
#### i3
#### vim
vim ist ein guter Texteditor. Er erweiter vi zu sehr. Eine eigene *Plugin*
Sprache sollte unnötig sein, vgl. `<`, `>` und `|` aus sam.

#### st
#### tmux
#### sh

## Mögliche Features
- *Hoch* und *Runter* markieren den Text eine Zeile oberhalb
	`Return` löst `send` aus.
	Das Verhalten ist dan ähnlich zu `history-backwards-search`
	Wobei hier unklar ist was *Zeile* heißt.
- *Links* und *Rechts* reduzieren(/erweitern) die Auswahl um ein *Wort*
- *History* persistieren
- `/` ist das cmd zum suchen. Ähnlich zu `[` bei test. Genauso `?`
- Eine Shell zeigt Text an. Es gibt *alten Text* und *neuen Text*.
	Der neue Text wird mit `Return` ausgeführt. Beide Texte sollten durch
	eine Textdatei repräsentiert werden. Mir ist noch unklar, ob diese Funktionen
	durch eine Shell oder eine Terminal abgedeckt werden.
	- Ich neige zu Terminal, allerdings sind dies Funktionen die evt. auch andere
		Programme interessieren könnten.
	- Es gibt folgende Elemente
		- text
		- input
		- selection/range
- Jump to `output point`

Richtig gemacht kann alles was die Maus kann auch mit der Tastatur erledigt werden.
Da das Terminal nur Text anzeigt, ist es egal wo der Text angezeigt wird, ob vterminal,
st, 9term etc.

Allerdings müssen die Eingaben auch von einem Programm bearbeitet werden.

**FRAGE**
Wenn alles nur Text ist und editierbar ist, ist das Terminal dann nicht ein Texteditor?

## Was macht eine Maus aus?
- Auswahl eines Bereiches
	- Ist keine Bereich ausgewählt, dass ist das geklickte *Wort*
	  der Bereich
- Aktion mit den gewähltem Bereich ausführen.
	- Die Anzahl der Aktionen ist begrenzt durch die, Anzahlt der Tasten der Maus.

Kann man auch mit der Tastatur.

## Meinung
Meine persönliche Meinung ist, dass eine Bedienung mit der Tastatur überlegen ist.
