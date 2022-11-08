# {{ .Name }}{{ if .Description }} - {{ .Description }}{{ end }}

{{ if .Sections }}
## Table of Contents
{{ range $title, $commands := .Sections }}
	- [{{ $title }}](#{{ $title | slugify }}){{ end }}

{{ end }}{{ range $title, $commands := .Sections }}
### {{ $title }}
{{ range $commands }}
```
$ {{ . }}
 
```
{{ end }}

{{ end }}
