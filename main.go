package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"slices"
)

func createCliDirectory() (string, error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    fmt.Println("Error while getting the home directory:", err)
    return "", err
  }
  cliDirectory := homeDir + "/.base64-to-file/"
  // create a directory if it doesn't exist
  if _, err := os.Stat(cliDirectory); os.IsNotExist(err) {
    err := os.Mkdir(cliDirectory, 0755)
    if err != nil {
      fmt.Println("Error while creating a directory:", err)
      return "", err
    }
  }
  return cliDirectory, nil
}

func createFileWith(cliDirectory string) (string, error) {
  input := os.Args[2]
  fmt.Println("Input is:", input)

  fileName := cliDirectory + "input.txt"
  err := os.WriteFile(fileName, []byte(input), 0644) 
  if(err != nil)  {
    return "", err
  }
  return fileName, nil
}

func showHelp() {
  fmt.Println("\nYou are seeing this due to some invalid input or you asked for help\n")
  fmt.Println("Refer to the following help section of base64-to-file tool")
  fmt.Println("Usage: btof [flag] [input]")
  fmt.Println("Flags:")
  fmt.Println("  -h, --help   shows help")
  fmt.Println("  -i, --input <input_string>  takes input as a string")
  fmt.Println("  -f, --file  <file_name> <decoded_file_type> takes input from a file")
  fmt.Println("\nExamples:\n")
  fmt.Println("  btof -i 'aGVsbG8gd29ybGQ=' (This creates a txt file with the decoded content)\n")
  fmt.Println("  btof -f input.txt pdf (This takes base64 content from input.txt and creates a pdf file with the decoded content)\n")
  fmt.Println("  btof -f input.txt (This takes base64 content from input.txt and creates a txt file with the decoded content)\n")
  fmt.Println("\n")
}

func showHelpAndReturnTrueIfHelpShown() bool {
  if len(os.Args) < 3 {
    showHelp()
    return true
  }
  validFlags := []string{"-h", "--help", "-i", "--input", "-f", "--file"}

  flag := os.Args[1]
  if(flag == "-h" || flag == "--help" || slices.Contains(validFlags, flag) == false) {
    fmt.Println("\nInvalid flag:", flag, "\n")
    showHelp()
    return true
  }
  return false
}

func main() {
  cliDirectory, err := createCliDirectory()
  if err != nil {
    fmt.Println("Error while creating a directory:", err)
    return
  }
  isHelpShown := showHelpAndReturnTrueIfHelpShown()
  if isHelpShown {
    return
  }

  var fileName string
  var fileType string
  fileType = "txt"

  flag := os.Args[1]
  if(flag == "-i" || flag == "--input") {
    fileName, err = createFileWith(cliDirectory)
    if err != nil {
      fmt.Println("Error while processing the input:", err)
      return
    }
  }

  if(flag == "-f" || flag == "--file") {
    fileName = os.Args[2]
    if(len(os.Args) > 3) {
      fileType = os.Args[3]
    }
    fmt.Println("Reading the file:", fileName)
  }

  // read the file content and save it in a variable named input
  inputByte, err := os.ReadFile(fileName)
  if err != nil {
    fmt.Println("Error while reading the file:", err)
    return
  }
  input := string(inputByte)

  // decode the input and store it in a byte array
  decoded, err := base64.StdEncoding.DecodeString(input)
  if err != nil {
    fmt.Println("Error while decoding the string:", err)
    return
  }

  // write to a new pdf file with the decoded content
  decodedFileName := cliDirectory + "decoded." + fileType
  err = os.WriteFile(decodedFileName, decoded, 0644)  // 0 110 100 100  --> 0 rwx rwx rwx
  if err != nil {
    fmt.Println("Error while writing the decoded content to a new file:", err)
    return
  }
  fmt.Println("The file with decoded content can be found at:", decodedFileName)
  fmt.Println("\nRun `open " + decodedFileName + "` to open the file\n")
}



