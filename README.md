# base64-to-file

**base64-to-file** is a simple command-line tool that decodes Base64-encoded content and creates a file with the provided file type extension. It can be useful when you have Base64-encoded data that you want to decode and save as a file, such as images, documents, or any other file type.

## Usage

btof [flag] [input]

### Flags

- `-h, --help`: Shows help.
- `-i, --input <input_string>`: Takes input as a string.
- `-f, --file <file_name> <decoded_file_type>`: Takes input from a file.

## Examples

1. Decoding from a string:  ```btof -i 'aGVsbG8gd29ybGQ='```
   This creates a **txt** file with the decoded content.

2. Decoding from a file and specifying the file type: ```btof -f input.txt pdf```
   This takes base64 content from `input.txt` and creates a **pdf** file with the decoded content.

3. Decoding from a file without specifying the file type (defaults to txt): btof -f input.txt
   This takes base64 content from `input.txt` and creates a **txt** file with the decoded content.

## Installation

Run the command `brew install yogeshnikam671/btof/btof` to install this CLI tool.
