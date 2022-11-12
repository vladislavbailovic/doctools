# {{ .Name }}{{ if .Description }} - {{ .Description }}{{ end }}

{{ if .Sections }}
## Table of Contents

- [Quick Start](#{{ "Quick Start" | slugify }}){{ range $title, $commands := .Sections }}
    - [{{ $title }}](#{{ $title | slugify }}){{ end }}

{{ end }}
## Quick Start

{{ range $title, $commands := .Sections }}
### {{ $title }}

```console{{ range $commands }}
$ {{ . }}{{ end }}
```

{{ end }}
