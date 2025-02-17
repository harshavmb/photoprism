package config

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/photoprism/photoprism/pkg/fs"
)

// DetachServer checks if server should detach from console (daemon mode).
func (c *Config) DetachServer() bool {
	return c.options.DetachServer
}

// HttpHost returns the built-in HTTP server host name or IP address (empty for all interfaces).
func (c *Config) HttpHost() string {
	if c.options.HttpHost == "" {
		return "0.0.0.0"
	}

	return c.options.HttpHost
}

// HttpPort returns the built-in HTTP server port.
func (c *Config) HttpPort() int {
	if c.options.HttpPort == 0 {
		return 2342
	}

	return c.options.HttpPort
}

// HttpMode returns the server mode.
func (c *Config) HttpMode() string {
	if c.options.HttpMode == "" {
		if c.Debug() {
			return "debug"
		}

		return "release"
	}

	return c.options.HttpMode
}

// HttpCompression returns the http compression method (none or gzip).
func (c *Config) HttpCompression() string {
	return strings.ToLower(strings.TrimSpace(c.options.HttpCompression))
}

// TemplatesPath returns the server templates path.
func (c *Config) TemplatesPath() string {
	return filepath.Join(c.AssetsPath(), "templates")
}

// CustomTemplatesPath returns the path to custom templates.
func (c *Config) CustomTemplatesPath() string {
	if p := c.CustomAssetsPath(); p != "" {
		return filepath.Join(p, "templates")
	}

	return ""
}

// TemplateFiles returns the file paths of all templates found.
func (c *Config) TemplateFiles() []string {
	results := make([]string, 0, 32)

	tmplPaths := []string{c.TemplatesPath(), c.CustomTemplatesPath()}

	for _, p := range tmplPaths {
		matches, err := filepath.Glob(regexp.QuoteMeta(p) + "/[A-Za-z0-9]*.*")

		if err != nil {
			continue
		}

		for _, tmplName := range matches {
			results = append(results, tmplName)
		}
	}

	return results
}

// TemplateExists checks if a template with the given name exists (e.g. index.tmpl).
func (c *Config) TemplateExists(name string) bool {
	if found := fs.FileExists(filepath.Join(c.TemplatesPath(), name)); found {
		return true
	} else if p := c.CustomTemplatesPath(); p != "" {
		return fs.FileExists(filepath.Join(p, name))
	} else {
		return false
	}
}

// TemplateName returns the name of the default template (e.g. index.tmpl).
func (c *Config) TemplateName() string {
	if s := c.Settings(); s != nil {
		if c.TemplateExists(s.Templates.Default) {
			return s.Templates.Default
		}
	}

	return "index.tmpl"
}

// StaticPath returns the static assets' path.
func (c *Config) StaticPath() string {
	return filepath.Join(c.AssetsPath(), "static")
}

// StaticFile returns the path to a static file.
func (c *Config) StaticFile(fileName string) string {
	return filepath.Join(c.AssetsPath(), "static", fileName)
}

// BuildPath returns the static build path.
func (c *Config) BuildPath() string {
	return filepath.Join(c.StaticPath(), "build")
}

// ImgPath returns the path to static image files.
func (c *Config) ImgPath() string {
	return filepath.Join(c.StaticPath(), "img")
}
