package main

import "../losegui"
import "fmt"

func main() {
    hOut := losegui.GetOutputHandle()
    hIn := losegui.GetInputHandle()
    losegui.SetConsoleTitle("test")
    losegui.Cls(hOut)
    fmt.Printf("Press any key to continue.")
    losegui.GetCh(hIn)
    losegui.GetText(hIn, hOut, losegui.COORD{13, 13}, 30)
    for i:=0;i<10;i++ {
        losegui.DrawChar(hOut, 'x', losegui.COORD{int16(i), int16(i)})
    }
    losegui.DrawLine(hOut, 'x', 0, losegui.COORD{0, 10}, 40)
    losegui.DrawBox(hOut, losegui.COORD{0, 0}, losegui.COORD{89, 15})
    losegui.DrawText(hOut, "Hello World! :D", 2, 0, 1)
    losegui.DrawText(hOut, "Welcome to formatted text!!!!!!", 2, 0, 3)
    losegui.DrawText(hOut, "Press any key to continue.", 0, 0, 16)
    losegui.GetCh(hIn)
}
