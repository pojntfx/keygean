package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type PasswordModalStory struct {
	Story

	modalOpen bool
}

func (c *PasswordModalStory) Render() app.UI {
	c.EnableShallowReflection()

	return c.WithRoot(
		app.Div().Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("Open Password Modal").
				OnClick(func(ctx app.Context, e app.Event) {
					c.modalOpen = !c.modalOpen
				}),
			app.If(
				c.modalOpen,
				&components.PasswordModal{
					OnSubmit: func(password string) {
						app.Window().Call("alert", "Successfully entered a password")

						c.modalOpen = false
					},
					OnCancel: func() {
						c.modalOpen = false

						c.Update()
					},
				},
			),
		),
	)
}
