package main

import (
	"struct2table"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("struct2table", func(w *unison.Window) {
		struct2table.New().Layout(w.Content())
	})
}
