# Sum Calculator with Parallel Processing

This program reads a JSON file containing an array of objects and calculates the sum of all numbers in the file using goroutines for parallel processing.

## Overview

- **Language:** Go (Golang)
- **Features:**
    - Generates a JSON file with 1,000,000 objects.
    - Reads and processes the JSON file using streaming to minimize memory usage.
    - Utilizes goroutines and waitgroups for parallel computation.
    - Provides memory usage metrics and allocation statistics.
    - Allows the number of worker goroutines to be specified as a command-line argument.

## File Structure

- `main.go`: The main program file containing all the code.

## Requirements

- **Go version:** Go 1.13 or higher is recommended.
- **Git:** For cloning the repository (if applicable).

## Installation

1. **Clone the repository** (if hosted on GitHub):

   ```bash
   git clone https://github.com/nurtidev/sum_calculator.git
   cd sum_calculator
   ```

2. **Ensure Go is installed** on your system:

   ```bash
   go version
   ```

   If Go is not installed, download it from the [official website](https://golang.org/dl/).

## Usage

### Build the Program

Compile the program using the following command:

```bash
go build -o sum_calculator main.go
```

This will generate an executable named `sum_calculator`.

### Run the Program

The program accepts three command-line arguments:

1. **File Path:** Path to the JSON file to read or generate.
2. **Number of Workers:** Number of goroutines to use for parallel processing.
3. **Generate Flag:** Specify `true` to generate a new JSON file, or `false` to use an existing file.

#### Example: Generate JSON File and Calculate Sum

To generate a JSON file and calculate the sum:

```bash
./sum_calculator data.json 10 true
```

- `data.json`: The JSON file to generate and read.
- `10`: Use 10 worker goroutines.
- `true`: Generate the JSON file before processing.

#### Example: Calculate Sum Using Existing JSON File

If you have already generated the JSON file and want to calculate the sum:

```bash
./sum_calculator data.json 10 false
```

- `false`: Do not generate the JSON file; use the existing one.

### Program Output

The program will display:

- **Total Sum:** The sum of all numbers in the JSON file.
- **Memory Used:** The amount of memory used during processing.
- **Number of Allocations:** The number of memory allocations made.

#### Sample Output

```
Generating JSON file...
JSON file generated successfully.
Processing JSON file...
Total sum: -1593
Memory used: 4.75 MB
Number of allocations: 1001560
```

## Code Explanation

- **Generating JSON File:**
    - Generates 1,000,000 objects with random integers between -10 and 10.
    - Writes the objects to the specified JSON file using streaming to minimize memory usage.

- **Processing JSON File:**
    - Reads the JSON file using a streaming decoder (`json.Decoder`).
    - Uses a buffered channel to pass objects to worker goroutines.
    - Each worker computes a local sum of the numbers it processes.
    - Uses a `sync.WaitGroup` to wait for all goroutines to finish.
    - Protects the shared `totalSum` variable using a `sync.Mutex`.

- **Memory Metrics:**
    - Collects memory statistics before and after processing using `runtime.ReadMemStats`.
    - Calculates the memory used and the number of allocations during processing.

## Customization

- **Number of Objects:**
    - Change the number of objects generated by modifying the value in the `generateJSONFile` function call.

- **Range of Random Numbers:**
    - Adjust the range `rand.Intn(21) - 10` to change the range of generated numbers.

## Performance Tips

- **Number of Workers:**
    - Set the number of workers to match the number of CPU cores for optimal performance.
    - Use the `GOMAXPROCS` environment variable if you need to control the number of OS threads.

- **Memory Usage:**
    - The program is designed to minimize memory usage by streaming data.
    - Monitor memory usage if processing very large files or when running on memory-constrained systems.

## Troubleshooting

- **Invalid Number of Workers:**
    - Ensure that the second argument is a valid integer greater than 0.

- **File Access Errors:**
    - Check that you have read/write permissions for the specified file path.

- **Go Version Compatibility:**
    - Ensure you are using a compatible version of Go (1.13 or higher recommended).

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## Contact

For any questions or suggestions, please contact 
- email: [nurtilek.develop@gmail.com](mailto:nurtilek.develop@gmail.com)
- telegram: @nurtilek_assankhan