package url

import (
	"net/url"
	"path"
	"strings"
)

func URLJoin(base string, parts ...string) (string, error) {
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	joinedPath := path.Join(parts...)
	argURL, err := url.Parse(joinedPath)
	if err != nil {
		return "", err
	}
	joinedURL := baseURL.ResolveReference(argURL).String()
	if !baseURL.IsAbs() && !strings.HasPrefix(base, "/") {
		return joinedURL[1:], nil
	}
	return joinedURL, nil
}
