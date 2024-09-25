package v1

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ysmilda/prusalink-go/pkg/printer"
	"github.com/ysmilda/prusalink-go/pkg/utils"
)

type filesHandler struct {
	conn *printer.Conn
}

// Returns a list of files and folders in the given storage at the given path
// TODO: Implement oneOf (FileInfo, PrintFileInfo, FirmwareFileInfo, FolderInfo)
// TODO: Add support for recursive listing?
func (f filesHandler) List(storage, path string) (*FileInfo, error) {
	return utils.ParseAsJSON[FileInfo](f.conn.Get(fmt.Sprintf("/api/v1/files/%s/%s", storage, path)))
}

// Creates a new folder in the given storage at the given path.
func (f filesHandler) CreateFolder(storage, path string) error {
	_, err := f.conn.Put(
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
		return printer.ErrNonGcodeFile
	}
	if len(content) == 0 {
		return printer.ErrEmptyFile
	}
	_, err := f.conn.Put(
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
		return printer.ErrNonGcodeFile
	}
	_, err := f.conn.Post(fmt.Sprintf("/api/v1/files/%s/%s", storage, path), nil)
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
	_, err := f.conn.Delete(
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
	if errors.Is(err, printer.ErrNotFound) {
		return printer.ErrStorageNotFound
	}
	if errors.Is(err, printer.ErrConflict) {
		return printer.ErrAlreadyExists
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
