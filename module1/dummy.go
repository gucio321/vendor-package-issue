//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy

import _ "github.com/gucio321/vendor-package-issue/module1/vendor"
