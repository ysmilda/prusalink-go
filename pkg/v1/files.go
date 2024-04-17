package v1

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ysmilda/prusalink-go/pkg/rest"
	"github.com/ysmilda/prusalink-go/pkg/utils"
)

type filesHandler struct {
	printer *Printer
}

// Returns a list of files and folders in the given storage at the given path
// TODO: Implement oneOf (FileInfo, PrintFileInfo, FirmwareFileInfo, FolderInfo)
// TODO: Add support for recursive listing?
func (f filesHandler) List(storage, path string) (interface{}, error) {
	return parseAsJSON[interface{}](f.printer.get(fmt.Sprintf("/api/v1/files/%s/%s", storage, path)))
}

// Creates a new folder in the given storage at the given path.
func (f filesHandler) CreateFolder(storage, path string) error {
	_, err := f.printer.put(
		fmt.Sprintf("/api/v1/files/%s/%s", storage, path),
		nil,
		map[string]string{
			"Content-Length": "0",
			"create-folder":  "?1",
		},
	)
	if err != nil {
		return fmt.Errorf("could not create folder: %w", f.parseError(err))
	}
	return nil
}

// Uploads a .gcode or .bgcode file to the given storage at the given path.
func (f filesHandler) Upload(
	storage, folderPath, fileName string, content []byte, overwrite, printAfterUpload bool,
) error {
	if !hasValidExtension(fileName) {
		return ErrNonGcodeFile
	}
	if len(content) == 0 {
		return ErrEmptyFile
	}
	_, err := f.printer.put(
		fmt.Sprintf(
			"/api/v1/files/%s/%s",
			storage,
			folderPath+"/"+fileName,
		),
		content,
		map[string]string{
			"Content-Type":       "application/octet-stream",
			"Content-Length":     fmt.Sprintf("%d", len(content)),
			"Print-After-Upload": fmt.Sprintf("?%d", utils.BoolToInt(printAfterUpload)),
			"Overwrite":          fmt.Sprintf("?%d", utils.BoolToInt(overwrite)),
		},
	)
	if err != nil {
		return fmt.Errorf("could not upload file: %w", f.parseError(err))
	}
	return nil
}

// Starts the print of the file at the given path in the given storage.
func (f filesHandler) StartPrint(storage string, path string) error {
	if !hasValidExtension(path) {
		return ErrNonGcodeFile
	}
	_, err := f.printer.post(fmt.Sprintf("/api/v1/files/%s/%s", storage, path), nil)
	if err != nil {
		return fmt.Errorf("could not start print: %w", f.parseError(err))
	}
	return nil
}

func (f filesHandler) State(storage, path string) error {
	// TODO: Implement
	return nil
}

func (f filesHandler) Delete(storage, path string, deleteNonEmptyFolder bool) error {
	_, err := f.printer.delete(
		fmt.Sprintf("/api/v1/files/%s/%s", storage, path),
		map[string]string{
			"Force": fmt.Sprintf("?%d", utils.BoolToInt(deleteNonEmptyFolder)),
		},
	)
	if err != nil {
		return fmt.Errorf("could not delete file: %w", f.parseError(err))
	}
	return nil
}

func (f filesHandler) parseError(err error) error {
	if errors.Is(err, rest.ErrNotFound) {
		return ErrStorageNotFound
	}
	if errors.Is(err, rest.ErrConflict) {
		return ErrAlreadyExists
	}
	return err
}

func hasValidExtension(fileName string) bool {
	for _, extension := range []string{".gcode", ".bgcode"} {
		if strings.HasSuffix(fileName, extension) {
			return true
		}
	}
	return false
}
