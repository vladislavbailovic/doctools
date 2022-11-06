ADR {{ .Number | printf "%03d" }}: {{ .Title }}
===============================================


Status
------

{{ range $idx, $s := .Status }}{{ if $idx }}, {{ end }}{{ $s.String }}{{ end }}


Context
-------

{{ .Context }}


Decision
--------

{{ .Decision }}


Consequences
------------

{{ .Consequences }}
