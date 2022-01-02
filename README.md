# Translationator!

This app takes an English sentence (or a paragraph or whatever you're feeling), and proceeds to translate it from 
English into a bunch of other languages, and then back to English.  The goal is to convert a normal sentence into 
something totally crazy.  

To install (run from translationator root directory):

```shell
go install translationator
```

To run the app, just run:
```shell
translationator "google-translate-api-key" "Target phrase" "Optional # of iterations"
# Ex:
translationator "lolfakeapikeyblahblahblah" "Hello, I am writing words in English" 20
```
