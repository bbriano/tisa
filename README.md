Tisa
====

Tisa replace tab characters (ab)used for alignment with spaces.
It only replace tab characters after the first non-tab character.
Or in other words, Tisa preserves tab characters used for indentation.

Synopsis
--------

	$ tisa [ nspace ] <file

Examples
--------

Vim:

	:execute '%!tisa ' . &ts

Acme:

	Edit ,|tisa 4

Install
-------

	$ go install github.com/bbriano/tisa

About
-----

	T|ab
	I|ndent
	S|pace
	A|lign
