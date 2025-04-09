package main

import (
    "fmt"
    "net"
)
var dict = map[string]string{ 
    "red": "красный",
    "green": "зеленый",
    "blue": "синий",
    "yellow": "желтый",
}

     
func main() {
    listener, err := net.Listen("tcp", ":4545") 
     
    if err != nil {
        fmt.Println(err) 
        return
    } 
    defer listener.Close() 
    fmt.Println("Server is listening...")
    for { 
        conn, err := listener.Accept() 
        if err != nil { 
            fmt.Println(err) 
            conn.Close() 
            continue
        } 
        go handleConnection(conn)  
    } 
}

func handleConnection(conn net.Conn) { 
    defer conn.Close()
    for {
        input := make([]byte, (1024 * 4)) 
        n, err := conn.Read(input) 
        if n == 0 || err != nil {
            fmt.Println("Read error:", err)
            break
        }    
        source := string(input[0:n])
        target, ok := dict[source]
        if !ok{           
            target = "undefined"
        }
        fmt.Println(source, "-", target)
        conn.Write([]byte(target))
    }
}