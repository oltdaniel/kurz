# Documentation

This is a small documentation on how `kurz` works in the background. The api
documentation can be found [here in `api.md`](https://github.com/oltdaniel/kurz/blob/master/doc/api.md).


## Generation

We use `abcdefghijklmnopqrstuvwxyz0123456789#._-:?!` as our charset. It is used with [`nanoid`](https://github.com/matoous/go-nanoid) to generate an 6 character long random identifier. The total number of possible urls is `6^55 = 6,285,195,213,566,005,335,561,053,533,150,026,217,291,776`, or in words `a lot`. Since it is also possible to use your own slug, the amount of urls that can be generated is ∞ _(infinity)_.
