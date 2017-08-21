// This sample program demonstrates how to use an unbuffered
// channel to simulate a game of tennis between two goroutines.
// 此示例 展示如何使用无缓冲的通道
// 来模拟2个 goroutine之间的网球比赛
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.
	court := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// 如果 通道赋值 放到这里，那么就会发生死锁，因为此时所有的goroutine 都还没执行起来
	// goroutine里面可以处于阻塞状态
	//court <- 1

	// Launch two players.
	go player("Nadal", court)
	go player("Djokovic", court)
	//go player("Gongyao", court)

	// Start the set.
	// 为什么要放到这个位置？？？
	court <- 1

	// Wait for the game to finish.
	wg.Wait()
}

// player simulates a person playing the game of tennis.
// player模拟一个选手在打网球
func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done.
	// 函数结束时调用
	defer wg.Done()


	for {
		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		// 选择随机数，用这个随机数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
