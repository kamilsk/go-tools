Flag `-ungroup` added possibility to reset custom import sorting ([diff][]).

## Example

Input:

```go
import (
	"context"
	"github.com/pkg/b"
	"github.com/pkg/a"

	"github.com/pkg/y"
	"github.com/pkg/x"
	"github.com/pkg/z"
)
```

Simple `goimports` output:

```go
import (
	"context"

	"github.com/pkg/a"
	"github.com/pkg/b"

	"github.com/pkg/x"
	"github.com/pkg/y"
	"github.com/pkg/z"
)
```

With `-ungroup` flag output:

```go
import (
	"context"

	"github.com/pkg/a"
	"github.com/pkg/b"
	"github.com/pkg/x"
	"github.com/pkg/y"
	"github.com/pkg/z"
)
```

## How to install

### Homebrew

```bash
$ brew install kamilsk/tap/goimports
```

### Binary

```bash
$ curl -sSfL https://bit.ly/install-goimports | sh -s -- -b $(go env GOPATH)/bin
```

## How to run

As usual, just add `-ungroup` flag.

Happy hacking!

<!-- references -->

[diff]: https://github.com/golang/tools/compare/master...kamilsk:extended
