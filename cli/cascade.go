package cli

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"net/url"
)

func CascadeHandler(_ *url.URL) error {
	return Cascade()
}

// TODO: use batch operations
func Cascade() error {

	ca := nod.Begin("cascading reductions...")
	defer ca.End()

	props := []string{data.TitleProperty, data.BookCompletedProperty, data.ArtsHistoryOrderProperty}

	absReduxDir, err := pasu.GetAbsRelDir(data.Redux)
	if err != nil {
		return ca.EndWithError(err)
	}

	rdx, err := kvas.NewReduxWriter(absReduxDir, props...)
	if err != nil {
		return ca.EndWithError(err)
	}

	// cascading data.BookCompletedProperty
	bcpa := nod.NewProgress(" " + data.BookCompletedProperty)
	defer bcpa.End()

	ids := rdx.Keys(data.TitleProperty)
	bcpa.TotalInt(len(ids))

	for _, id := range ids {
		bcpa.Increment()
		if val, ok := rdx.GetFirstVal(data.BookCompletedProperty, id); ok && val != "" {
			continue
		}
		if err := rdx.ReplaceValues(data.BookCompletedProperty, id, "false"); err != nil {
			return ca.EndWithError(err)
		}
	}

	bcpa.EndWithResult("done")

	ca.EndWithResult("done")

	return nil
}
