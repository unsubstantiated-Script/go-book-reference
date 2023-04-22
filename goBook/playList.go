package goBook

import "goBook/goBook/gadget"

// The interface
type Player interface {
	// The things required
	Play(string)
	Stop()
}

//Passing in the interface and the songs
func playList(device Player, songs []string) {
	// Playing through the list of songs
	for _, song := range songs {
		device.Play(song)
	}

	// Then stopping
	device.Stop()
}

func PlaySongs() {
	mixtape := []string{"Jessie's Girl", "Tootsie Roll", "Be Like Mike"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}
