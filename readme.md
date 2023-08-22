# Ignite to Feather

This repository contains the minimal changes required to convert a chain scaffolded using [Ignite CLI](https://docs.ignite.com/) to conform to the [Feather Core](https://github.com/terra-money/feather-core) interface.

## Steps

Please see [this commit](https://github.com/AaronCQL/ignite-to-feather/commit/e1b51387a382cd0ebd4e53e207956e916e5e818d) for a full example diff of the changes required.

### 1. Rename the `cmd/<app>d` directory to `cmd/feather-cored`

This ensures that the `Makefile` that will be added later references the correct directory. If your IDE does not automatically update the import changes caused by step 1, you can fix it manually (remember to replace `<go_module>` with your project's Go module):

```go
import (
  // "<go_module>/cmd/<app>d/cmd"
  "<go_module>/cmd/feather-cored/cmd"
)
```

### 2. Modify the `app/simulation_test.go` file

Rename the `BenchmarkSimulation` function to `TestFullAppSimulation`, and make sure it's taking in `testing.T` as args and not `testing.B`

```go
// OLD
// func BenchmarkSimulation(b *testing.B) {
//   ...
// }

// NEW
func TestFullAppSimulation(t *testing.T) {
  ...
}
```

### 3. Copy the `config` directory of `feather-core`

Copy the [`config` directory from `feather-core`](https://github.com/terra-money/feather-core/tree/main/config) into the root of your project. There should be exactly two files in the directory: `config.go` and `config.json`.

### 4. Copy the `Makefile` of `feather-core`

Copy the [`Makefile` from `feather-core`](https://github.com/terra-money/feather-core/blob/main/Makefile) into the root of your project. Though not essential, the `HTTPS_GIT` variable should be set to the correct GitHub URL of your project.

**Warning**: this makefile should NOT be edited unless you know what you're doing.

### 5. Copy the `contrib/devtools/Makefile` of `feather-core`

Copy the [`contrib/devtools/Makefile` from `feather-core`](https://github.com/terra-money/feather-core/blob/main/contrib/devtools/Makefile) into the root of your project.

**Warning**: this makefile should NOT be edited unless you know what you're doing.

### 6. Modify the `app/app.go` file

Firstly, declare new variables to be used as [ldflags during compile time](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications):

```go
// OLD
// const (
//   AccountAddressPrefix = "cosmos"
//   Name                 = "pluto"
// )

// NEW
// Note: 'const' is changed to 'var' and the values are no longer initialised.
// DO NOT change the names and values of these variables! They are populated by the `init` function.
var (
  AccountAddressPrefix string
  Name                 string
  BondDenom            string
  CoinType             uint32
)
```

Then, update the `init` function by loading configurations from `config/config.json` (note that the added code MUST be placed at the start of the `init` function):

```go
func init() {
  // ADD THIS TO START OF FUNCTION
  // Load and use config from config.json
  config, err := cfg.Load()
  if err != nil {
   panic(err)
  }
  Name = config.AppName
  BondDenom = config.BondDenom
  AccountAddressPrefix = config.AddressPrefix
  // Feather chains' coin type should follow Terra's coin type.
  // WARNING: changing this value will break feather's assumptions and functionalities.
  CoinType = 330

  // REST OF ORIGINAL FUNCTION
  ...
}
```

### 7. Modify the `cmd/feather-cored/cmd/config.go` file

Add two additional lines of configurations immediately before calling `config.Seal()`:

```go
config.SetCoinType(app.CoinType)     // Set coin type
sdk.DefaultBondDenom = app.BondDenom // Set default bond denom
```
