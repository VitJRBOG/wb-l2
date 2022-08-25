package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

func ExecuteStateExample() {
	player := NewMusicPlayer()

	var s string

	for {
		fmt.Println("Print new state: ")
		fmt.Scan(&s)

		switch s {
		case "1":
			player.Play()
		case "2":
			player.Pause()
		case "3":
			player.Stop()
		}
	}
}

type MusicPlayer struct {
	playback State
	paused   State
	stopped  State

	currentState State
}

func NewMusicPlayer() *MusicPlayer {
	player := &MusicPlayer{}

	playbackState := &PlaybackState{
		MusicPlayer: player,
	}

	pausedState := &PausedState{
		MusicPlayer: player,
	}

	stoppedState := &StoppedState{
		MusicPlayer: player,
	}

	player.SetState(stoppedState)
	player.playback = playbackState
	player.paused = pausedState
	player.stopped = stoppedState

	return player
}

func (p *MusicPlayer) Play() {
	p.currentState.Playback()
}

func (p *MusicPlayer) Pause() {
	p.currentState.Paused()
}

func (p *MusicPlayer) Stop() {
	p.currentState.Stopped()
}

func (p *MusicPlayer) SetState(s State) {
	p.currentState = s
}

type State interface {
	Playback()
	Paused()
	Stopped()
}

type PlaybackState struct {
	MusicPlayer *MusicPlayer
}

func (s *PlaybackState) Playback() {
	fmt.Println("Playback is already running")
}

func (s *PlaybackState) Paused() {
	s.MusicPlayer.SetState(s.MusicPlayer.paused)
	fmt.Println("Playback has been paused")
}

func (s *PlaybackState) Stopped() {
	s.MusicPlayer.SetState(s.MusicPlayer.stopped)
	fmt.Println("Playback has been stopped")
}

type PausedState struct {
	MusicPlayer *MusicPlayer
}

func (s *PausedState) Playback() {
	s.MusicPlayer.SetState(s.MusicPlayer.playback)
	fmt.Println("Playback is resumed")
}

func (s *PausedState) Paused() {
	fmt.Println("Playback is already paused")
}

func (s *PausedState) Stopped() {
	s.MusicPlayer.SetState(s.MusicPlayer.stopped)
	fmt.Println("Playback has been stopped")
}

type StoppedState struct {
	MusicPlayer *MusicPlayer
}

func (s *StoppedState) Playback() {
	s.MusicPlayer.SetState(s.MusicPlayer.playback)
	fmt.Println("Playback is resumed")
}

func (s *StoppedState) Paused() {
	s.MusicPlayer.SetState(s.MusicPlayer.paused)
	fmt.Println("Playback has been paused")
}

func (s *StoppedState) Stopped() {
	fmt.Println("Playback is already stopped")
}
