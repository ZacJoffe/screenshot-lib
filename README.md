# Screenshot Library
Simple Golang libraries for taking screenshots. It includes functions for taking screenshots of the whole screen or just a selected region.

You can either use [`ImageMagick`](https://wiki.archlinux.org/index.php/ImageMagick) as or [`maim`](https://github.com/naelstrof/maim) for the backend. Both act similarly, so just pick whichever you prefer.


# Installation and Use
Install the library:
```
# imagemagick
go get -u github.com/ZacJoffe/screenshot-lib/imagemagick

# maim
go get -u github.com/ZacJoffe/screenshot-lib/maim
```

Import the library in your code:
```
import (
    "github.com/ZacJoffe/screenshot-lib/imagemagick"
    "github.com/ZacJoffe/screenshot-lib/maim"
)
```

Use the `Screen` function for taking a screenshot of the whole screen, and the `Select` function for taking a screenshot of a selected region.

See `example.go` in the root folder for small example program.