package models

import (
	"image"
	"math"
	"retro-hamster/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// DrawSprite draws a sprite from the sprite sheet
func DrawSprite(spritesheet *ebiten.Image, screen *ebiten.Image, x, y float64, spriteX, spriteY, spriteWidth, spriteHeight int) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	// Draw the sub-image of the sprite sheet at the desired position
	screen.DrawImage(spritesheet.SubImage(image.Rect(spriteX, spriteY, spriteX+spriteWidth, spriteY+spriteHeight)).(*ebiten.Image),
		&opts,
	)
}

func DrawAssetSprite(spritesheet *ebiten.Image, screen *ebiten.Image, x, y float64, asset assets.AssetSprite) {
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

func DrawButton(g *Game, screen *ebiten.Image, button *GameAction) {
	ss, ssExists := g.ImageAssets[assets.AssetKey_Static_PNG]
	if !ssExists {
		return
	}

	btnSelectedRect := assets.Sprite_Button_Focused

	if button.Selected || button.Focused {
		DrawSprite(ss.Image, screen, button.X, button.Y, btnSelectedRect.X, btnSelectedRect.Y, btnSelectedRect.W, btnSelectedRect.H)
	} else {
		DrawAssetSprite(ss.Image, screen, button.X, button.Y, assets.Sprite_Button)
	}

	if font, fontExists := g.FontAssets[assets.AssetKey_Sunny_Font_TTF]; fontExists {
		txtOps := &text.DrawOptions{}
		txtOps.GeoM.Translate(float64(button.X)+float64(assets.Sprite_Button.W/2), float64(button.Y)+float64(assets.Sprite_Button.H)/3.5)
		txtOps.ColorScale.ScaleWithColor(COLOR_PINK)
		txtOps.PrimaryAlign = text.AlignCenter
		text.Draw(screen, button.Text, &text.GoTextFace{
			Source: font.Font,
			Size:   float64(assets.Sprite_Button.H)/2 - math.Mod(float64(assets.Sprite_Button.H), 16),
		}, txtOps)
	}

}

func WriteToScene(stringText string, size float64, g *Game, screen *ebiten.Image, opts *text.DrawOptions) {
	if font, fontExists := g.FontAssets[assets.AssetKey_Sunny_Font_TTF]; fontExists {
		txtOps := &text.DrawOptions{}
		*txtOps = *opts
		text.Draw(screen, stringText, &text.GoTextFace{
			Source: font.Font,
			Size:   size,
		}, txtOps)
	}
}
