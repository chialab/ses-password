SES SMTP Password converter
===========================

Simple CLI to convert IAM credentials to SMTP password that can be used to
authenticate againsts AWS SES.

Installation
------------

If you have Go installed, run `go get github.com/chialab/ses-password` and you’re
ready to Go.

Otherwise, download a pre-build binary and copy it to a directory that’s in your
`$PATH`. Usually, `/usr/local/bin` is a good place to put the executable in.

Usage
-----

The CLI reads IAM Secret Access Key from stdin and outputs the SMTP password
to stdout. Although this is not user-friendly, it makes piping easier as no
formatting or parsing are required.

On macOS, you may also pipe from and to pasteboard. Assuming you already have
the IAM Secret Access Key in your clipboard, run:

```sh
pbpaste | ses-password | pbcopy
```

This will read the secret access key from your clipboard, pass it through the
executable, and put the SMTP password in your clipboard.

Other OS-es provide similar functionality with equivalent commands
(e.g. `xclip` on Ubuntu).

Contributing
------------

Any contribution is welcome. Issues, Pull Requests and any form of suggestion
is welcome.

Fork this repository and make commits, then open a Pull Request against `master`,
or contact us at <dev@chialab.io> to get in touch.
