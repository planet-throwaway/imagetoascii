package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertSmallImage(t *testing.T) {
	_, converted := Convert(smallImageData)
	assert.Equal(t, converted, smallImageAscii)
}

func TestConvertBigImage(t *testing.T) {
	// Want to be sure that we resize the image to 80 columns
	_, converted := Convert(bigImageData)
	assert.Equal(t, converted, bigImageAscii)
}

const smallImageData = `
iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAIAAACRXR/mAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6
JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAACXBIWXMAAAsTAAALEwEAmpwYAAACcmlU
WHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIg
eDp4bXB0az0iWE1QIENvcmUgNS40LjAiPgogICA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cu
dzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPgogICAgICA8cmRmOkRlc2NyaXB0aW9uIHJk
ZjphYm91dD0iIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYv
MS4wLyIKICAgICAgICAgICAgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIj4K
ICAgICAgICAgPHRpZmY6WVJlc29sdXRpb24+NzI8L3RpZmY6WVJlc29sdXRpb24+CiAgICAgICAgIDx0
aWZmOkNvbXByZXNzaW9uPjU8L3RpZmY6Q29tcHJlc3Npb24+CiAgICAgICAgIDx0aWZmOlhSZXNvbHV0
aW9uPjcyPC90aWZmOlhSZXNvbHV0aW9uPgogICAgICAgICA8eG1wOkNyZWF0b3JUb29sPkZseWluZyBN
ZWF0IEFjb3JuIDUuNi4zPC94bXA6Q3JlYXRvclRvb2w+CiAgICAgICAgIDx4bXA6TW9kaWZ5RGF0ZT4y
MDE3LTA1LTE5VDIyOjI0OjE4PC94bXA6TW9kaWZ5RGF0ZT4KICAgICAgPC9yZGY6RGVzY3JpcHRpb24+
CiAgIDwvcmRmOlJERj4KPC94OnhtcG1ldGE+CkGjZOIAAAE5SURBVFjD7dlLDoQgDAbgYTkHmvsfwAO5
xJ0htbR/H6BmYKk8PhtNsZRa6+d5rSzWRFYpBelmXcXDAikRn4EV0VhlOgvU7Bu98v35ZQpLNV01bCPE
EEs2gSAhctLSvXtXk8kB4rqrszdGmEwxm8rC3zOGRUyJIFxGWYjpnNQtVl+yGSy2vxwwiXXOQuZVWWD/
ViaxeqFqx+/bEBaRdVm9gCeyhICZWchHij+GjQXmXYSFf0A8qxeqR7NS8sz9LFYWYo3INoh+sRbrJpZD
FtnqzGaB1neyIrK/Z5lkY1nugJENFnsljeUrMbT7vjSWTyYUZ0axTFNHBiqsRJlpiM4SqkhZGVP+sfbU
tyZUlEJFykhijrKyZKZSamalGfnjDVUDIziwKaVk3ynG6BODhDMfmeh87HVC9n7WAcQuuaxCZRjoAAAA
AElFTkSuQmCC`

const smallImageAscii = `..................................................
.....................nnnnnnnnnnnnn................
..............nnn$$$$$$$$$$$$$$$$$$$$n............
..........n$$$$$$$$$$CCC______CCCC$$$$$...........
........n$$$$$CCCCC_______________CC$$$$$n........
.......$$$$$$_______________________CC$$$$$.......
.....n$$$$C______CCC___________________$$$$$n.....
.....$$$$C__CCC__$$$________CCC__$$$____C$$$$.....
....$$$$C___$$$$_$$$________$$$__$$$_____C$$$$....
....$$$C____C$$$$$$$________$$$$C$$$______C$$$....
...$$$$______CCCCCCC________C$$$$$$$_______$$$....
...$$$C_______________________CCCCCC_______$$$....
...$$$_____________________________________$$$....
...$$$_____________________________________$$$....
...$$$$________________$$$____CCC__________$$$....
...n$$$C_______________$$$____$$$__________$$$....
....$$$$C______________$$$____$$$__________$$$....
.....$$$$C_____________$$$$_C$$$$__________$$$....
......$$$$$CC__________$$$$$$$$$C_________C$$$....
.......n$$$$$$CC________CCCCCCC_________C$$$$$....
.........nn$$$$$$CCC_________________CC$$$$$$.....
............nn$$$$$$$CCCC____CCCCCC$$$$$$$nn......
................nn$$$$$$$$$$$$$$$$$$$$nn..........
......................nnnnnnnnnn..................
..................................................
`

