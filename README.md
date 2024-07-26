# goTetris
A modern Tetris game written in Go.

## Code

This project uses the ebiten library as a game engine.
The update loop is divided between  :
- the update function (in the update.go file) which handles the game logic 
- the draw function (in the draw.go file) which displays the game content

## Features

This game includes most of the modern features of a Tetris game, including:
- delayed auto-shift (DAS)
- Hold piece
- Hard drops (with a ghost piece for guidance)
- Instant soft drops
- SRS rotation system (with wall kicks)
- 7-bag randomizer
- next piece preview (up to 5 pieces)