# ChessImager

## Embedded pieces

The embedded chess pieces are free for personal use, but not commercial use. See : https://clipart-library.com/clip-art/chess-pieces-silhouette-14.htm

Colorful chess pieces (Brian Provan) : Public Domain. See: https://opengameart.org/content/colorful-chess-pieces

## Todo 

* Create a CLI tool?
* Highlighted square x:s : https://elzr.com/blag/img/2018/chess-pieces/chess-moves.png
* Fix BoardImage (not implemented yet)
* Fix Move (foundations implemented)
* Handle error in hexToRGBA
* Implement PGN : White player, Black player, move count etc
* Settings should be a global variable, might simplify things 
* rendererRankAndFile should use getSquareBox for RankAndFileInSquare
* Validate FEN characters
* 
## Possible future todo:s
* Implement Possible Moves For/to square - show moves that a piece can do, or show pieces that can move to a square.
* Validation of settings file? CHeck if font file exists, size out of square boundary, etc
* Select corner for RankAndFileInSquare => RankAndFileTopLeft, RankAndFileTopRight, etc
