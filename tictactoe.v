import readline { read_line }

fn main() {
	mut player := "Player 1"
	mut player_symbol := symbol(player)
	mut game := ["0", "1", "2", "3", "4", "5", "6", "7", "8"]
	mut running := true
	mut choice := 0
	mut disp_game := format_game(game)

	println(disp_game)

	for running {
		choice = read_line("\nEnter a number between 0 and 8 :\n> ")?.int()

		if choice <= 8 && choice >= 0 {

			if game[choice] != "O" && game[choice] != "X" {


				game[choice] = player_symbol
				disp_game = format_game(game)

				if won(player_symbol, game) {
					println("\n$player WON THE GAME !\nGG!")
					running = false
					println(disp_game)
					read_line("Press enter to exit ...")?

				} else{
				player = if player == "Player 1" { "Player 2" } else { "Player 1" }
				player_symbol = symbol(player)
				

				println("\n\n\n\n\nIt's now to :\n$player\n\n$disp_game\n")
				}

			} else{
				println("\nYou can't do this. Retry.")
			}
		} else{
			println("\nYou can't do this. Retry.")
		}
	}


}

fn symbol(player string) string {
	return if player == "Player 1" { "X" } else { "O" }
}

fn won(symbol string, game []string) bool{
	
	return (game[0] == symbol && game[1] == symbol && game[2] == symbol) || (game[3] == symbol && game[4] == symbol && game[5] == symbol) || (game[6] == symbol && game[7] == symbol && game[8] == symbol) || (game[0] == symbol && game[4] == symbol && game[8] == symbol) || (game[6] == symbol && game[4] == symbol && game[2] == symbol) || (game[0] == symbol && game[3] == symbol && game[6] == symbol) || (game[1] == symbol && game[4] == symbol && game[7] == symbol) || (game[2] == symbol && game[5] == symbol && game[8] == symbol)
}

fn format_game(game []string) string {
	return game[6] + " | " + game[7] + " | " + game[8] + "\n" + game[3] + " | " + game[4] + " | " + game[5] + "\n" + game[0] + " | " + game[1] + " | " + game[2]
}