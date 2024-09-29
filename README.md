# smtp-cli

> 📨 Send emails right from your terminal.

We use this tool to send an informational email whenever someone logs in to one of our servers.
Intended usage is via [pam_exec(8)](https://linux.die.net/man/8/pam_exec).

Emails are sent directly via SMTP. The config is read from a json file, see [Config](#config).

The message body is produced via go templates, see [Templates](#templates).

## Installation

Download an archive (`.tar.gz`) from the latest release and run the `install.sh` script.

## Usage

Add the following snippet to `/etc/pam.d/sshd`:

```
session    required     pam_exec.so /usr/local/bin/smtp-cli -config /usr/local/etc/smtp-cli/config.json -template /usr/local/etc/smtp-cli/template.gotmpl
```

## Config

```json
{
    "host":  "mail.example.org",
	"port":  465,
	"username":  "no-reply@example.org",
	"password":  "",
	"from":  "monitor <no-reply@example.org>",
	"to":  [
		"info <info@example.org>"
	],
	"subject": "New login to the server"
}
```

## Templates

The following variables are made available to the go template:

| Variable | Type |
| --- | --- |
| Host       | string |
| User       | string |
| RemoteUser | string |
| RemoteHost | string |
| Tty        | string |
| Timestamp  | time.Time |

## License

[MIT](LICENSE)
