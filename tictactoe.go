package main

import (
	"fmt"
)

func main() {
	
	player := "Player 1"
	player_symbol := symbol(player)
	game := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}

	running := true
	var win bool
	var choice uint8
	var disp_game string

	disp_game = format_game(game)
	fmt.Println(disp_game)

	for running {

		fmt.Printf("\nEnter a number between 0 and 8 :\n> ")
		fmt.Scanln(&choice)
		
		if choice <= 8 && choice >= 0 {

			if game[choice] != "O" && game[choice] != "X" {


				game[choice] = player_symbol

				win = won(player_symbol, game)
				
				if win {
					fmt.Printf("\n%v WON THE GAME !\nGG!", player)
					running = false
					fmt.Scanf("Press enter to exit ...")
				} else{
				player = next_player(player)
				player_symbol = symbol(player)
				
				disp_game = format_game(game)
				fmt.Printf("\n\n\n\n\nIt's now to :\n%v\n\n%v\n", player, disp_game)
				}

			} else{
				fmt.Printf("\nYou can't do this. Retry.")
			}
		} else{
			fmt.Printf("\nYou can't do this. Retry.")
		}
	}

}

func symbol(player string) string{
	if player == "Player 1"{
		return "O"
	} else{
		return "X"
	}
	
}

func next_player(player string) string{

	if player == "Player 1" {

		return "Player 2"
	
	} else {
	
		return "Player 1"
	}
}

func won(symbol string, game [9]string) bool{
	
	return (game[0] == symbol && game[1] == symbol && game[2] == symbol) || (game[3] == symbol && game[4] == symbol && game[5] == symbol) || (game[6] == symbol && game[7] == symbol && game[8] == symbol) || (game[0] == symbol && game[4] == symbol && game[8] == symbol) || (game[6] == symbol && game[4] == symbol && game[2] == symbol) || (game[0] == symbol && game[3] == symbol && game[6] == symbol) || (game[1] == symbol && game[4] == symbol && game[7] == symbol) || (game[2] == symbol && game[5] == symbol && game[8] == symbol)
}

func format_game(game [9]string) string{
	return game[6] + " | " + game[7] + " | " + game[8] + "\n" + game[3] + " | " + game[4] + " | " + game[5] + "\n" + game[0] + " | " + game[1] + " | " + game[2]
}