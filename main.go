package main

import (
	"struct2table"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("struct2table", func(w *unison.Window) {
		w.Content().AddChild(struct2table.New().Layout())
	})
}
