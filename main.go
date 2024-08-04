package main

import (
	"struct2table"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("struct2table", func(w *unison.Window) {
		w.Content().AddChild(struct2table.New().Layout())
	})
}
