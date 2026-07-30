module github.com/zed-pkg-test/go-app

go 1.22

require github.com/zed-pkg-test/go-lib v1.0.0

// zed installs go-lib into [install].dir; this replacement is the whole seam
// between the two package managers. Ordinary proxy-sourced requires above are
// resolved by Go exactly as usual.
replace github.com/zed-pkg-test/go-lib => ./.vendor/.zed/zed-pkg-test/go-lib
