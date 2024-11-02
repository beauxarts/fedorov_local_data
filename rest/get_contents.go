package rest

import (
	"github.com/beauxarts/fedorov/rest/compton_pages"
	"github.com/boggydigital/nod"
	"net/http"
)

func GetContents(w http.ResponseWriter, r *http.Request) {

	// GET /contents?id

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, nod.ErrorStr("missing book id"), http.StatusInternalServerError)
		return
	}

	var err error
	if rdx, err = rdx.RefreshReader(); err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		return
	}

	//raReader, err := data.NewArtsReader(litres_integration.ArtsTypeReviews)
	//if err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//artsReviews, err := raReader.ArtsReviews(id)
	//if err != nil {
	//	http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	//	return
	//}

	if p := compton_pages.Contents(); p != nil {
		if err = p.WriteResponse(w); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
		}
	}
}
