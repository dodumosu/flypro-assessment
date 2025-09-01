package config

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

/*
	this has been adapted from https://github.com/smilingthrone13/url-builder
*/

type credentials struct {
	user     string
	password string
}

type Builder struct {
	Scheme      string
	Host        string
	Port        int
	Credentials *credentials
	Path        []string
	Query       map[string][]string
	Anchor      string
}

func NewURLBuilder() *Builder {
	return &Builder{
		Path:  make([]string, 0),
		Query: make(map[string][]string),
	}
}

func (b *Builder) WithScheme(scheme string) *Builder {
	b.Scheme = scheme
	return b
}

func (b *Builder) WithHost(host string) *Builder {
	b.Host = strings.TrimSuffix(host, "/")
	return b
}

func (b *Builder) WithPort(port int) *Builder {
	b.Port = port
	return b
}

func (b *Builder) WithCredentials(user string, password string) *Builder {
	b.Credentials = &credentials{
		user:     user,
		password: password,
	}
	return b
}

func (b *Builder) WithPath(elements ...string) *Builder {
	b.Path = append(b.Path, elements...)
	return b
}

func (b *Builder) WithQuery(key string, values ...string) *Builder {
	b.Query[key] = append(b.Query[key], values...)
	return b
}

func (b *Builder) WithAnchor(anchor string) *Builder {
	b.Anchor = strings.Trim(anchor, "#/")
	return b
}

func (b *Builder) Build() (string, error) {
	if b.Host == "" {
		return "", fmt.Errorf("host is required")
	}

	// check given host
	// todo: can't detect if given ipv6 contains port, so result string might be broken.
	if strings.Contains(b.Host, "/") || // assume host contains scheme
		strings.Count(b.Host, ":") == 1 { // assume host contains port (valid ipv6 have at least 2 colons)
		return "", fmt.Errorf("host contains forbidden symbols")
	}

	rawBaseUrl := fmt.Sprintf("%s://%s", b.Scheme, b.Host)

	if b.Port > 0 {
		if b.Port > 65535 {
			return "", fmt.Errorf("port must be in range [1, 65535]")
		}
		rawBaseUrl = fmt.Sprintf("%s:%d", rawBaseUrl, b.Port)
	}

	u, err := url.Parse(rawBaseUrl)
	if err != nil {
		return "", err
	}

	if b.Credentials != nil {
		if b.Credentials.user == "" && b.Credentials.password == "" {
			return "", fmt.Errorf("empty credentials")
		}

		if b.Credentials.password == "" {
			u.User = url.User(b.Credentials.user)
		} else {
			u.User = url.UserPassword(b.Credentials.user, b.Credentials.password)
		}

	}

	if len(b.Path) > 0 {
		u = u.JoinPath(b.Path...)
	}

	for k, v := range b.Query {
		if k == "" {
			return "", fmt.Errorf("query key is empty")
		}
		if i := slices.Index(v, ""); i != -1 {
			return "", fmt.Errorf("query query value for key %s", k)
		}
	}

	u.RawQuery = url.Values(b.Query).Encode()

	if b.Anchor != "" {
		u.Fragment = b.Anchor
	}

	return u.String(), nil
}
