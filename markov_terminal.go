package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "math/rand"
    "time"
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

    // Random seed used for selection of words.  Use time for something approximating randomness.
    rand.Seed(int64(time.Now().Nanosecond()))

    // Grab first word in sentence.
    current_slice := dictionary["$"]
    current_word := current_slice[rand.Intn(len(current_slice))]
    words := []string{current_word}

    // Continue until we reach the end of a sentence.
    for !strings.ContainsAny(current_word, punctuation) {

        // Find new slice and new word from that slice.
        current_slice = dictionary[current_word]
        current_word = current_slice[rand.Intn(len(current_slice))]

        // Add new word to sentence.
        words = append(words, current_word)
    }

    // Join all words into proper sentence and return.
    return strings.Join(words, " ")
}

// Version of generate_sentence meant to be used as a goroutine.
func go_generate_sentence(dictionary map[string][]string, channel chan string) {
    out := generate_sentence(dictionary)
    channel <- out
}

// Generate N sentences, using generate_sentence as a helper function and goroutines.
func generate_sentences(dictionary map[string][]string, count int) []string {

    // Put sentences here.
    sentences := make([]string, count)

    c := make(chan string, count)

    for i := 0; i < count; i++ {
        go go_generate_sentence(dictionary, c)
    }

    for i := 0; i < count; i++ {
        sentences[i] = <- c
    }

    return sentences
}

func main() {

    // Get dictionary
    dictionary := generate_dictionary()

    //for k := range dictionary {
    //    fmt.Println("key:", k, "val:", dictionary[k])
    //}

    // Generate sentence.
    //sentence := generate_sentence(dictionary)

    sentences := generate_sentences(dictionary, 4)
    fmt.Println("number sentences:", len(sentences))
    for i, v := range sentences {
        fmt.Println(i, v)
    }


}

