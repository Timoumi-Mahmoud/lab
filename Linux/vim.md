###

```vim
:set scrolloff=8
:set number  :set relativenumber
:set rnu

dt  <char>

:%s/dog/cat/g
V + >    : intend

:r fileToPast
:r! date

set -o vi

:set noic # ignore case
:set hls # highlight search after doing / (all result) also with :set ic for case
:nohls   
I can do /SIG.* to highlight every word that begin with SIG


':s/foo/baz<CR>' search and replace

dt (    => fmt.Println("vim-go")  ==> ("vim-go")

V + "by   // yank in b register -> :reg then I can do "bp

f +  <firstCharacter>
F +  <firstCharacter> reverse

t +  <firstCharacter>
T +  <firstCharacter> reverse

{ and } vertical movement

'[m and ]m' move my curosor to { and }
'%' go the end of ( or { or [

yyp and yyP

V highlight a line

:reg  => keep track of delete and yank

SHIFT + i => go to first no white space character in a line

CNTRL + arrow keys


:Ex replace current buffer with explorer
:Vex 
:Sex


/////
CNTRL + w + (s horizontal , v vertical) 
close with CNTRL + w + o

////  Marks
global
m + Capital letter
' + Capital letter
internal 
m + lower letter
' + lower letter

/// Jumplist
CNTRL + o --->forward i<---backward

:reg % current file , # previous file

SHIFT + D delete to the end of the file

CNTRL + u 
CNTRL + d ==> jump up and down by half page

```
