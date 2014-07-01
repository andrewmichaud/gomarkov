package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "rand"
)

const source_file_path string = "/Users/amichaud/Downloads/holmes.txt"

// We use these to tell when we've found the end of a sentence.
const punctuation = ".?!"

// We use this to mark words that may start a sentence in the dictionary.
const start_marker = "$"

// Using the file in source_file_path as a source, create a dictionary for our Markov generator.
func generate_dictionary() map[string][]string {


    // Dictionary we'll return.
    dictionary := make(map[string][]string)

    // Open the file.
    file, err := os.Open(source_file_path)

    // Error handling.
    if (err != nil) {
        panic(err)
    }

    // Make sure we close the file.
    defer file.Close()

    // Scanner to grab words with.
    scanner := bufio.NewScanner(bufio.NewReader(file))

    // Previous word storage.
    previous := ""

    // Read all words.
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {

        // Get words.
        word := scanner.Text()

        // Error check.
        if err != nil {
            panic(err)
        }

        // For first word in file, create entry for sentence-starting words.
        if previous == "" {
            dictionary["$"] = []string{word}

        // For other words that came after the end of a sentence, add to sentence-starting words.
        } else if strings.ContainsAny(previous, punctuation) {
            oldVal := dictionary[start_marker]
            dictionary[start_marker] = append(oldVal, word)

        // For any other words, look up the entry for the previous word and add to it.
        } else {
            oldVal := dictionary[previous]
            dictionary[previous] = append(oldVal, word)
        }
        previous = word
    }

    // Scanner error checking.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input:", err)
    }

    return dictionary
}

// Generate a sentence, using Markov text generation and the provided dictionary.
func generate_sentence(dictionary map[string][]string) string {

    current_word = dictionar
}

func main() {

    // Get dictionary
    dictionary := generate_dictionary()

    for k := range dictionary {
        fmt.Println("key:", k, "val:", dictionary[k])
    }


}

