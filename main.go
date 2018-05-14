package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//Cons is the body of concrete
type Cons struct {
	Height   float64
	Width    float64
	As0      float64
	DegreeID int
	Scale    string
	Moment   float64
	Shear    float64
	Torque   float64
}

//Degree is only used by Binder
type Degree struct {
	ID  int
	Val string
}

//Value Nothing to say
func Value() []*Degree {
	return []*Degree{
		{0, "C25"},
		{1, "C30"},
		{2, "C35"},
		{3, "C40"},
	}
}

var concrete = new(Cons)

func main() {
	// walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
	walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
	var db *walk.DataBinder
	var outTe *walk.TextEdit
	ico, _ := walk.NewIconFromFile(".\\3dEdgeON.ico")
	MainWindow{
		Title:   "Concrete Calc",
		Icon:    ico,
		Layout:  VBox{},
		MinSize: Size{400, 600},
		DataBinder: DataBinder{
			AssignTo:   &db,
			DataSource: concrete,
			AutoSubmit: true,
		},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2, Spacing: 2},
				Children: []Widget{
					Label{Text: "Width(mm):"},
					NumberEdit{Value: Bind("Width")},
					Label{Text: "Height(mm):"},
					NumberEdit{Value: Bind("Height")},
					Label{Text: "Moment(kN·m):"},
					NumberEdit{Value: Bind("Moment")},
					Label{Text: "Shear(kN):"},
					NumberEdit{Value: Bind("Shear")},
					Label{Text: "Torque(kN·m):"},
					NumberEdit{Value: Bind("Torque")},
					Label{Text: "Degree:"},
					ComboBox{
						Value:         Bind("DegreeID"),
						CurrentIndex:  1,
						DisplayMember: "Val",
						BindingMember: "Id",
						Model:         Value(),
					},
					Label{Text: "Scale:"},
					ComboBox{
						Editable:     true,
						Value:        Bind("Scale"),
						CurrentIndex: 1,
						Model:        []string{"1.1", "1.25", "1.3", "1.35"},
					},
					Label{Text: "As0:"},
					NumberEdit{
						Suffix: " mm",
						Value:  Bind("As0"),
					},
				},
			},
			VSeparator{},
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{Text: "Resule:"},
					HSpacer{},
				},
			},
			TextEdit{
				AssignTo: &outTe,
				ReadOnly: true,
			},
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "Calc",
						OnClicked: func() {
							outTe.SetText(concrete.Calc())
						},
					},
				},
			},
		},
	}.Run()
}
