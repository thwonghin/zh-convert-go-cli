# zhconvert CLI

This is a command-line wrapper for the [Fanhuaji](https://zhconvert.org) API. It allows you to convert text between various Chinese variants (Simplified, Traditional, Taiwan, Hongkong, etc.) via CLI.

## Features

* Full support for all Fanhuaji conversion options
* Accepts input via stdin, writes output to stdout

## Usage

```sh
# Convert text using stdin
cat input.txt | ./zhconvert --converter=Traditional > output.txt
```

### Example

```sh
echo "繁體字測試" | ./zhconvert --converter=Simplified
```

## Building

```sh
make build
```

## Disclaimer

This tool is built on top of the [Fanhuaji API](https://zhconvert.org). Please respect their terms of use and properly attribute their service if you redistribute or build upon this tool.

## License

MIT
