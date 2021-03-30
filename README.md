# textbundle-to-epub

Command line tool for converting textbundle to epub.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gonejack/textbundle-to-epub)
![Build](https://github.com/gonejack/textbundle-to-epub/actions/workflows/go.yml/badge.svg)
[![GitHub license](https://img.shields.io/github/license/gonejack/textbundle-to-epub.svg?color=blue)](LICENSE)

### Install
```shell
> go get github.com/gonejack/textbundle-to-epub
```

### Usage
```shell
> textbundle-to-epub *.textbundle
```
```
Usage:
  textbundle-to-epub [-o output] [--title title] [--cover cover] *.textbundle

Flags:
  -o, --output string   output filename (default "output.epub")
      --title string    epub title (default "TextBundles")
      --cover string    cover image
  -v, --verbose         verbose
  -h, --help            help for textbundle-to-epub
```
