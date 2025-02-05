package mouse

import (
	"github.com/go-vgo/robotgo"
	"github.com/haroflow/go-macros/automation"
)

func Commands() []automation.Command {
	moduleName := "mouse"
	return []automation.Command{
		{
			ModuleName:  moduleName,
			MethodName:  "move",
			Parameters:  "x: int, y: int",
			Description: "Moves the mouse cursor to a point on the screen in one step.",
			Action:      Move,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "moveSmooth",
			Parameters:  "x: int, y: int",
			Description: "Moves the mouse cursor to a point on the screen smoothly.",
			Action:      MoveSmooth,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "moveRelative",
			Parameters:  "x: int, y: int",
			Description: "Moves the mouse cursor to a point relative to the current mouse position in one step.",
			Action:      MoveRelative,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "click",
			Parameters:  "",
			Description: "Triggers a left click.",
			Action:      Click,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "rightClick",
			Parameters:  "",
			Description: "Triggers a right click.",
			Action:      RightClick,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "doubleClick",
			Parameters:  "",
			Description: "Triggers a left double-click.",
			Action:      DoubleClick,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "leftDown",
			Parameters:  "",
			Description: "Press the left mouse button.",
			Action:      LeftDown,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "leftUp",
			Parameters:  "",
			Description: "Unpress the left mouse button.",
			Action:      LeftUp,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "rightDown",
			Parameters:  "",
			Description: "Press the left mouse button.",
			Action:      RightDown,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "rightUp",
			Parameters:  "",
			Description: "Unpress the left mouse button.",
			Action:      RightUp,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "drag",
			Parameters:  "x: int, y: int",
			Description: "Press the left mouse button on the current position and drag to another position on screen.",
			Action:      Drag,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "dragRelative",
			Parameters:  "x: int, y: int",
			Description: "Press the left mouse button on the current position and drag to another position on screen relative to the current position.",
			Action:      DragRelative,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "getX",
			Parameters:  "",
			Description: "Returns the current mouse X position.",
			Action:      GetX,
		},
		{
			ModuleName:  moduleName,
			MethodName:  "getY",
			Parameters:  "",
			Description: "Returns the current mouse Y position.",
			Action:      GetY,
		},
	}
}

func Move(x, y int) {
	robotgo.Move(x, y)
}

func MoveRelative(x, y int) {
	robotgo.MoveRelative(x, y)
}

func MoveSmooth(x, y int) {
	robotgo.MoveSmooth(x, y)
}

func Click() {
	robotgo.Click("left", false)
}

func DoubleClick() {
	robotgo.Click("left", true)
}

func RightClick() {
	robotgo.Click("right", false)
}

func LeftDown() {
	_ = robotgo.MouseDown("left")
}

func RightDown() {
	_ = robotgo.MouseDown("right")
}

func LeftUp() {
	_ = robotgo.MouseUp("left")
}

func RightUp() {
	_ = robotgo.MouseUp("right")
}

func GetX() int {
	x, _ := robotgo.Location()
	return x
}

func GetY() int {
	_, y := robotgo.Location()
	return y
}

func Position() (x, y int) {
	return robotgo.Location()
}

func Drag(x, y int) {
	robotgo.DragSmooth(x, y)
}

func DragRelative(x, y int) {
	dx := GetX() + x
	dy := GetY() + y
	robotgo.DragSmooth(dx, dy)
}
