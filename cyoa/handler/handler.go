package handler

import (
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
		res.Write([]byte("not found"))
		return
	}

	json, marshalArcErr := arc.Marshal()

	if marshalArcErr != nil {
		logger.Error(marshalArcErr)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "application/json")
	res.Write(json)
}
