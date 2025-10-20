package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <input.pdb>")
        return
    }

    inputFile := os.Args[1]
    outputFile := "cleaned_" + inputFile

    inFile, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inFile.Close()

    outFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outFile.Close()

    scanner := bufio.NewScanner(inFile)
    writer := bufio.NewWriter(outFile)

    fmt.Println("Deleted lines (HOH with zero occupancy):")

    for scanner.Scan() {
        line := scanner.Text()

        if strings.HasPrefix(line, "HETATM") && strings.Contains(line[17:20], "HOH") {
            occupancyStr := strings.TrimSpace(line[54:60])
            occupancy, err := strconv.ParseFloat(occupancyStr, 64)
            if err != nil || occupancy == 0.0 {
                fmt.Println(line) // Report the deleted line
                continue          // Skip writing this line
            }
        }

        writer.WriteString(line + "\n")
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
    }

    writer.Flush()
    fmt.Printf("Cleaning complete. Output written to %s\n", outputFile)
}
