# AGENTS.md - Context for AI Assistants

This file provides technical context about the advent-of-code Bazel monorepo for AI assistants working on the code.

## Overview

This is a Bazel monorepo containing Advent of Code solutions for multiple years (currently 2025, with 2024 and 2023 to be migrated). It uses `go.work` to manage multiple independent Go modules - one per year plus a shared utilities module. Requires Go 1.25.3+ and Bazel 8.4+.

## Project Structure

The monorepo is organized as follows:

```
advent-of-code/
├── util/                    # Shared utilities module (used by all years)
│   ├── go.mod
│   ├── common/              # Generic helpers (Ptr)
│   ├── strconv/             # String conversion (MustParseInt)
│   ├── math/                # Math utilities (GCD, LCM)
│   ├── matrix/              # Matrix operations (Matrix type, Transpose, etc.)
│   ├── geometry/            # Geometric types (Point3D)
│   ├── interval/            # Interval operations
│   └── set/                 # Set data structure
├── years/
│   └── 2025/                # Advent of Code 2025 solutions
│       ├── go.mod
│       ├── day1/            # Day 1 solutions (1a.go, 1b.go, tests, inputs)
│       ├── day2/
│       └── ...
├── MODULE.bazel             # Bazel module configuration
├── BUILD.bazel              # Root BUILD file with gazelle setup
├── go.mod                   # Root go.mod (for Bazel Go SDK)
├── go.work                  # Go workspace file listing all modules
└── .bazelrc                 # Bazel configuration
```

Each year is a separate Go module for dependency isolation. The shared `util` module provides common utilities used across all years.

## Key Patterns

### Import Paths

- Year modules: `github.com/lonnblad/advent-of-code/years/2025`
- Shared utilities: `github.com/lonnblad/advent-of-code/util/<subpackage>`
  - Examples: `util/matrix`, `util/strconv`, `util/interval`, `util/geometry`

### Build System

- **Bazel** handles builds, dependency management, and test execution
- **Gazelle** auto-generates BUILD.bazel files for Go code (`bazel run //:gazelle -- -r .`)
- **Go workspace** (`go.work`) manages multiple modules in a single workspace

### Shared Utilities Module

The `util/` module is organized into focused sub-packages:

- **`common/`**: Generic helpers like `Ptr[T]` for creating pointers
- **`strconv/`**: String conversion utilities (`MustParseInt`)
- **`math/`**: Mathematical operations (`GCD`, `LCM`)
- **`matrix/`**: Matrix operations and types
  - `Matrix` type with methods: `IsWithinBounds()`, `GetAdjacentCells()`, `Find()`, `String()`
  - Functions: `IntMatrix()`, `StringMatrix()`, `NewMatrix()`, `Transpose()`, `Rotate45Degrees()`
- **`geometry/`**: Geometric types (`Point3D` with `Distance()` method)
- **`interval/`**: Interval operations (`Interval` type with methods like `Contains()`, `HasOverlap()`, `Union()`, `Intersection()`)
- **`set/`**: Set data structure (`Set[T]` with methods like `Add()`, `Contains()`, `Values()`, `Size()`)

**Dependencies between sub-packages:**
- `interval` imports `strconv` for parsing
- `matrix` imports `strconv` for parsing integers

### Day Structure

Each day follows a consistent structure:

```
dayX/
├── Xa.go                    # Part A solution
├── Xa_test.go               # Part A tests
├── Xa_input.example.txt     # Example input for tests
├── Xa_input.txt             # Real input
├── Xb.go                    # Part B solution
├── Xb_test.go               # Part B tests
├── Xb_input.example.txt     # Example input for tests
└── Xb_input.txt             # Real input
```

**Test Pattern:**
- Tests use `testify` for assertions (`require.Equal`, `require.NoError`)
- Tests embed input files using `//go:embed`
- Tests verify both example and real inputs

**Function Naming:**
- Solution functions: `DayXA(input string) (_ int, err error)` for part A, `DayXB(...)` for part B
- Helper functions: Lowercase, descriptive names

## Adding New Days

When asked to create a new day (e.g., "create day 09"), use this template:

### Step 1: Create Directory Structure

Create `years/2025/day9/` directory with the following files:
- `9a.go` - Part A solution
- `9a_test.go` - Part A tests
- `9a_input.example.txt` - Example input for Part A
- `9a_input.txt` - Real input for Part A (can be empty initially)
- `9b.go` - Part B solution
- `9b_test.go` - Part B tests
- `9b_input.example.txt` - Example input for Part B
- `9b_input.txt` - Real input for Part B (can be empty initially)

### Step 2: Template Files

**`9a.go` template:**
```go
package day9

func Day9A(input string) (_ int, err error) {
	// Implementation here
	return 0, nil
}
```

