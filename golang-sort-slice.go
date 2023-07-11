	players := []struct {
		Name string
		Wins int
	}{
		{"Chris", 10},
		{"Cleo", 30},
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Wins > players[j].Wins
	})

	fmt.Printf("%v", players)
	// [{Cleo 30} {Chris 10}]
