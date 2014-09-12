

package main


import (
       "os"
       "fmt"
       )






func main() {

    msg := "No DB Specified, Y/N to create a new one"
    new, err:= CheckDB(msg)

}



func CheckDB() {

           msg := "No DB Specified, Y/N to create a new one"
           var i string
           fmt.Println(msg)
           fmt.Scan(&i)
           if i == "y" {
              CreateDB()
           } else {
              os.Exit(0)
           }

}