const bigImageData = `
iVBORw0KGgoAAAANSUhEUgAAAgAAAAIACAAAAADRE4smAAAAAXNSR0IArs4c6QAAAAlwSFlzAAALEwAA
CxMBAJqcGAAAAjBpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRv
YmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRm
PSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpE
ZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRv
YmUuY29tL3hhcC8xLjAvIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29t
L3RpZmYvMS4wLyI+CiAgICAgICAgIDx4bXA6Q3JlYXRvclRvb2w+QWNvcm4gdmVyc2lvbiAzLjUuMTwv
eG1wOkNyZWF0b3JUb29sPgogICAgICAgICA8dGlmZjpDb21wcmVzc2lvbj41PC90aWZmOkNvbXByZXNz
aW9uPgogICAgICAgICA8dGlmZjpZUmVzb2x1dGlvbj43MjwvdGlmZjpZUmVzb2x1dGlvbj4KICAgICAg
ICAgPHRpZmY6WFJlc29sdXRpb24+NzI8L3RpZmY6WFJlc29sdXRpb24+CiAgICAgIDwvcmRmOkRlc2Ny
aXB0aW9uPgogICA8L3JkZjpSREY+CjwveDp4bXBtZXRhPgoHbBTiAAAGV0lEQVR42u3dW27bSBCGUSrI
Cmr/a+wtJI9JBDIsUWz27Zy3ASYXaz7U/JJl+1U2VvbDQyAABIAAEAACQAAIAAEgAASAABAAAkAACAAB
IAAEgAAQAAJAAAgAASAABIAAEAACYBA/PQQrCBcAASAABIARuPTqcwEQAAJAABiBK88/FwABIAABIAA8
C2Dqse8CIAAEgAAwAhedf8UFQAAIAAFgBNp/LgACEICHQAAYgSy4/1wABCAAD4EAMAIZV3EBEAACQAAY
gROKSvvPBUAAAvAQCAABIAAEgAAQAAJgGfmXgt9fkCwevPqqvgrsAiAAAXgIBIAReGGKhE3oAiAABIAA
mHoEXhmKRuHND+j3j+zufyIXwP8CEAACwAisv2Esw6/m382/ZbgACEAACAAj8MF5YwleVq4+5i4AAkAA
CAABcMuzgJKYmCez1POB/GQvLgACQAAIgC5G4PE0SS9DbyS98qi5AAgAASAAmo/Ak2X44bBZ9TXCuLC0
XQAEgAAQAL2NwLf5cvlTxpMvwcg+gC4AAkAACIDBRuDukLn82eKHZlEf8++ODzRcAASAABAA7Ubgybq5
8injuV8jvOOjCxcAASAABIAA6OlZwPHm/exF4iGfCjz/9d8uAAJAAAiALkfg7g7KzEHfdtQFQAAIAAEw
xQjcm4MnmzC2CabgrX/9cAEQAAJAAAw7AndX0skU9KKgC4AAEAACoNEIrPqtHi9/x6E+hAuAABAAAmC+
ERh7/2wJugAIAAEgAATAxM8C2hvyTQHFBUAACAABIAAEgAAQAALgEWO+Enj8poBVfwD5dviIuAAIAAEg
AKYZgelFtMwSDBcAASAABMAiIzDzs2bu/T6S83xBUnEBEAACQAACwLOAyV19VXig1Z/5tgnFBUAACAAB
MNcIzEyg47eLTvMKb3EBEAACQAAsNQLfJlB6DuICCAABIACMwMnmYFT/E/LCBUAACAABYASOuATLTb86
7votXQAEgAAQAEbgyfqKz/71+n8hFwABIAAEgADwLGC16e0CIAAEgACYegSu9S3bXQAEgAAQAAJAAAgA
ASAA+nXbK4HHb7/0cqALgAAQAAJg2hF4orsvi8YFQAAIQAAYgbdJ/+wWS9AFQAAIAAEgACZ5FpB+KoAL
gAAQAAJgshH4Zwluh3NwmdeDj39ouQuAABAAAmDEEejT+vn91/ED5gL4XwACQAAYgZWUvVlkPboACAAB
IAAaeqXXWGRWXiUDf3G5TwfjfwEIAAEw6wisugejzex0ARAAAkAACIB55d8PUDLPBbxT1AVAAAgAATCC
15XFFpm5eCffbN4FQAAIAAHQfgSeLEE7zQVAAAgAATDtCOx9CXr10AVAAAgAAfD8CGw5wXw1kQuAABAA
AuD/vvxGkX5auAuAABAAAkAArPYs4LOnAp4LuAAIAAEgAPrwqrDLHv/q8dSfbYC6AAgAASAAqo7ARkuw
5fZ0ARAAAkAAjKXaTw/PfI7Y/nMBEAACQABMNgL7mGTmnwuAABAAAkAA9PAswBMAFwABIAAEgBH4nfCf
0AVAAAgAASAABIAAEAACIGHKTwf7TLALgAAQAAJAAAgAASAABMCbgV8J9HZAFwABIAAEgAAQAAJAAAgA
AZAz25tCvR/UBUAACAABIAAEgAAQAAJgT4NXAmPbvGLnAiAABIAAEAACQAAIAAHwrLqvBJZt82XcLgAC
QAAIAAEgAASAABAAAqC1Vl8e7p2hLgACQAAIAAEgAASAABAAT3rglcD27wyNf/8uuAAIAAEgACNw7g8v
Tv650jKMmr+5C4AAEAACYPARWPWNgfH4XyMe/xhdAASAABAAAmC4ZwFDzOT7nm24AAgAASAAlhuBpePB
9NkCDRcAASAABIAReGFuxeGyevz1wA//wKj6u7sACAABIABmHYGZkTXDV3AXFwABIAAEgBG4t5AG+qRq
ZD8mFwABIAAEgBE4wxKM9EfiAiAABIAAEAADebUcsfHVtM786r/+nV/btm2vxJ+Qfm5SXAAEgAAQAEbg
rSMwM7IavZBcXAAEgAAQAEZgwyX4/AgsLgACQAAIACPwmSXY/uXA4gIgAASAADACm0zB8tWvNv9cAASA
ABAAQ47AdTaZC4AAEAACQADU19W3ix/te4i6AAgAASAAjEBL0AVAAAgAATCM3xLTqux9whCIAAAAAElF
TkSuQmCC
`

