## How It Works

The program:
1. Accepts one or more filenames as command-line arguments
2. Processes each file sequentially
3. For each file:
   - Counts lines by scanning line-by-line
   - Counts words by splitting each line into fields
   - Counts characters by measuring the length of each line
4. Displays individual file statistics
5. Shows total counts when processing multiple files

## Error Handling

If a file cannot be opened, the program will print an error message to stderr and continue processing other files.

## Usage

To run the program, use the following command:

```
go build -o wc .
./wc a.txt b.txt
```

### Output Format

The program outputs the following for each file:

```
[lines] [words] [chars] [filename]
1 1 10 a.txt
```

If multiple files are provided, a total count is displayed at the end.

### Example

```
[lines] [words] [chars] [filename]
1 1 10 a.txt
1 1 10 b.txt
2 2 20 total
```