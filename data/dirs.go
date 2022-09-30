package data

import (
	"path/filepath"
	"strconv"
)

const (
	relReduxDir         = "_redux"
	relMyBookFreshDir   = "my_books_fresh"
	relMyBookDetailsDir = "my_books_details"
	relDownloadsDir     = "downloads"
	relCoversDir        = "covers"

	relCookiesFilename = "cookies.txt"

	CoverExt = ".jpg"
)

var rootDir = ""

func ChRoot(d string) {
	rootDir = d
}

func Pwd() string {
	return rootDir
}

func AbsMyBooksFreshDir() string {
	return filepath.Join(rootDir, relMyBookFreshDir)
}

func AbsMyBooksDetailsDir() string {
	return filepath.Join(rootDir, relMyBookDetailsDir)
}

func AbsDownloadsDir() string {
	return filepath.Join(rootDir, relDownloadsDir)
}

func AbsDownloadPath(id int64, file string) string {
	return filepath.Join(AbsDownloadsDir(), strconv.FormatInt(id, 10), file)
}

func AbsReduxDir() string {
	return filepath.Join(rootDir, relReduxDir)
}

func AbsCoverPath(id int64) string {
	return filepath.Join(AbsCoverDir(), strconv.FormatInt(id, 10)+CoverExt)
}

func AbsCoverDir() string {
	return filepath.Join(rootDir, relCoversDir)
}

func AbsCookiesFilename() string {
	return filepath.Join(rootDir, relCookiesFilename)
}
