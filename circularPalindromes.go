package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the circularPalindromes function below.
 */
func check(s string) bool{
    for i:=0; i<=len(s)/2; i++{

        if s[i] != s[len(s)-1-i]{
            return false
        }
    }
    return true
}
func findPalindrome(s string) int32{
    //var longest int32
    for i := len(s); i >= 2; i--{
        start := 0
        end := start + i
        for end < len(s)+1{
            if check(s[start:end]){
                fmt.Println(s[start:end])
                return int32(len(s[start:end]))
            }
            start++
            end++
        }
    }
    return 1
}
func circularPalindromes(s string) []int32 {
    /*
     * Write your code here.
     */
    var palindrome []int32

    copyString := s
    for i := 0; i < len(s); i++{
        palindrome = append(palindrome,findPalindrome(copyString)) 
        copyString = copyString[1:] + string(copyString[0])
        fmt.Println(palindrome)
    }

    return palindrome
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

   nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
   checkError(err)
    n := int32(nTemp)
    n+=1
    s := readLine(reader)

    result := circularPalindromes(s)

    for resultItr, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if resultItr != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
