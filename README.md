# [Zone01](https://www.zone01kisumu.ke/) - Go-Reloaded Tool

[Raymond Ogwel](https://learn.zone01kisumu.ke/git/rogwel)'s Go- Reloaded project

## Getting Started

### Summary of what the tool does

In this project we will be making a simple text completion/editing/auto-correction tool.
The project is a tool that does correction on a text file
and displays the correct file content into another file, therefore the project is a simple auto-correction tool.

### Introduction

- This program has been written in [Golang](https://go.dev/).
- This project respects the basic rules of [good practices](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/good-practices/README.md).
- The project has a test file accompanying it as well for testing.

### What the tool does in detail

The tool receives as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output). Next is a list of possible modifications  that the program executes

- Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")

- Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")

- Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")

- Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

- Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
    - For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

- Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
    - Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".

- The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")
    - If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")

- Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!")

### To Start You Off

- Create an input file, say you create sample and give it an extension of .txt.

``
$touch sample.txt
``

- Populate the sample file with the raw text and save it ie.

``
punctuation (cap) is the easiest part of grammar in this language ,err ... english (cap) .
``

- Continue without creating the second file(the output file)
- Use the commands below to run the tool to modify your provided text:

``
    $go run . sample.txt result.txt
``

- Assuming that your output file is to be named result
- This saves the modified text onto result.txt
- Use the command below to view your modified text on the command line :

``
    $cat result.txt
``

- Your results will be modified as in this provided example:

``
Punctuation is the easiest part of grammar in this language, err... English.
``

### Testing

The project comes with its own test file, which can be tested at the command line using the command :

``
$go test
``

The output is :

``
$ok
``

This means that your program has been successfully tested!