const bigImageAscii = `""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""l++JJJJ
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""!)))n#$$$$$$
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""">C%$$$$$$#\lll
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""-a$$h/]]]];""""
"""""I/Xv~~;""""""""""""""""""""""""""""""""""""""""""""""""""""_b$$*]""""""""""
"""""|$$$$$adQ|||""""""""""""""""""""""""""""""""""""""""""""""fo$$M_"""""""xddf
"""""|$$8O888$$$$*n?ll""""""""""""""""""""""""""""""""""""""";f$$$Ol"""""""b$$$$
"""""|$$*"""")vW$$$$$BmuI"""""""""""""""""""""""""""""""""")r#$$Wv!""""""""o$$$$
"""""|$$*""""""Il~nn*8$$888q?""""""""""""""""""""""""""""Ih$$B8fI""""""""]h$$$$$
"""""|$$*"""""""""""";)zmp$$$Q}""""""""""""""""""""""""""O$$W1""""""""""iC$$$$$$
"""""|$$*""Irt;"""""""""",?w@$$z""""""""""""""""""""""""?o$$O"""""""""""|$$$$$$$
"""""|$$*""w$$v"""""""""""""w$$o?"""""""""""""""""""""""r$$W}"""""""""""|$$$$$$$
"""""|$$*""w$$v""""""""""""">88$Wk|I"""""""""",lllllllllu$$o"""""""""""xW$$$$$B8
"""""|$$*""w$$v"""""""""""""""<#$$$#f]"">]]]vLO$$$$$$$$$$$$o"""""""""""tbbbbOt+"
"""""|$$*""w$$v""""""""""""""""l~X%$$$**8$$$$$88wXXXXXXC8@$b""""""""""""""""""""
"""""|$$*""w$$v""""""""""""""""""">coooooo||||I""""""""""]|"""""""""""""""""""""
"""""|$$*""w$$M}""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"""""|$$*"""0$$$OY<"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"""""|$$*""""|L#$$p<"""""""""""""""""""""""""""""""""""t]"""""""""""""""""""""""
"""""}W$$p>""""}tbC<""""";]>""""""""""""""""""""""""""0$W}""""""""""""""""""""""
""""""tM$$&oj""""""""""""#$&[""""""""""""""""""""""""""+i"""""""""""""""""""""""
"""""""InW$$$>"""""""""""~n{:"""":~i""""""""""""""""""""""""""""""""""""""""""""
"""""""""?h$$8bzI""""""""""""""""m$8|"""""""""""""""""""""""""""""""""""""""""""
"""""""""";\nM$$8u!""""""""""""lj8$$r"""""""""""""""""""""""""""""""""""""""""""
"""""""""""""lL$$8\"""""""""""<#$$#vi"""""""""""""""""""""""""""""""""""""""""""
""""""""""""""v$$o"""""""""""!J$$#I"""""""""""""""""""""""""""""""""""""""""""""
""""""""""""""v$$o"""""""""""+$$$/""""""""""""""""""""""""""""""""""""""""""""""
""""""""""""""v$$o"""""""""""8$$%<""""""""""""""""""""""""""""""""""""""""""""""
""""""""""""""v$$M_""""""""""$$$|"""""""""""""""""""""""""""""""""""""""""""""""
""""""""""""""]a$$U""""""""""$$$|"""""""""""""""""""""""""""""""""""""""""""""""
"""""""""""""""O$$*]"""""""""$$$|"""""""""""""""""""""""""""""""""""""""""""""""
"""""""""""""""iM$$t,""""""""$$$|,?J*znn{:""""""""":l\nl""""""""""""""""""""""""
""""""""""""""""O$$$Z""""""""$$$|~$$$$$$M]"""""l]]Yq$$$o""""""""""""""""""""""""
"""""""""""""""""1B$$k)""""""$$$|~$$$$$Jl""""I_W$$$880?<""""""""""""""""""""""""
""""""""""""""""""<b$$pi"""""$$$v~*$#B8t"";]XW$$$ht"""""""""""""""""""""""""""""
"""""""""""""""""""t8$$&v""""$$$$*M$W@Booo#$$$%x+:""""""""""""""""""""""""""""""
"""""""""""""""""""""u$$$Y+""ltYbbbbbbbbbbbbz<""""""""""""""""""""""""""""""""""
""""""""""""""""""""";/o$$w"""""""""""""""""""""""""""""""""""""""""""""""""""""
`
