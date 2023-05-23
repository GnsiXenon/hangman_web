package hangmanWeb

import (
	HangMan "github.com/GnsiXenon/HangMan/src"
)

type Loop struct {
	Letter  string
	Visible bool
}
type DataHangman struct {
	Word      string
	Stock     string
	Attempt   int
	InWord    []bool
	Letter    string
	HidenWord string
	Text      string
	Keyboard1 []Loop
	Keyboard2 []Loop
	Keyboard3 []Loop
	GameStart bool
	file      string
	AddWord   bool
}

var data = DataHangman{
	GameStart: false,
}

func Initial(file string) DataHangman {
	guessWord, t := HangMan.StarGame(file)
	hW := HangMan.Prints(guessWord, t)
	data.Word = guessWord
	data.Stock = ""
	data.Attempt = 10
	data.InWord = t
	data.Letter = ""
	data.HidenWord = hW
	data.Text = ""
	data.file = file
	data.AddWord = false
	data.Keyboard1 = []Loop{
		{Letter: "A", Visible: true},
		{Letter: "Z", Visible: true},
		{Letter: "E", Visible: true},
		{Letter: "R", Visible: true},
		{Letter: "T", Visible: true},
		{Letter: "Y", Visible: true},
		{Letter: "U", Visible: true},
		{Letter: "I", Visible: true},
		{Letter: "O", Visible: true},
		{Letter: "P", Visible: true},
	}
	data.Keyboard2 = []Loop{
		{Letter: "Q", Visible: true},
		{Letter: "S", Visible: true},
		{Letter: "D", Visible: true},
		{Letter: "F", Visible: true},
		{Letter: "G", Visible: true},
		{Letter: "H", Visible: true},
		{Letter: "J", Visible: true},
		{Letter: "K", Visible: true},
		{Letter: "L", Visible: true},
		{Letter: "M", Visible: true},
	}
	data.Keyboard3 = []Loop{
		{Letter: "W", Visible: true},
		{Letter: "X", Visible: true},
		{Letter: "C", Visible: true},
		{Letter: "V", Visible: true},
		{Letter: "B", Visible: true},
		{Letter: "N", Visible: true},
	}
	data.GameStart = true
	return data
}
func GoodLetter(letter string) DataHangman {
	data.Letter = letter
	data.HidenWord, data.Attempt, data.InWord, data.Stock, data.Text = HangMan.Guess(data.Word, data.Letter, data.InWord, data.Stock, data.Attempt)
	return data
}
func CaseStatus(guess string) DataHangman {
	fl := data.Keyboard1
	sl := data.Keyboard2
	tl := data.Keyboard3
	for i := 0; i < len(fl); i++ {
		if fl[i].Letter == guess {
			fl[i].Visible = false
		} else if sl[i].Letter == guess {
			sl[i].Visible = false
		}
	}
	for i := 0; i < len(tl); i++ {
		if tl[i].Letter == guess {
			tl[i].Visible = false
		}
	}
	return data
}

func AddWord(newWord string) DataHangman {
	data.Text, data.AddWord = HangMan.AddWords(data.file, newWord)
	return data
}

func GameStatus(letter string) (DataHangman, string) {
	data = GoodLetter(letter)
	if data.Attempt == 0 && HangMan.Success(data.InWord) == false {
		return data, "../template/lose.html"
	} else if HangMan.Success(data.InWord) == true {
		return data, "../template/win.html"
	} else {
		data = CaseStatus(letter)
	}
	return data, "../template/game.html"
}
