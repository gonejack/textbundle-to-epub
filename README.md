# textbundle-to-epub

Command line tool for converting textbundles to epub.

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
