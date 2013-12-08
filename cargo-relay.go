package main

import (
        "bufio"
        "net"
        "strconv"
        "strings"
        "fmt"
        "lang"
        "time"
        "gcfg"
        "github.com/VividCortex/ewma"

)
type Metric struct {
    Mkey string  // metric key
    Mval int   // metric value
    Mts  int     // metric unix-epoch
}

type Config struct {
        Global struct {
            Hosts string
            Port string
        }
        Queue struct {
            Max int
        }
}

func main() {
        var cfg Config
        err := gcfg.ReadFileInto(&cfg, "cargo.conf")
        address := fmt.Sprintf("%s:%s", []byte(cfg.Global.Hosts), []byte(cfg.Global.Port))
        ln, err := net.Listen("tcp", address)
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
    e := ewma.NewMovingAverage()
    for {
        old_len := queue.Len()
        time.Sleep(1000 * time.Millisecond)
        new_len := queue.Len()
        msg_cnt := float64(new_len - old_len)
        e.Add(msg_cnt)
        str := fmt.Sprintf("Msg cnt: %v | Msg send: %v | ewma: %v", new_len, msg_cnt, e.Value())
        fmt.Println(str)
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
                        ts,_ := strconv.ParseInt(m[2], 10, 0)
                        val,_ := strconv.Atoi(m[1])
                        m := &Metric{m[0], val, int(ts)}
                        //fmt.Println("Metric ", ts, m)
                        queue.Push(m)
                }
        }

        if err := scanner.Err(); err != nil {
                // handling error
        }
}