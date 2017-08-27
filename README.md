# slack-typer [![Travis Ci Status](https://travis-ci.org/joshheinrichs/slack-typer.svg?branch=master)](https://travis-ci.org/joshheinrichs/slack-typer)
Bug your coworkers with constant typing notifications.

## Installing 

Either grab a build on the [releases page](https://github.com/joshheinrichs/slack-typer/releases) or run:

```bash
$ go get -u github.com/joshheinrichs/slack-typer
$ go install github.com/joshheinrichs/slack-typer
```

## Usage

```bash
$ slack-typer -slack-token SLACK_TOKEN -user USER [-h]
```

Slack tokens can be generated at https://api.slack.com/custom-integrations/legacy-tokens
