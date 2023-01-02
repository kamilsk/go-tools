## Motivation

As a Go developer, I actively use `goimports` as a code formatter, but it has
non-deterministic behavior that affects a final diff. For example, for the input
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
the output will be
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
The proper result for me is
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
So, I need to patch its behaviour and make a result deterministic.

## What's changed

- The `goimports` has the deterministic behaviour.
- The tool is available by `brew install kamilsk/tap/goimports`.
- The installation is also possible by
  ```bash
  $ curl -sSfL https://install.octolab.org/goimports | sh -s -- -b /usr/local/bin
  ```

## Related issues

- https://github.com/golang/tools/pull/68
- https://github.com/golang/tools/pull/308
- https://github.com/golang/tools/compare/master...kamilsk:extended

<p align="right">made with ❤️ for everyone by <a href="https://www.octolab.org/">OctoLab</a></p>
