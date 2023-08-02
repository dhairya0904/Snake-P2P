package main

type Config struct {
	RendezvousString string
	ProtocolID       string
	ListenHost       string
	ListenPort       int
	NodeType         string
	logLevel         string
	peerAddress      string
}

type Part struct {
	X int
	Y int
}

type GameStateUpdade struct {
	FoodPos          Part   `json:"foodPos"`
	Parts            []Part `json:"parts"`
	Xspeed           int    `json:"xSpeed"`
	Yspeed           int    `json:"ySpeed"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	OpponenetFoodPos Part   `json:"opponentFoodPos"`
}
