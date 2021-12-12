## Motivation

`goimports` has inconsistent behavior that affects a final diff. For example

- input

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

- output

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

The flag `-ungroup` added possibility to reset custom import sorting.

## Related issues

- https://github.com/golang/tools/pull/68
- https://github.com/golang/tools/pull/308
- https://github.com/golang/tools/compare/master...kamilsk:extended

## What's changed

- The `goimports` supports `-ungroup` flag to reset custom import sorting.
- The tool is available by `brew install kamilsk/tap/goimports`.
- The installation is also possible by
  `curl -sSfL https://raw.githubusercontent.com/kamilsk/go-tools/extended/bin/install | sh -s -- -b /usr/local/bin`.

<!-- references -->
