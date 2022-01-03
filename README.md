# Translationator!

This app takes an English sentence (or a paragraph or whatever you're feeling), and proceeds to translate it from 
English into a bunch of other languages, and then back to English.  The goal is to convert a normal sentence into 
something totally crazy.  

## Prerequisites
- Install Golang on your computer.  This comes in various shapes and forms based on your OS.
- Register with Google Cloud Console (https://console.cloud.google.com/) to get yourself a translation API key.

## Installation
To install (run from translationator root directory):

```shell
go install translationator
```

## Running the dang thing

To run the app, just run:
```shell
translationator "google-translate-api-key" "Target phrase" "Optional # of iterations"
```

Sample run:
```shell
> translationator "lolfakeapikeyblahblahblah" "Hello, I am writing words in English" 10
I write greetings in English
```

## Developer Notes

TODO:
- Perhaps support JSON file configurations.  Maybe integrate with viper for this?
- The code needs comments lol.  Haven't done this because this first iteration is undoubtedly UGLY and needs refactors.
