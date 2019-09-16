SES SMTP Password converter
===========================

Simple CLI to convert IAM credentials to SMTP password that can be used to
authenticate againsts AWS SES.

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
