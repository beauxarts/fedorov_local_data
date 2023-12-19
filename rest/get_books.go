package rest

import (
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/fedorov/stencil_app"
	"github.com/boggydigital/nod"
	"golang.org/x/exp/slices"
	"net/http"
	"strconv"
	"time"
)

const (
	latestBooksLimit = 24
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	// GET /books

	showAll := r.URL.Query().Get("show-all") == "true"

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	myBooks, ok := rdx.GetAllValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
	if !ok {
		http.Error(w, nod.ErrorStr("no my books found"), http.StatusInternalServerError)
		return
	}

	if missingDetails, ok := rdx.GetAllValues(data.MissingDetailsIdsProperty, data.MissingDetailsIdsProperty); ok {
		filteredBooks := make([]string, 0, len(myBooks))
		for _, id := range myBooks {
			if slices.Contains(missingDetails, id) {
				continue
			}
			filteredBooks = append(filteredBooks, id)
		}
		myBooks = filteredBooks
	}

	booksByType := make(map[string][]string)
	bookTypeTotals := make(map[string]int)

	for _, id := range myBooks {
		bt, _ := rdx.GetFirstVal(data.BookTypeProperty, id)
		bookTypeTotals[bt]++
		if !showAll && len(booksByType[bt]) >= latestBooksLimit {
			continue
		}
		booksByType[bt] = append(booksByType[bt], id)
	}

	DefaultHeaders(w)

	updated := "recently"
	if scu, ok := rdx.GetFirstVal(data.SyncCompletedProperty, data.SyncCompletedProperty); ok {
		if scui, err := strconv.ParseInt(scu, 10, 64); err == nil {
			updated = time.Unix(scui, 0).Format(time.RFC1123)
		}
	}

	if err := app.RenderGroup(
		stencil_app.NavLatestBooks,
		stencil_app.BookTypeOrder,
		booksByType,
		stencil_app.BookTypeTitles,
		bookTypeTotals,
		updated,
		r.URL,
		rdx,
		w); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}
}
