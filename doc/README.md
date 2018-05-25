# Documentation

This is a small documentation on how `kurz` works in the background. The api
documentation can be found [here in `api.md`](https://github.com/oltdaniel/kurz/blob/master/doc/api.md).

## Duration

Public links will be held in the database for 10 days, user created links for
infinitie.

## Generation

We use `abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789` as our charset.
It is used with [`nanoid`](https://github.com/matoous/go-nanoid) to generate an 6 
character long random identifier. The total number of possible urls is `6^58 = 
1,357,602,166,130,257,152,481,187,563,160,405,662,935,023,616`, or in words `a lot`.
Since it is also possible to use your own slug, the amount of urls that can be generated is ∞ _(infinity)_.
