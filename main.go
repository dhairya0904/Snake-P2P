package main

import (
	"time"

	"github.com/rs/zerolog"
)

func main() {

	cfg := parseFlags()
	if cfg.logLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	node := Node{
		RendezvousString: cfg.RendezvousString,
		ListenHost:       cfg.ListenHost,
		ListenPort:       cfg.ListenPort,
		ProtocolID:       cfg.ProtocolID,
	}
	node.InitializeNode()
	host := node.CreateHost()
	node.InitializeNode()

	if len(cfg.peerAddress) == 0 {
		node.startMaster(host)
		time.Sleep(5 * time.Second)
	} else {
		node.connectWithPeer(host, cfg.peerAddress)
	}
	for {

	}
	// screen, err := tcell.NewScreen()

	// if err != nil {
	// 	log.Fatalf("%+v", err)
	// }
	// if err := screen.Init(); err != nil {
	// 	log.Fatalf("%+v", err)
	// }

	// defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	// screen.SetStyle(defStyle)

	// game := Game{
	// 	Screen: screen,
	// 	Node:   &node,
	// }

	// go game.Run()
	// for {
	// 	switch event := game.Screen.PollEvent().(type) {
	// 	case *tcell.EventResize:
	// 		game.Screen.Sync()
	// 	case *tcell.EventKey:
	// 		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
	// 			game.Screen.Fini()
	// 			os.Exit(0)
	// 		} else if event.Key() == tcell.KeyUp && game.snakeBody.Yspeed == 0 {
	// 			game.snakeBody.ChangeDir(-1, 0)
	// 		} else if event.Key() == tcell.KeyDown && game.snakeBody.Yspeed == 0 {
	// 			game.snakeBody.ChangeDir(1, 0)
	// 		} else if event.Key() == tcell.KeyLeft && game.snakeBody.Xspeed == 0 {
	// 			game.snakeBody.ChangeDir(0, -1)
	// 		} else if event.Key() == tcell.KeyRight && game.snakeBody.Xspeed == 0 {
	// 			game.snakeBody.ChangeDir(0, 1)
	// 		} else if event.Rune() == 'y' && game.GameOver {
	// 			go game.Run()
	// 		} else if event.Rune() == 'n' && game.GameOver {
	// 			game.Screen.Fini()
	// 			os.Exit(0)
	// 		}
	// 	}
	// }
}
