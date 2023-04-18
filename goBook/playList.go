package goBook

import "goBook/goBook/gadget"

func playList(device gadget.TapePlayer, songs []string) {

	// Playing through the list of songs
	for _, song := range songs {
		device.Play(song)
	}

	// Then stopping
	device.Stop()
}

func PlaySongs() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Tootsie Roll", "Be Like Mike"}
	playList(player, mixtape)
}
