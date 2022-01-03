# Translationator!

This app takes an English sentence (or a paragraph or whatever you're feeling), and proceeds to translate it from 
English into a bunch of other languages, and then back to English.  The goal is to convert a normal sentence into 
something totally crazy.  

## Prerequisites
- Install Golang on your computer.  This comes in various shapes and forms based on your OS.
- Register with Google Cloud Console (https://console.cloud.google.com/) to get yourself a translation API key or 
  service account JSON file (the JSON file is preferred for ease of use).

## Installation
To install (run from translationator root directory):

```shell
go install translationator
```

## Running the dang thing

To run the app with a JSON file in the default directory (~/translationator/translation-credentials.json), just run:
```shell
translationator -t "Target phrase"
```

To run the app with an API key via Google Cloud, just run:
```shell
translationator -a myApiKeyFromGoogleCloud -t "Target phrase"
```

Sample run:
```shell
> translationator -t "Hello, I am writing words in English"
I write greetings in English
```
