package losegui

/*
#include "losegui.h"

HANDLE GetOutputHandle() {
    return GetStdHandle(STD_OUTPUT_HANDLE);
}

HANDLE GetInputHandle() {
    return GetStdHandle(STD_INPUT_HANDLE);
}

CHAR * chartoCHAR(char * chr) {
    return chr;
}
*/
import "C"
import "unsafe"

//TODO Add in proper foreground and background colour changing

//TODO Maybe make another COORD object that's an interface to turn the existing one into
//a C compatable one or something ... the current method is annoying
type COORD struct {
    X int16
    Y int16
}

type SMALL_RECT struct {
    X1 int16
    Y1 int16
    X2 int16
    Y2 int16
}

func GetOutputHandle() (C.HANDLE) {
    return C.GetOutputHandle()
}

func GetInputHandle() (C.HANDLE) {
    return C.GetInputHandle()
}

func GetCh(hIn C.HANDLE) (byte) {
    return byte(C.getch(hIn))
}

func FGetCh(hIn _Ctype_HANDLE) (byte) {
    return byte(C.fgetch(hIn))
}

func GetText(hIn C.HANDLE, hOut C.HANDLE, pos COORD, maxlength int) (string) {
    text := C.getText(hIn, hOut, C.COORD{_Ctype_SHORT(pos.X),_Ctype_SHORT(pos.Y)}, C.int(maxlength))
    defer C.free(unsafe.Pointer(text))

    return C.GoString(text)
}

func Cls(hOut C.HANDLE) {
    C.cls(hOut)
}

func Clear(hOut C.HANDLE, from, to COORD) {
    C.clear(hOut, C.COORD{_Ctype_SHORT(from.X),_Ctype_SHORT(from.Y)},
            C.COORD{_Ctype_SHORT(to.X),_Ctype_SHORT(to.Y)})
}

//TODO Make the C function use a COORD instead of X and Y
func DrawChar(hOut C.HANDLE, chr byte, pos COORD) {
    C.drawChar(hOut, C.char(chr), C.int(pos.X), C.int(pos.Y))
}

func DrawText(hOut C.HANDLE, text string, align, xOff, yPos int) (COORD) {
    cText := C.CString(text)
    defer C.free(unsafe.Pointer(cText))

    pos := C.drawText(hOut, cText, C.int(align), C.int(xOff), C.int(yPos))
    return COORD{int16(pos.X), int16(pos.Y)}
}

func DrawLine(hOut C.HANDLE, chr byte, vertical int8, from COORD, count uint) {
    C.drawLine(hOut, C.char(chr), C.BOOL(vertical),
               C.COORD{_Ctype_SHORT(from.X),_Ctype_SHORT(from.Y)}, C.uint(count))
}

func DrawBox(hOut C.HANDLE, topLeft, bottomRight COORD) {
    C.drawBox(hOut, C.COORD{_Ctype_SHORT(topLeft.X),_Ctype_SHORT(topLeft.Y)},
              C.COORD{_Ctype_SHORT(bottomRight.X),_Ctype_SHORT(bottomRight.Y)})
}

//Wrappers for windows.h functions begin here

//Maybe make a "C.WinString" function
func SetConsoleTitle(text string) {
    cText := C.CString(text)
    defer C.free(unsafe.Pointer(cText))
    wText := C.chartoCHAR(cText)
    defer C.free(unsafe.Pointer(wText))

    C.SetConsoleTitle(wText)
}

func SetConsoleScreenBufferSize(hOut C.HANDLE, bSize COORD) (bool) {
    return int8(
        C.SetConsoleScreenBufferSize(
            hOut, C.COORD{_Ctype_SHORT(bSize.X),_Ctype_SHORT(bSize.Y)})) == 1
}

func SetConsoleWindowInfo(hOut C.HANDLE, absolute int8, wSize *SMALL_RECT) {
    C.SetConsoleWindowInfo(hOut, C.WINBOOL(absolute), &C.SMALL_RECT{
                           _Ctype_SHORT(wSize.X1),_Ctype_SHORT(wSize.Y1),
                           _Ctype_SHORT(wSize.X2),_Ctype_SHORT(wSize.Y2)})
}
