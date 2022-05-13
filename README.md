# googletranslate

## About

Translate source language to another language
<br /><br /><br />
## Defaults (without passing values)

```
go run main.go 
Options:
  -s string
        Source lanuguage[en (default "en")
  -st string
        Text to translate
  -t string
        Target language[fr] (default "fr")
```
<br /><br /><br />
## Example (translate English to Spanish)
```
go run main.go -s en -st hello -t es
Hola
```