# VarsizedInt - Variable-Sized Integer Encoding in Go

VarsizedInt is a Go package that offers functionality for encoding and decoding variable-sized integers. This encoding scheme is specifically designed to efficiently represent integers of different sizes, optimizing storage space for smaller values while accommodating larger integers.

## Installation

To use this package, you can include it in your Go project using the following go get command:

```bash
go get -u github.com/NIR3X/varsizedint

```

## Usage

Here is an example of how to use VarsizedInt in your Go code:

```go
package main

import (
	"fmt"
	"github.com/NIR3X/varsizedint"
)

func main() {
	// Example usage of VarsizedInt encoding and decoding

	// Encode and Decode a variable-sized integer
	originalValue := uint64(123456789)
	encodedData := make([]uint8, varsizedint.MaxSize)
	encodedSize := varsizedint.Encode(encodedData, originalValue)

	// Display encoded data
	fmt.Print("Encoded Data: ")
	for i := 0; i < encodedSize; i++ {
		fmt.Printf("%#02x ", encodedData[i])
	}
	fmt.Println()

	// Decode the data
	decodedValue := varsizedint.Decode(encodedData)
	fmt.Printf("Decoded Value: %d 0x%016x\n", decodedValue, decodedValue)
}
```

In this example, a variable-sized integer (`originalValue`) is encoded using VarsizedInt, and the resulting encoded data is displayed in hexadecimal format. Subsequently, the encoded data is decoded back to its original value. The example showcases the simplicity and effectiveness of VarsizedInt in encoding and decoding variable-sized integers. The `varsizedint.MaxSize` constant represents the maximum size in bytes that an encoded integer can occupy.

## Variable-Sized Integer Ranges and Efficiency

VarsizedInt is designed to efficiently encode variable-sized integers, optimizing storage based on the value range. The following table illustrates the relationship between the value range, encoded size, and the number of bits lost to size encoding:
| Value Range              | Encoded Size | Bits Lost to Size Encoding   |
| ------------------------ | ------------ | ---------------------------- |
| 0 - 127                  | 1 byte       | 1 bit                        |
| 0 - 16383                | 2 bytes      | 2 bits                       |
| 0 - 2097151              | 3 bytes      | 3 bits                       |
| 0 - 268435455            | 4 bytes      | 4 bits                       |
| 0 - 34359738367          | 5 bytes      | 5 bits                       |
| 0 - 4398046511103        | 6 bytes      | 6 bits                       |
| 0 - 562949953421311      | 7 bytes      | 7 bits                       |
| 0 - 72057594037927935    | 8 bytes      | 8 bits                       |
| 0 - 18446744073709551615 | 9 bytes      | 8 bits (limited by encoding) |

Feel free to integrate VarsizedInt into your Go projects to efficiently store variable-sized integers with reduced storage overhead.

## License

[![GNU AGPLv3 Image](https://www.gnu.org/graphics/agplv3-155x51.png)](https://www.gnu.org/licenses/agpl-3.0.html)

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
