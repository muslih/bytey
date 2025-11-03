# Bytey - Your Playful Size Formatter 🎭

> *Because 1048576 bytes deserves to be call## Contributing

Feel free to submit issues and enhancement requests!

## License

MIT License - see LICENSE file for details.h love. ❤️*

A fun and friendly Go library that converts file sizes into human-readable formats with personality!

## Features

- **Talk()** - Get playful size strings with random comments
- **Size()** - Clean, formal size formatting  
- **Whisper()** - Gentle size display with cute emojis
- **Scream()** - Dramatic size announcements for big files

## Installation

```bash
go get github.com/muslih/bytey
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/muslih/bytey"
)

func main() {
    b := bytey.NewBytey()
    
    // Different personality modes
    fmt.Println(b.Talk(1048576))    // "1MB — File's been hitting the buffet 🍝"
    fmt.Println(b.Size(1048576))    // "1MB"
    fmt.Println(b.Whisper(1048576)) // "1MB ✨"
    fmt.Println(b.Scream(1048576))  // "🚨 BYTEY ALERT 🚨 1MB — Much data, very wow 🐶"
}
```

## API Reference

### `NewBytey() *ByteyFacade`
Creates a new Bytey instance.

### `Talk(size int64) string`
Returns a human-readable size string with a playful random comment.

### `Size(size int64) string` 
Returns a formal, clean human-readable size string.

### `Whisper(size int64) string`
Returns a quiet, gentle version with a random emoji.

### `Scream(size int64) string`
Returns an exaggerated, dramatic output for large files.

## Examples

```go
b := bytey.NewBytey()

// Various file sizes
b.Talk(42)                    // "42B — Just a byte-sized snack 🍪"
b.Talk(1024)                  // "1KB — Baby file 👶 - so cute!"
b.Talk(1024 * 1024)           // "1MB — This file's been hitting the buffet 🍝"
b.Talk(1024 * 1024 * 1024)    // "1GB — Too heavy for email attachments 😅"

// Edge cases
b.Talk(-100)  // "Negative bytes? Are you reversing time, human? ⏳"
b.Talk(0)     // "0B — Just a byte-sized snack 🍪"
```

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run the demo
go run demo/main.go
```

## Contributing

Feel free to submit issues and enhancement requests!

## License

MIT License - see LICENSE file for details. - Your Bytes, But Friendlier

> *Because 1048576 bytes deserves to be called 1 MB - with love. ❤️*

bytey is a Go package that gives your file sizes some humanity.
No more staring at cold numbers - let bytey whisper,
*“Hey, that’s 2.5GB of cat videos right there.” 🐱📹*