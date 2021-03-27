# hexdump
Simple hexdump written in Go

Use the canonical hex+ASCII display, which
- displays the input offset in hexadecimal, followed by
- sixteen space-separated, two column, hexadecimal bytes, followed by
- the same sixteen bytes in ASCII if printable or '.' if not
