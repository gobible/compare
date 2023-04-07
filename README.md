# Gobible Compare

A simple tool to compare two bible versions, creating a new version with (++additons++) and (~~subtractions~~) in the output.

## Usage

```bash
go run . file1.json file2.json
```

Or, if you build the binary:

```bash
gobible-compare file1.json file2.json
```

## Example

Here is John 1:3 from the KJV to the WEB

```
KJV: All things were made by him; and without him was not any thing made that was made.
WEB: All things were made through him. Without him was not anything made that has been made.
```

Example JSON output in the new generated version:

```json
{
  "number": 3,
  "text": "All things were made (~~by him; and w~~)(++through him. W++)ithout him was not any(~~ ~~)thing made that (~~w~~)(++h++)as (++been ++)made."
}
```

