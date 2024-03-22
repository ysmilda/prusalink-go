# Prusalink-go

This package offers a library to communicate with Prusa's PrusaLink enabled printers. An CLI is included to allow for easier scripting.

This is still in early development and any feedback or contributions would be greatly appreciated.

## Implementation

The current implementation is based on the published OpenAPI specs and the related source code:

- [v1](https://github.com/prusa3d/Prusa-Link-Web/blob/master/spec/openapi.yaml)
- [source](https://github.com/prusa3d/Prusa-Firmware-Buddy/blob/master/lib/WUI/link_content/prusa_link_api_v1.cpp)

The implementation has been tested against my Prusa Mini+. It is possible that some features behave differently on different models. If you come accross one of these instances please open an issue so this can be resolved.

## TODO

- [ ] Implement support for the legacy api definition (`v0`)
- [ ] Support responses for file commands. These are `oneof(FileInfo, PrintFileInfo, FirmwareFileInfo, FolderInfo)`. For now they return an interface.
- [ ] Define clearer errors.
- [ ] Extend CLI to support cameras, currently unable to test this due to available hardware.
- [ ] Figure out what the `/update` endpoint is supposed to do.
- [ ] Allow setting CLI inputs via environment.
