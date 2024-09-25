# Prusalink-go

This package offers a library to communicate with Prusa's PrusaLink enabled printers. It supports the full API available on the printer directly. A CLI is included to allow for easier scripting.

This is still in early development and any feedback or contributions would be greatly appreciated.

## Implementation

The current implementation is based on the related source code:

]- [v1](https://github.com/prusa3d/Prusa-Firmware-Buddy/blob/master/lib/WUI/link_content/prusa_link_api_v1.cpp)

The implementation has been tested against my Prusa Mini+. It is possible that some features behave differently on different models. If you come accross one of these instances please open an issue so this can be resolved.

## Development

To run the tests in the package you will need to have a Prusa printer with PrusaLink enabled. You will need to create a `.envrc` file in the root of the project with the following content:

```bash
PRUSALINK_HOST=xxxx
PRUSALINK_KEY=xxxx
```

You can then run the tests with `go test ./...` or by using the vscode go plugin.