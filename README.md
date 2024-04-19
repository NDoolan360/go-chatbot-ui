# Go Chatbot UI

## Purpose

This project was completed as a part of a firebreak week I participated in at Kaluza. 🔥

The purpose of this project was to integrate it with an AI model trained on BigQuery schemas to be able to ask it questions about database tables and ultimately improve data visibility across the engineering teams.
An alternate suggestion was to have it integrate with a model trained on post-mortems from incidents and have the chatbot act as a way to assist with diagnosing and triaging incidents.

Sadly, the AI models were not able to be completed within the constraints of the firebreak so all it does at the moment is return the question in reverse. ↩️

## Features

- Minimal Material based UI
- Websocket based conversation with server using htmx websocket extension
- Messages rendered with markdown

## Usage

### Development

1. Install Go. [(mac-OS)][brew-install-go] - [(other)][go-website]

2. Run the entry program.

```shell
go run go-chatbot-ui.go
// Server listening on port:	http://localhost:3000
```

## Extertnal Components

- [Htmx](/client/public/scripts/htmx.min.js) - [Big Sky Software][htmx]
- [Htmx Websockets](/client/public/scripts/htmx.ws.min.js) - [Big Sky Software][htmx-ws]
- [Hyperscript](/client/public/scripts/hyperscript.min.js) - [Big Sky Software][hyperscript]
- [Beer CSS](/client/public/scripts/beercss.min.js) - [Beer CSS][beer-css]
- [md-block](/client/public/scripts/md-block.min.js) - [Lea Verou][beer-css]

[brew-install-go]: https://formulae.brew.sh/formula/go
[go-website]: https://go.dev/doc/install
[htmx]: https://github.com/bigskysoftware/htmx
[htmx-ws]: https://htmx.org/extensions/web-sockets/
[hyperscript]: https://github.com/bigskysoftware/_hyperscript
[beer-css]: https://github.com/beercss/beercss
[md-block]: https://md-block.verou.me/
