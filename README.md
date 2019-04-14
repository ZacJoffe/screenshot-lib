# Screenshot Library
A simple Golang library for taking screenshots. It includes functions for taking screenshots of the whole screen or just a selected region.

The library uses [`ImageMagick`](https://wiki.archlinux.org/index.php/ImageMagick) as its backend. Make sure to install it for proper use!


# Installation and Use
Install the library:
```
go get -u github.com/ZacJoffe/screenshot-lib/screenshot
```

Import the library in your code:
```
import (
    "github.com/ZacJoffe/screenshot-lib/screenshot"
)
```

Use the `Screen` function for taking a screenshot of the whole screen, and the `Select` function for taking a screenshot of a selected region.

See `example.go` in the root folder for small example program.