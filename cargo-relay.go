package main

import (
        "bufio"
        "net"
        "strconv"
        "strings"
        "fmt"
        "lang"
        "time"
)
type Metric struct {
    Mkey string  // metric key
    Mval int64   // metric value
    Mts  int     // metric unix-epoch
}

func main() {
        ln, err := net.Listen("tcp", "0.0.0.0:6666")
        if err != nil {
                // handle error
        }
        queue := lang.NewQueue()
        go consume(queue)
        for {
                conn, err := ln.Accept()
                if err != nil {
                        continue
                }
                go handle(conn, queue)
        }
        fmt.Println("End\n")
}

func consume(queue *lang.Queue) {
    for {
      time.Sleep(2000 * time.Millisecond)
      fmt.Println("Total items: ", queue.Len())
    }
}

func handle(conn net.Conn, queue *lang.Queue) {
        fmt.Println("Connection: %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())
        defer func() {
                fmt.Println("Closing connection: %s\n", conn.RemoteAddr())
                conn.Close()
        }()

        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
                if m := strings.Split(scanner.Text(), " "); len(m) > 2 {
                        if ts, err := strconv.ParseInt(m[2], 10, 0); err != nil {
                                // handling warning
                                continue
                        } else {
                                m := &Metric{"test.test1", 1, 2}
                                fmt.Println("Metric ", ts, m)
                                queue.Push(m)
                        }
                }
        }

        if err := scanner.Err(); err != nil {
                // handling error
        }
}