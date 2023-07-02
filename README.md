# Favicon-Hash-Checker

A simple GoLang script to fetch the favicon.ico file from a given domain and calculate its hash using MurmurHash3 algorithm.

## Usage

1. Clone the repository:

```bash
$ git clone https://github.com/dzhenway/Favicon-Hash-Checker.git

```Run the script, providing the domain name as a command-line argument:

$ **./favicon-hash-checker -d example.com**

Replace example.com with the domain name you want to check.

The script will fetch the favicon.ico file from the specified domain and calculate its hash using MurmurHash3. The resulting hash value will be displayed on the console.

Dependencies
This script depends on the following external libraries:

github.com/twmb/murmur3 - MurmurHash3 library for GoLang
Ensure that you have the dependencies installed before running the script.

Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to submit a pull request or open an issue in the repository.

License
This project is licensed under the MIT License. See the LICENSE file for more details.
