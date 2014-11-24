

package rafiki


import (
        "os"
        "os/exec"
        )


func ClearScreen() {

    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()

}








