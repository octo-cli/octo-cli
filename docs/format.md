
<!--- start function list --->

## General Functions

### coalesce

Returns the first non-empty argument.

`{{coalesce 0 1 2}}` outputs `1`

### default

Returns a default value if the given value is empty, zero or false.

`{{ default "foo" "bar"}}` outputs `bar`

`{{ default "foo" ""}}` outputs `foo` because `""` is an empty string.

`{{ default "foo" 12}}` outputs `12`

`{{ default "foo" 0}}` outputs `foo` because `0` is a zero value string.

`{{ list 1 2 3 | default "foo" }}` outputs `[1 2 3]`

`{{ list | default "foo" }}` outputs `foo` because `list` with no arguments returns an empty array.

`{{ default "foo" (obj "foo" "bar")}}` outputs `map[foo: bar]`

`{{ default "foo" (obj) }}` outputs `foo` because `obj` with no arguments returns an empty map.

`{{ default "foo" nil }}` outputs `foo` because `nil` is treated as empty

### empty

Returns `true` if its argument is empty.

`{{ empty "" }}` outputs `true`

`{{ empty "foo" }}` outputs `false`

### ternary

Takes two values and a test value. Returns the first value if the test value is true 
or the second value if the test value is false.

`{{ ternary "foo" "bar" true }}` outputs `foo`

`{{ ternary "foo" "bar" false }}` outputs `bar`

### toCsv

Renders its arguments as a line in csv format (including a trailing `\n` which is omitted from the examples below)

`{{ toCsv "foo", "bar" }}` outputs `foo,bar`

If its only argument is a list, it expands the list and treats that as its argument list.

`{{ list "foo", "bar" | toCsv }}` outputs `foo,bar`

### toJson

Returns a json representation of its argument.

`{{ obj "foo" "bar" | toJson }}` outputs `{"foo": "bar"}`

### toPrettyJson

Like `toJson` but adds indentation and linebreaks to make the output pretty.

### toRawJson

Like `toJson` but doesn't escape `&`, `<` and `>` in quoted strings.

### toTsv

Renders its arguments as a line in tsv format (including a trailing `\n` which is omitted from the examples below)

`{{ toTsv "foo", "bar" }}` outputs `foo bar` (the whitespace between `foo` and `bar` is a tab).

If its only argument is a list, it expands the list and treats that as its argument list.

`{{ list "foo", "bar" | toTsv }}` outputs `foo bar` (the whitespace between `foo` and `bar` is a tab).

### toYaml

Returns a yaml representation of its argument.

`{{ obj "foo" "bar" | toYaml }}` outputs: `foo: bar`

## String Functions

### cat

The `cat` function concatenates multiple strings together into one, separating
them with spaces:

`{{cat "hello" "beautiful" "world"}}` outputs `hello beautiful world`

### contains

Test to see if one string is contained inside of another:

`{{contains "cat" "catch"}}` outputs `true`.

`{{contains "dog" "catch"}}` outputs `false`.


### fromBase64

Decodes a base64 string:

`{{fromBase64 "aGVsbG8gd29ybGQ="}}` outputs `hello world`

### lower

Convert the entire string to lowercase:

`{{lower "HELLO"}}` outputs `hello`

### newLine

Appends a `\n` to the end of a string:

`{{newLine "hello world"}}` outputs `hello world\n`

### replace

Perform simple string replacement.

It takes three arguments:

- string to replace
- string to replace with
- source string

`{{"I Am Henry VIII" | replace " " "-"}}` outputs `I-Am-Henry-VIII`

### split

Splits a string into an array of strings:

`{{split "$" "foo$bar$baz"}}` outputs `[foo bar baz]`

### substr

Get a substring from a string. It takes three parameters:

- start (int)
- end (int)
- string (string)

`{{substr 0 5 "hello world"}}` outputs `hello`

### toString

Converts an object to a string.

### trim

The `trim` function removes space from either side of a string:

`{{trim "   hello    "}}` outputs `hello`

### trimAll

Remove given characters from the front or back of a string:

`{{trimAll "$" "$5.00"}}` outputs `5.00`.

### trimPrefix

Trim just the prefix from a string:

`{{trimPrefix "-" "-hello"}}` outputs `hello`

### trimSuffix

Trim just the suffix from a string:

`{{trimSuffix "-" "hello-"}}` outputs `hello`

