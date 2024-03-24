# NetCount

NetCount is a simple CLI tool designed to quickly calculate and display the number of IP addresses within specified CIDR (Classless Inter-Domain Routing) ranges. It reads CIDR ranges from standard input (stdin), making it versatile for both interactive use and as part of a pipeline in shell scripts.

## Features

- Reads CIDR ranges from standard input.
- Calculates the number of IP addresses in each CIDR range.
- Outputs the CIDR range alongside its corresponding count of IP addresses.

## Installation

To use NetCount, you'll need to have Go installed on your system. You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

With Go installed, you can clone this repository and build the tool with:

```
git clone https://github.com/yourusername/netcount.git
cd netcount
go build
```
This will create the `netcount` executable.

or

```
go install github.com/cybercdh/netcount@latest
```

## Usage

NetCount is designed to be simple to use. You can pipe CIDR ranges directly into NetCount or use it with files containing CIDR ranges.

### Direct Input

`echo "192.168.1.0/24" | ./netcount`

### Using a File

If you have a file named `cidrs.txt` with one CIDR range per line, you can use:

`cat cidrs.txt | ./netcount`

### Output

NetCount will output each CIDR range followed by the number of IP addresses it contains, like so:

`192.168.1.0/24, 256`

## Contributing

Contributions are welcome! If you have suggestions for improvements or encounter any issues, please feel free to open an issue or submit a pull request.

## License

NetCount is released under the MIT License. See the LICENSE file for more details.