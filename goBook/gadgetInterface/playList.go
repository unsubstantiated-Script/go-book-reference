package gadgetInterface

import "goBook/goBook/gadgetInterface/utils"

// The interface
type Player interface {
	// The things required
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	//Asserting this type to make it work
	recorder, ok := player.(utils.TapeRecorder)

	// Safety checking an interface against it's accepted type. Avoids a runtime panic.
	if ok {
		recorder.Record()
	}
	player.Play("Test Track")
	recorder.Stop()
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
	var player Player = utils.TapePlayer{}
	playList(player, mixtape)
	player = utils.TapeRecorder{}
	playList(player, mixtape)

	TryOut(utils.TapeRecorder{})
	TryOut(utils.TapePlayer{})
}