### trunc

Truncate a string (and add no suffix)

`{{trunc 5 "hello world"}}` outputs `hello`.

`{{trunc -5 "hello world"}}` outputs `world`.

### upper

Convert the entire string to uppercase:

`{{upper "hello"}}` outputs `HELLO`

## Array Functions

### compact

Returns a copy of an array with empty values removed. 

`{{list 1 "a" "foo" "" | compact}}` outputs `[1 a foo]`

### first

Returns the first item in an array.

`{{first list(1 2 3)}}` outputs `1`

### has

Tests whether an array has a particular element.

`{{ list 1 2 3 | has 1 }}` outputs `true`

`{{ list 1 2 3 | has 4 }}` outputs `false`

### join

Joins an array into a single string with the given separator.

`{{ list "hello" "world" | join "_" }}` outputs `hello_world`

If it has multiple value arguments, it will treat them as if they were values in an array.

`{{ join "_" "hello" "world" }}` outputs `hello_world`

### last

Returns the last item in an array.

`{{last list(1 2 3)}}` outputs `3`. 

### list

Turn a sequence of items into an array slice:

`{{$myList := list 1 2 3 4 5}}` outputs `[1 2 3 4 5]`.

### slice

Returns a slice of an array the first argument after the array is the index 
of the first item to return and the second argument is the index of the last
item to return. If the last argument is omitted, the returned slice will contain
all values after the first argument. 

`{{slice (list 1 2 3 4 5) 3}}` outputs `[4 5]`.

`{{slice (list 1 2 3 4 5) 1 3}}` outputs `[2 3]`.

`{{slice (list 1 2 3 4 5) 0 3}}` outputs `[1 2 3]`.

### sortAlpha

Sorts a list of strings into alphabetical (lexicographical) order.

`{{list "z" "x" "y" | sortAlpha}}` outputs `[x y z]`

### toStrings

Convert a list, slice, or array to a list of strings.

### uniq

Returns a copy of an array with all the duplicates removed:

`{{list 1 1 1 2 | uniq}}` outputs `[1 2]`

## Object Functions

"Object" refers to a json object. Sprig refers to these as dictionary functions, but 
we are dealing in json here, so this uses the json terminology.

In the examples below, the variable `$myObject` will contain an object unmarshalled from <br/> 
`{ "foo": "bar", "baz": "qux", "hey": "bear" }` and `$myOtherObject` is from <br/> 
`{ "one": 1, "two": 2, "hey": "bear" }`.

### get

Returns the value from one key in an object.

`{{get $myObject "foo"}}` outputs `bar`

### hasKey

Returns true or false depending on whether the object contains the given key.

`{{hasKey $myObject "foo}}` outputs `true`

`{{hasKey $myObject "missing key}}` outputs `false`

### keys

Returns a list of all the keys in one or more objects.
Since object keys are _unordered_, the keys will not be in a predictable order.
They can be sorted with `sortAlpha`.

`{{keys $myObject | sortAlpha}}` outputs `[baz foo hey]`

Multiple objects, will have their keys concatenated. Use the `uniq`
function along with `sortAlpha` to get a unique, sorted list of keys.

`{{keys $myObject $myObject $myOtherObject | uniq | sortAlpha }}` outputs `[baz foo hey one two]`

### obj

Creates an object based on key/value pairs. 

`{{obj "foo" "bar" "one" 1 "bool" false}}` outputs `map[bool:false foo:bar one:1]`

### omit

Returns a copy of the object with the given keys omitted.

`{{omit $myObject "foo" "baz"}}` outputs `{hey: bear}`

### pick

Selects just the given keys out of an object, creating a new object.

`{{pick $myObject "foo" "baz"}}` outputs `{baz: qux, foo: bar}`

### pluck

Given multiple objects and one key, pluck returns a list of values from all objects
that contain the key. If an object does not contain the key, that object will not have
an item in the list and the list will be shorter than the number of objects in the `pluck`
call.

`{{pluck "hey" $myObject $myOtherObject}}` outputs `[bear bear]`

`{{pluck "one" $myObject $myOtherObject}}` outputs `[1]`

When `pluck`'s only argument is a list of objects, it treats the call as if the objects in 
that list were its arguments.

`{{list $myObject $myOtherObject | pluck "hey"}}` outputs `[bear bear]`

<!--- end function list --->
