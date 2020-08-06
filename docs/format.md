By default, octo-cli outputs formatted json results with line-breaks and indenting. That
may be all you need. If you are so inclined, you can pipe the output to something like `jq` and 
ignore the rest of this page. That's a perfectly valid approach, especially if you are already
familiar with `jq`.

## Flags

There are three flags that control the output.

### `--format <template>`

The `--format` flag allows you to format output using [go templates](https://golang.org/pkg/text/template/). For 
example, you can output just the title of an issue with `--format '{{.title}}'`. To see the login of the issue submitter 
use `--format '{{.user.login}}'`

There are quite a few template functions available to use. See them in [Template Functions](template-functions.md)

### `--output-each <path>`

Instead of outputting the response body as one entity, iterate over an array found at `<path>`.

If you are making a request that returns a list, you may want to use `--output-each .`.

On a command that returns a pull request, if you are only interested in its labels, you may
want to use `--output-each .labels`

`--output-each` is primarily useful in conjunction with `--format`.

### `--raw-output`

Don't do any formatting. Just output the response body exactly as-is. When `--raw-output` is set,
`--format` and `--output-each` are ignored.

## Examples

- List numbers for open issues

```shell
$ octo issues list-for-repo --repo golang/go --per_page 3 \
    --output-each . \
    --format '{{.number | newLine}}'

40576
40575
40574
```

`--output-each .` causes it to iterate over each item in the result,
`.number` outputs the issue number and `| newLine` appends "\n".

- Output tsv

```shell
octo pulls get --repo octo-cli/octo-cli --pull_number 5 \
    --format '{{ toTsv .user.login .title }}'

WillAbides	Add MIT license
```

You can use this in bash to set multiple environment variables from one response:

```shell
$ read -r PR_AUTHOR PR_TITLE <<< "$(octo pulls get \
    --repo octo-cli/octo-cli --pull_number 5 \
    --format '{{ toTsv .user.login .title }}')"
$ echo $PR_AUTHOR
WillAbides
$ echo $PR_TITLE
Add MIT license
```

Use `toTsv` with `column -ts $'\t'` to output a human-readable table:

```shell script
$ octo issues list-for-repo --repo golang/go --direction asc --per_page 5  \
    --output-each . \
    --format '{{toTsv .number .title .user.login}}' \
    | column -ts $'\t'

101  doc: manual pages for Go tools                                                       gopherbot
377  proposal: spec: various changes to :=                                                agl
395  proposal: spec: use (*[4]int)(x) to convert slice x into array pointer               rogpeppe
449  cmd/vet: warn about unused struct or array, ignoring assignment to field or element  gopherbot
463  gccgo: compilation fails on Darwin                                                   gopherbot
```
