package main

import (
	"sync"
	"math/rand"
	"fmt"
	"sync/atomic"
	"runtime"
	"time"
)

//定义两个用来传递数据的结构，通道是用来接收返回值的，比如执行是否成功
type readOp struct{
	key int
	resp chan int
}

type writeOp struct{
	key int
	val int
	resp chan bool
}
//互斥锁，保证原子操作
func testMutex(){
	var stat = make(map[int]int)

	//互斥锁
	var mutex = &sync.Mutex{}
	//记录操作次数
	var ops int64 = 0
	//100个协程
	for r := 0; r < 100; r++{
		go func(){
			total := 0
			for{
				key := rand.Intn(5)
				mutex.Lock()
				total += stat[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++{
		go func(){
			key := rand.Intn(5)
			value := rand.Intn(100)
			mutex.Lock()
			stat[key] = value
			mutex.Unlock()
			atomic.AddInt64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second * 5)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops: ", opsFinal)
	mutex.Lock()
	fmt.Println("state:", stat)
	mutex.Unlock()
}


//第二个方法是使用通道的方法，使得一个变量只在一个协程中进行读写，而其他需要进行通信的可以通过通道传递

func testChan(){
    var ops int64
    reads := make(chan *readOp)
    writes := make(chan *writeOp)
    go func(){
    	var state = make(map[int]int)
    	for{
    		select{
    		case read := <- reads:
    			read.resp <- state[read.key]
    			case write := <- writes:
    				state[write.key] = write.val
    				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++{
		go func(){
			for{
				read := &readOp{key: rand.Intn(5),
				resp: make(chan int)}
				reads <- read
				<- read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++{
		go func(){
			for{
				write := &writeOp{
					key : rand.Intn(5),
					val: rand.Intn(100),
					resp : make(chan bool),
				}
				writes <- write
				<- write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second * 2)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops: ", opsFinal)

}
func gobyexamplechapter11(){
    //testMutex()
    testChan()
}