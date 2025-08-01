package scenes

import (
	"image"
	"retro-hamster/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// DrawSprite draws a sprite from the sprite sheet
func DrawSprite(spritesheet *ebiten.Image, screen *ebiten.Image, x, y int, spriteX, spriteY, spriteWidth, spriteHeight int) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	// Draw the sub-image of the sprite sheet at the desired position
	screen.DrawImage(spritesheet.SubImage(image.Rect(spriteX, spriteY, spriteX+spriteWidth, spriteY+spriteHeight)).(*ebiten.Image),
		&opts,
	)
}

func DrawAssetSprite(spritesheet *ebiten.Image, screen *ebiten.Image, x, y int, asset assets.AssetSprite) {
	DrawSprite(spritesheet, screen, x, y, asset.X, asset.Y, asset.W, asset.H)
}

// DrawSprite draws a sprite from the sprite sheet
func DrawAssetSpriteWithOptions(spritesheet *ebiten.Image, screen *ebiten.Image, asset assets.AssetSprite, opts ebiten.DrawImageOptions) {
	screen.DrawImage(spritesheet.SubImage(image.Rect(asset.X, asset.Y, asset.X+asset.W, asset.Y+asset.H)).(*ebiten.Image),
		&opts,
	)
}

func DrawAssetSpriteWithOptionsWithBoundsCorrect(spritesheet *ebiten.Image, screen *ebiten.Image, asset assets.AssetSprite, opts ebiten.DrawImageOptions) {
	sub := spritesheet.SubImage(image.Rect(asset.X, asset.Y, asset.X+asset.W, asset.Y+asset.H)).(*ebiten.Image)
	opts.GeoM.Translate(float64(sub.Bounds().Dx()), 0) // Corrects after flipping
	screen.DrawImage(sub,
		&opts,
	)
}
