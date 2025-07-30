package models

import (
	"bytes"
	"image"
	"log"
	"retro-hamster/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameAction struct {
	Id         int
	Index      int
	Text       string
	Focused    bool
	Selected   bool
	Available  bool
	Counter    int
	X, Y, W, H int
	OnClickFn  func()
}

func (a *GameAction) Contains(x, y int) bool {
	if a == nil {
		return false
	}

	if x == -1 || y == -1 {
		return false
	}

	return x >= a.X && x <= a.X+a.W && y >= a.Y && y <= a.Y+a.H
}

func (a *GameAction) OnClick() {
	if a.OnClickFn != nil {
		a.OnClickFn()
	}
}

type GameController = uint8

const (
	GameController_None     GameController = 0
	GameController_Mouse    GameController = 1 << 0
	GameController_KeyBoard GameController = 1 << 2
	GameController_Gamepad  GameController = 1 << 3
)

type Game struct {
	ScreenW   int
	ScreenH   int
	UTIL_EXIT bool

	currentState IGameState

	TimerStart time.Time

	ImageAssets map[assets.AssetKey]GameAssetImg
	FontAssets  map[assets.AssetKey]GameAssetFont

	// Game Pad
	GamepadButtonIsDown bool
	GamepadIDsBuf       []ebiten.GamepadID
	GamepadIDs          map[ebiten.GamepadID]struct{}

	// Mouse
	CURSOR_X int
	CURSOR_Y int

	lastController uint8
}

func NewGame(w, h int) *Game {
	return &Game{
		ScreenW:             w,
		ScreenH:             h,
		UTIL_EXIT:           false,
		GamepadButtonIsDown: false,
		GamepadIDsBuf:       make([]ebiten.GamepadID, 0),
		GamepadIDs:          map[ebiten.GamepadID]struct{}{},
	}
}

func (g *Game) LoadGameAssets(allAssets map[assets.AssetKey]assets.AssetRef) {
	g.ImageAssets = make(map[assets.AssetKey]GameAssetImg)
	g.FontAssets = make(map[assets.AssetKey]GameAssetFont)

	for _, ref := range allAssets {

		if ref.AssetType == assets.AssetType_PNG {
			img, _, err := image.Decode(bytes.NewReader(ref.Data))
			if err != nil {
				log.Fatal(err)
			}
			ebImg := ebiten.NewImageFromImage(img)

			g.ImageAssets[ref.Key] = GameAssetImg{
				Asset: ref,
				Image: ebImg,
			}
		} else if ref.AssetType == assets.AssetType_TTF {
			f, err := text.NewGoTextFaceSource(bytes.NewReader(ref.Data))
			if err != nil {
				log.Fatal(err)
			}

			g.FontAssets[ref.Key] = GameAssetFont{
				Asset: ref,
				Font:  f,
			}
		}
	}
}

func (g *Game) Update() error {
	// Gamepad Connection Events
	g.GamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.GamepadIDsBuf[:0])

	for _, id := range g.GamepadIDsBuf {
		// log.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
		g.GamepadIDs[id] = struct{}{}
	}
	for id := range g.GamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			// log.Printf("gamepad disconnected: id: %d", id)
			delete(g.GamepadIDs, id)
		}
	}

	return g.currentState.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func (g *Game) ChangeState(state IGameState) {
	g.currentState = state
	state.OnTransition()
}

func (g *Game) ResetTimer() {
	g.TimerStart = time.Now()
}

func (g *Game) TimerCount() time.Duration {
	return time.Since(g.TimerStart)
}

func (g *Game) SetLastController(controller GameController) {
	g.lastController = controller
}

func (g *Game) GetLastController() GameController {
	return g.lastController
}

func (g *Game) OnMouseMoved(run func(x, y int)) (cursorMoved bool) {
	/** Mouse **/
	x, y := ebiten.CursorPosition()
	cursorMoved = false

	if x != g.CURSOR_X || y != g.CURSOR_Y {
		cursorMoved = true
	}

	if cursorMoved {
		run(x, y)

		g.CURSOR_X = x
		g.CURSOR_Y = y

		g.SetLastController(GameController_Mouse)
	}

	return cursorMoved
}

func (g *Game) OnMouseLeftClick(run func(x, y int)) (mouseClicked bool) {
	/** Mouse **/
	mouseClicked = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)

	if mouseClicked {
		x, y := ebiten.CursorPosition()
		run(x, y)
		g.SetLastController(GameController_Mouse)
	}

	return mouseClicked
}

func (g *Game) HandleGamePadEvents(button ebiten.StandardGamepadButton, handler func()) (handled bool) {
	for id := range g.GamepadIDs {
		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
				if inpututil.IsStandardGamepadButtonJustPressed(id, b) && !g.GamepadButtonIsDown {
					if b == button {
						handler()
						g.GamepadButtonIsDown = true
						return true
					}
				}
				if inpututil.IsStandardGamepadButtonJustReleased(id, b) {
					g.GamepadButtonIsDown = false
				}
			}
		}
	}

	return false
}
