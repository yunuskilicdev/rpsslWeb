package main

const (
	lose = "lose"
	tie = "tie"
	win = "win"
)
var choices Choices

var results [6][6]string

func FillChoices() {
	rock := Choice{
		Id:   1,
		Name: "rock",
	}
	paper:=Choice{
		Id:   2,
		Name: "paper",
	}
	scissors := Choice{
		Id:   3,
		Name: "scissors",
	}
	lizard := Choice{
		Id:   4,
		Name: "lizard",
	}
	spock := Choice{
		Id:   5,
		Name: "spock",
	}

	choices = Choices{
		rock,
		paper,
		scissors,
		lizard,
		spock,
	}

	results[rock.Id][rock.Id] = tie
	results[rock.Id][paper.Id] = lose
	results[rock.Id][scissors.Id] = win
	results[rock.Id][lizard.Id] = win
	results[rock.Id][spock.Id] = lose

	results[paper.Id][rock.Id] = win
	results[paper.Id][paper.Id] = tie
	results[paper.Id][scissors.Id] = lose
	results[paper.Id][lizard.Id] = lose
	results[paper.Id][spock.Id] = win

	results[scissors.Id][rock.Id] = lose
	results[scissors.Id][paper.Id] = win
	results[scissors.Id][scissors.Id] = tie
	results[scissors.Id][lizard.Id] = win
	results[scissors.Id][spock.Id] = lose

	results[lizard.Id][rock.Id] = lose
	results[lizard.Id][paper.Id] = win
	results[lizard.Id][scissors.Id] = lose
	results[lizard.Id][lizard.Id] = tie
	results[lizard.Id][spock.Id] = win

	results[spock.Id][rock.Id] = win
	results[spock.Id][paper.Id] = lose
	results[spock.Id][scissors.Id] = win
	results[spock.Id][lizard.Id] = lose
	results[spock.Id][spock.Id] = tie
}