**`9a_test.go` template:**
```go
package day9_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day9"
	"github.com/stretchr/testify/require"
)

//go:embed 9a_input.example.txt
var exampleInput9A string

const expectedExampleOutput9A = -1 // Update with actual expected value

func Test_9A_Example(t *testing.T) {
	actualOutput9A, actualErr := day9.Day9A(exampleInput9A)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput9A, actualOutput9A)
}

//go:embed 9a_input.txt
var input9A string

const expectedOutput9A = -1 // Update with actual expected value

func Test_9A_Input(t *testing.T) {
	actualOutput9A, actualErr := day9.Day9A(input9A)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput9A, actualOutput9A)
}
```

**`9b.go` template:**
```go
package day9

func Day9B(input string) (_ int, err error) {
	// Implementation here
	return 0, nil
}
```

**`9b_test.go` template:**
```go
package day9_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code/years/2025/day9"
	"github.com/stretchr/testify/require"
)

//go:embed 9b_input.example.txt
var exampleInput9B string

const expectedExampleOutput9B = -1 // Update with actual expected value

func Test_9B_Example(t *testing.T) {
	actualOutput9B, actualErr := day9.Day9B(exampleInput9B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutput9B, actualOutput9B)
}

//go:embed 9b_input.txt
var input9B string

const expectedOutput9B = -1 // Update with actual expected value

func Test_9B_Input(t *testing.T) {
	actualOutput9B, actualErr := day9.Day9B(input9B)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutput9B, actualOutput9B)
}
```

**Input files:**
- `9a_input.example.txt` - Copy example input from Advent of Code
- `9a_input.txt` - Copy real input from Advent of Code (can be empty initially)
- `9b_input.example.txt` - Usually same as Part A example, or Part B specific example
- `9b_input.txt` - Usually same as Part A input, or Part B specific input

### Step 3: Post-Creation Steps

1. Run `bazel run //:gazelle -- -r .` to generate BUILD.bazel files
2. Run `bazel mod tidy` if you added new external dependencies
3. Run `bazel test //years/2025/day9:day9_test` to verify tests pass
4. Update expected values in test files once you have working solutions

### Key Points

- Package name: `day9` (lowercase, no leading zero)
- Function names: `Day9A`, `Day9B` (capital D, capital A/B, no leading zero)
- Test package: `day9_test` (package name + `_test`)
- Test function names: `Test_9A_Example`, `Test_9A_Input`, `Test_9B_Example`, `Test_9B_Input`
- File names: Use leading zero for single digits (`9a.go` not `09a.go`)
- Import path: `github.com/lonnblad/advent-of-code/years/2025/day9`

## Adding New Years

1. Copy the structure from an existing year (e.g., `years/2025/`)
2. Create `years/YYYY/go.mod` with module path `github.com/lonnblad/advent-of-code/years/YYYY`
3. Update all import paths in the new year's code
4. Add `./years/YYYY` to `go.work`
5. Run `go work sync` to sync workspace dependencies
6. Run `bazel run //:gazelle -- -r .` to generate BUILD files
7. Run `bazel mod tidy` to update MODULE.bazel

## Common Commands

```bash
# Build all code
bazel build //...

# Build specific year
bazel build //years/2025/...

# Run all tests
bazel test //...

# Run tests for specific year
bazel test //years/2025/...

# Run tests for specific day
bazel test //years/2025/day1:day1_test

# Update BUILD files after adding/modifying Go code
bazel run //:gazelle -- -r .

# Update dependencies
bazel mod tidy
go work sync
```

## Testing Strategy

- **Unit Tests**: Test individual solution functions with example and real inputs
- **Test Files**: Colocated with solution files (`Xa_test.go` next to `Xa.go`)
- **Input Files**: Embedded using `//go:embed` directive
- **Assertions**: Use `testify/require` for assertions
- **Test Naming**: `Test_XA_Example` and `Test_XA_Input` for part A, similar for part B

## Current Status

- **2025**: Fully migrated and working (all 8 tests passing)
- **2024**: To be migrated in future phase (has its own util package that will be consolidated)
- **2023**: To be migrated in future phase (no util package, but has duplicate GCD/LCM functions)

## Notes

- **BUILD.bazel files**: Auto-generated by gazelle - don't edit them manually
- **Module independence**: Each year is a separate Go module for dependency isolation
- **Shared utilities**: All years use the shared `util` module to avoid code duplication
- **Test failures**: Some tests may fail - this is expected and reflects the state of the original projects
- **Go version**: Root `go.mod` uses Go 1.24.2 for Bazel compatibility, but `go.work` and year modules use Go 1.25.3

## Utility Package Design Decisions

- **Sub-packages**: Utilities are split into focused sub-packages rather than a single package for better organization and clearer dependencies
- **Internal dependencies**: Sub-packages can import each other (e.g., `interval` imports `strconv`)
- **No circular dependencies**: Package structure avoids circular imports
- **Merged features**: Matrix package includes features from both 2024 and 2025 implementations (Find() from 2024, GetAdjacentCells() from 2025)

