package handler

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/prxg22/gophercises/cyoa/story"
	"github.com/prxg22/gophercises/utils/logger"
)

type StoryHandler struct {
	*story.Story
}

func (sh *StoryHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path == "/" {
		logger.Info("redirecting to /intro")
		http.Redirect(res, req, "/intro", http.StatusFound)
		return
	}

	paths := strings.Split(path, "/")
	arcKey := paths[1]
	arc := sh.Story.GetArc(arcKey)

	if arcKey == "" || arc == nil {
		logger.Warn("Arc \"%v\" not found!", arcKey)
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("arc not found"))
		return
	}

	tmpl, tmplErr := template.New("template.html").ParseFiles("template.html")

	if tmplErr != nil {
		logger.Error(tmplErr)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(tmplErr.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "text/html")
	logger.Info("%v", tmpl)
	err := tmpl.Execute(res, *arc)
	if err != nil {
		logger.Error(err)
	}

}
