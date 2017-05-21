package main

import (
	"github.com/stretchr/testify/assert"
	"image/color"
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

func TestConvertBadImage(t *testing.T) {
	// Want to be sure that we return an error with a bogus image
	err, converted := Convert("baddata")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "image: unknown format")
	assert.Empty(t, converted)
}

func TestPixelColorToAscii(t *testing.T) {
	grayScaleBlack := color.Gray{0}
	asciiBlack := pixelColorToAscii(grayScaleBlack)
	assert.Equal(t, "$", asciiBlack)

	grayScaleWhite := color.Gray{255}
	asciiWhite := pixelColorToAscii(grayScaleWhite)
	assert.Equal(t, " ", asciiWhite)
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

const smallImageAscii = `                                                  
                     nnnnnnnnnnnnn                
              nnn$$$$$$$$$$$$$$$$$$$$n            
          n$$$$$$$$$$JJJ______JJJJ$$$$$           
        n$$$$$JJJJJ_______________JJ$$$$$n        
       $$$$$$_______________________JJ$$$$$       
     n$$$$J______JJJ___________________$$$$$n     
     $$$$J__JJJ__$$$________JJJ__$$$____J$$$$     
    $$$$J___$$$$_$$$________$$$__$$$_____J$$$$    
    $$$J____J$$$$$$$________$$$$J$$$______J$$$    
   $$$$______JJJJJJJ________J$$$$$$$_______$$$    
   $$$J_______________________JJJJJJ_______$$$    
   $$$_____________________________________$$$    
   $$$_____________________________________$$$    
   $$$$________________$$$____JJJ__________$$$    
   n$$$J_______________$$$____$$$__________$$$    
    $$$$J______________$$$____$$$__________$$$    
     $$$$J_____________$$$$_J$$$$__________$$$    
      $$$$$JJ__________$$$$$$$$$J_________J$$$    
       n$$$$$$JJ________JJJJJJJ_________J$$$$$    
         nn$$$$$$JJJ_________________JJ$$$$$$     
            nn$$$$$$$JJJJ____JJJJJJ$$$$$$$nn      
                nn$$$$$$$$$$$$$$$$$$$$nn          
                      nnnnnnnnnn                  
                                                  
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

const bigImageAscii = `^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^I~~UUUU
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^l)))n#$$$$$$
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^iC%$$$$$$#|lll
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^_a$$h\????:^^^^
^^^^^;/Xu~~:^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^+b$$*?^^^^^^^^^^
^^^^^($$$$$apQ(((^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^to$$M+^^^^^^^xppf
^^^^^($$8O888$$$$*x-lI^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^:t$$$OI^^^^^^^b$$$$
^^^^^($$*^^^^)vW$$$$$Bmn;^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^)r#$$Wvl^^^^^^^^o$$$$
^^^^^($$*^^^^^^;l<xx*8$$888q-^^^^^^^^^^^^^^^^^^^^^^^^^^^^;h$$B8f;^^^^^^^^?k$$$$$
^^^^^($$*^^^^^^^^^^^^:)zmp$$$L[^^^^^^^^^^^^^^^^^^^^^^^^^^O$$W1^^^^^^^^^^!C$$$$$$
^^^^^($$*^^Ir/:^^^^^^^^^^,-w@$$z^^^^^^^^^^^^^^^^^^^^^^^^-o$$O^^^^^^^^^^^($$$$$$$
^^^^^($$*^^w$$v^^^^^^^^^^^^^w$$o-^^^^^^^^^^^^^^^^^^^^^^^j$$W[^^^^^^^^^^^($$$$$$$
^^^^^($$*^^w$$v^^^^^^^^^^^^^i88$Wk(;^^^^^^^^^^"llllllllln$$o^^^^^^^^^^^xW$$$$$B8
^^^^^($$*^^w$$v^^^^^^^^^^^^^^^>#$$$#t?^^i???uLO$$$$$$$$$$$$o^^^^^^^^^^^tbbbb0t~^
^^^^^($$*^^w$$v^^^^^^^^^^^^^^^^I~X%$$$**8$$$$$88mXXXXXXC8@$b^^^^^^^^^^^^^^^^^^^^
^^^^^($$*^^w$$v^^^^^^^^^^^^^^^^^^^ivoooooo((((;^^^^^^^^^^?(^^^^^^^^^^^^^^^^^^^^^
^^^^^($$*^^w$$M[^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^($$*^^^0$$$OY>^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^($$*^^^^|L#$$p>^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^t?^^^^^^^^^^^^^^^^^^^^^^^
^^^^^[W$$pi^^^^[tbC>^^^^^:?i^^^^^^^^^^^^^^^^^^^^^^^^^^0$W[^^^^^^^^^^^^^^^^^^^^^^
^^^^^^/M$$&oj^^^^^^^^^^^^*$&]^^^^^^^^^^^^^^^^^^^^^^^^^^~!^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^;xW$$$i^^^^^^^^^^^<x},^^^^,~!^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^-h$$8bc;^^^^^^^^^^^^^^^^m$8(^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^:|xM$$8ul^^^^^^^^^^^^If8$$j^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^IL$$8|^^^^^^^^^^^>#$$#u!^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^v$$o^^^^^^^^^^^lU$$#;^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^v$$o^^^^^^^^^^^~$$$\^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^v$$o^^^^^^^^^^^8$$%>^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^v$$M+^^^^^^^^^^$$$(^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^?a$$U^^^^^^^^^^$$$(^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^O$$*?^^^^^^^^^$$$(^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^!M$$t,^^^^^^^^$$$("-J*zxx},^^^^^^^^^,l|xI^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^0$$$Z^^^^^^^^$$$(<$$$$$$M?^^^^^l??Yq$$$o^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^^{B$$k)^^^^^^$$$(<$$$$$Ul^^^^;_W$$$880?>^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^^^>b$$p!^^^^^$$$u<*$*B8t^^:?XW$$$ht^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^^^^t8$$Wu^^^^$$$$*M$W@%ooo*$$$%r~,^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^^^^^^n$$$Y~^^ItXbbbbbbbbbbbbc>^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
^^^^^^^^^^^^^^^^^^^^^:\o$$w^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`
