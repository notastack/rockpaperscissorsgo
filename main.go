package main

import (
        "fmt"
        "math/rand"
        "time"
	"log"
	"os/user"
	"os"
)

func choice(v int) {
	//convert the number to the RPS choice
	choice := []string{"Rock","Paper","Scissors"}
        fmt.Println(choice[v - 1])
}

func result(u int , c int) {
	//tells you if you won or not
    if u == c { 
        fmt.Println("It's a tie. Try again next time.")
        return
    }
	//use a circular array to calculate the results because elif are ugly, take too much space and really slow
	possible := []int{1,2,3}
	var rot int
	rot = u -1
	rotateL(possible , rot)
	enemy := possible[1]
	if c == enemy {
        fmt.Println("You lost. Better luck next time.")
	}else {
        fmt.Println("You won, congrats!")		
	}
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func rotateL(a []int, i int) {
	//rotate the array to the left * i
    i = i % len(a)
    if i < 0 {
        i += len(a)
    }
    for c := 0; c < gcd(i, len(a)); c++ {
        t := a[c]
        j := c
        for {
            k := j + i
            if k >= len(a) {
                k -= len(a)
            }
            if k == c {
                break
            }
            a[j] = a[k]
            j = k
        }
        a[j] = t
    }
}

func main() {
		//the game
		//also you lost
		var usrplay int
		// var cmtplay
		user, err := user.Current()
		//introduce the player
		username := user.Username
		fmt.Printf("Hello %s\n", username)
		fmt.Println("Lets play Rock Paper Scissors")
		//you chose what you want to use
		fmt.Println("Type 1 for Rock, 2 for Paper or 3 for Scissors")
		fmt.Scanln(&usrplay)
		fmt.Println("You choose")
		choice(usrplay)
		fmt.Println("The computer choose")
		//builds up the hype
		time.Sleep(1 * time.Second)
		fmt.Println("3...")
		time.Sleep(1 * time.Second)
		fmt.Println("2...")
		time.Sleep(1 * time.Second)
		fmt.Println("1...")
		time.Sleep(1 * time.Second)
		var cpuplay int
		//generate a random choice for the cpu
		rand.Seed(time.Now().UnixNano())
        	cpuplay = rand.Intn(3) + 1
// for testing purposes		cpuplay = 1
		choice(cpuplay)
		//output the result
		result(usrplay , cpuplay)

}
