package config

import "html/template"

// holds application config
type AppConfig struct {
	UseCache bool
	// TODO: template.Templateとは？
	TemplateCache map[string]*template.Template
}