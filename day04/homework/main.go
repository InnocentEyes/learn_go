package main

import "fmt"

var (
	coins = 50
	users = []string{"Mathhew","Sarah","Augustuss","Heidi","Emilie","Peter","Giana","Adriano","Aaron","Elizabeth"}
	distribution = make(map[string]int,len(users))
)

func dispatchCoin() (int,map[string]int) {
	for _, user := range users {
		distribution[user] = 0
		for i := 0; i < len(user); i++ {
			value := distribution[user]
			switch user[i] {
			case 'e', 'E':
				value++
				coins--
				distribution[user] = value
			case 'i', 'I':
				value += 2
				coins -= 2
				distribution[user] = value
			case 'o', 'O':
				value += 3
				coins -= 3
				distribution[user] = value
			case 'u', 'U':
				value += 4
				coins -= 4
				distribution[user] = value
			}
		}
	}
	return coins,distribution
}

func main() {
	much,res := dispatchCoin()
	fmt.Println(much)
	fmt.Println(res)
}
