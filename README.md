# ankiconnect

A lightweight Golang interface for ankiconnect. This library depends on [Anki](https://apps.ankiweb.net/index.html) and the [ankiconnect](https://ankiweb.net/shared/info/2055492159) addon.

[![Build](https://github.com/atselvan/ankiconnect/actions/workflows/build.yaml/badge.svg)](https://github.com/atselvan/ankiconnect/actions/workflows/build.yaml)
[![Release](https://github.com/atselvan/ankiconnect/actions/workflows/release.yaml/badge.svg)](https://github.com/atselvan/ankiconnect/actions/workflows/release.yaml)
[![reference](https://img.shields.io/badge/godoc-docs-blue.svg?label=&logo=go)](https://godoc.org/github.com/atselvan/ankiconnect)
[![Go Report Card](https://goreportcard.com/badge/github.com/atselvan/ankiconnect)](https://goreportcard.com/report/github.com/atselvan/ankiconnect)
[![codecov](https://codecov.io/gh/atselvan/ankiconnect/branch/main/graph/badge.svg?token=05D9AMNEUF)](https://codecov.io/gh/atselvan/ankiconnect)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=atselvan_ankiconnect&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=atselvan_ankiconnect)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fatselvan%2Fankiconnect.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fatselvan%2Fankiconnect?ref=badge_shield)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

* Get Anki deck names.
* Create a new deck in Anki.
* Delete a existing deck from Anki.
* Add a new card to an existing deck in Anki.

## Installation

```go
go get github.com/atselvan/ankiconnect
```

## Usage

> [Anki](https://apps.ankiweb.net/index.html) should be running in the background with the [ankiconnect](https://ankiweb.net/shared/info/2055492159) addon installed for below examples to work.

### Ping

```go
client := ankiconnect.NewClient()
restErr := client.Ping()
if restErr != nil {
	log.Fatal(restErr)
}
```

### Get Decks

```go
client := ankiconnect.NewClient()
decks, restErr := client.Decks.GetAll()
if restErr != nil {
	log.Fatal(restErr)
}
fmt.Println(decks)
```

### Create Deck

```go
client := ankiconnect.NewClient()
restErr := client.Decks.Create("New Deck")
if restErr != nil {
	log.Fatal(restErr)
}
```

### Delete Deck

```go
client := ankiconnect.NewClient()
restErr := client.Decks.Delete("New Deck")
if restErr != nil {
	log.Fatal(restErr)
}
```

### Create Note

```go
client := ankiconnect.NewClient()

note := ankiconnect.Note{
	DeckName: "New Deck",
	ModelName: "Basic-a39a1",
	Fields: ankiconnect.Fields{
		Front: "Front data",
		Back: "Back data",
	},
}

restErr := client.Notes.Add(note)
if restErr != nil {
	log.Fatal(restErr)
}
```

### Get Notes

```go
client := ankiconnect.NewClient()

// Get the Note Ids of cards due today
nodeIds, restErr := client.Notes.Get("prop:due=0")
if restErr != nil {
	log.Fatal(restErr)
}

// Get the Note data of cards due today
notes, restErr := client.Notes.Get("prop:due=0")
if restErr != nil {
	log.Fatal(restErr)
}

```

### Get Cards

```go
client := ankiconnect.NewClient()

// Get the Card Ids of cards due today
nodeIds, restErr := client.Cards.Get("prop:due=0")
if restErr != nil {
	log.Fatal(restErr)
}

// Get the Card data of cards due today
notes, restErr := client.Cards.Get("prop:due=0")
if restErr != nil {
	log.Fatal(restErr)
}

```

### Sync local data to Anki Cloud
```go
client := ankiconnect.NewClient()

restErr := client.Sync.Trigger()
if restErr != nil {
	log.Fatal(restErr)
}
```
