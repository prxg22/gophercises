package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(res, req, dest, http.StatusFound)
		} else {
			fallback.ServeHTTP(res, req)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYAML(yml)

	if err != nil {
		return nil, err
	}

	m := buildMap(pathUrls)

	return MapHandler(m, fallback), nil
}

func parseYAML(yml []byte) ([]pathUrl, error) {
	var data []pathUrl
	err := yaml.Unmarshal(yml, &data)

	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %v", err)
	}

	return data, nil
}

func buildMap(data []pathUrl) map[string]string {
	pathUrls := make(map[string]string)

	for _, pu := range data {
		pathUrls[pu.Path] = pu.Url
	}

	return pathUrls
}
