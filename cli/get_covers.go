package cli

import (
	"errors"
	"github.com/beauxarts/fedorov/data"
	"github.com/beauxarts/litres_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/url"
	"strconv"
	"strings"
)

func GetCoversHandler(u *url.URL) error {
	ids := strings.Split(u.Query().Get("id"), ",")
	return GetCovers(ids)
}

func GetCovers(ids []string) error {

	gca := nod.NewProgress("fetching covers...")
	defer gca.End()

	rxa, err := kvas.ConnectReduxAssets(data.AbsReduxDir(), nil, data.MyBooksIdsProperty)
	if err != nil {
		return gca.EndWithError(err)
	}

	if len(ids) == 0 {
		var ok bool
		ids, ok = rxa.GetAllUnchangedValues(data.MyBooksIdsProperty, data.MyBooksIdsProperty)
		if !ok {
			err = errors.New("no my books found")
			return gca.EndWithError(err)
		}
	}

	gca.TotalInt(len(ids))

	dc := dolo.DefaultClient

	for _, id := range ids {

		idn, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			gca.Error(err)
			gca.Increment()
			continue
		}

		cu := litres_integration.CoverUrl(idn)

		if err := dc.Download(cu, nil, data.AbsCoverDir()); err != nil {
			gca.Error(err)
			gca.Increment()
			continue
		}

		gca.Increment()
	}

	gca.EndWithResult("done")

	return nil
}
