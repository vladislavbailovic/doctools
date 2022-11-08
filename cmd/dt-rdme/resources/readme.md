# {{ .Name }}{{ if .Description }} - {{ .Description }}{{ end }}

{{ if .Sections }}
## Table of Contents
{{ range $title, $commands := .Sections }}
	- [{{ $title }}](#{{ $title | slugify }}){{ end }}

{{ end }}
## Quick Start

{{ range $title, $commands := .Sections }}
### {{ $title }}

```sh{{ range $commands }}
$ {{ . }}{{ end }}
```

{{ end }}
