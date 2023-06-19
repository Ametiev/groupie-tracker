package web

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func Start(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	artists, err := Get_artists()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	files := []string{
		"./ui/html/home.html",
		"./ui/html/base.html",
	}
	parsed_home_page, err := template.ParseFiles(files...)
	if parsed_home_page == nil || err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = parsed_home_page.Execute(w, artists)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	xurl := strings.Split(url, "?id=")
	id, _ := strconv.Atoi(xurl[1])
	if url != "/info/?id="+xurl[1] {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if id > 52 || id < 1 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	files := []string{
		"./ui/html/info.html",
		"./ui/html/base.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = UnmarshallRelations()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, Artists[id-1])
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, status int) {
	files := []string{
		"./ui/html/error.html",
		"./ui/html/base.html",
	}
	parsed_err_page, err := template.ParseFiles(files...)
	if parsed_err_page == nil || err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var Result ErrorStr
	Result.Status = status
	Result.Message = http.StatusText(status)
	err = parsed_err_page.Execute(w, Result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
