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
    switch v {
        case 1:
        fmt.Println("Rock")
        case 2:
        fmt.Println("Paper")
	case 3:
        fmt.Println("Scissors")
        default:
        fmt.Println("Nothing. You're not playing nice!")
	os.Exit(0)    
    }
}

func result(u int , c int) {
    if u == c { 
        fmt.Println("It's a tie. Try again next time.")
        return
    }
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

		var usrplay int
		// var cmtplay
		user, err := user.Current()
		if err != nil {
		log.Fatalf(err.Error())
		}
		username := user.Username
		fmt.Printf("Hello %s\n", username)
		fmt.Println("Lets play Rock Paper Scissors")
		fmt.Println("Type 1 for Rock, 2 for Paper or 3 for Scissors")
		fmt.Scanln(&usrplay)
		fmt.Println("You choose")
		choice(usrplay)
		fmt.Println("The computer choose")	
		time.Sleep(1 * time.Second)
		fmt.Println("3...")
		time.Sleep(1 * time.Second)
		fmt.Println("2...")
		time.Sleep(1 * time.Second)
		fmt.Println("1...")
		time.Sleep(1 * time.Second)
		var cpuplay int
		rand.Seed(time.Now().UnixNano())
        	cpuplay = rand.Intn(3) + 1
// for testing purposes		cpuplay = 1
		choice(cpuplay)
		result(usrplay , cpuplay)

}
