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

// Returns a list of files and folders in the given storage at the given path.
// The returned FileInfo contains information about the files and folders in the given path.
// The list is only one level deep.
func (f filesHandler) List(storage, path string) (*FileInfo, error) {
	return printer.ParseAsJSON[FileInfo](f.conn.Get(fmt.Sprintf("/api/v1/files/%s/%s", storage, path)))
}

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

// Upload uploads a .gcode or .bgcode file to the given storage at the given path.
// Overwrite indicates whether the file should be overwritten if it already exists, and printAfterUpload
// indicates whether the file should be printed after upload.
func (f filesHandler) Upload(
	storage, folderPath, fileName string, content []byte, overwrite, printAfterUpload bool,
) error {
	if !hasValidExtension(fileName) {
		return printer.ErrNonGcodeFile
	}
	if len(content) == 0 {
		return printer.ErrEmptyFile
	}

	headers := map[string]string{
		"Content-Length": fmt.Sprintf("%d", len(content)),
	}
	if strings.HasSuffix(fileName, ".bgcode") {
		headers["Content-Type"] = "application/gcode+binary"
	} else {
		headers["Content-Type"] = "text/x.gcode"
	}
	if printAfterUpload {
		headers["Print-After-Upload"] = "?1"
	}
	if overwrite {
		headers["Overwrite"] = "?1"
	} else {
		headers["Overwrite"] = "?0"
	}

	_, err := f.conn.Put(
		fmt.Sprintf(
			"/api/v1/files/%s/%s",
			storage,
			folderPath+"/"+fileName,
		),
		content,
		headers,
	)
	if err != nil {
		return fmt.Errorf("could not upload file: %w", f.parseError(err))
	}
	return nil
}

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

func (f filesHandler) Info(storage, path string) (*FileInfo, error) {
	return printer.ParseAsJSON[FileInfo](f.conn.Head(fmt.Sprintf("/api/v1/files/%s/%s", storage, path)))
}

func (f filesHandler) Delete(storage, path string, deleteNonEmptyFolder bool) error {
	_, err := f.conn.Delete(
		fmt.Sprintf("/api/v1/files/%s/%s", storage, path),
		map[string]string{
			"Force": fmt.Sprintf("?%d", utils.BoolToInt(deleteNonEmptyFolder)),
		},
	)
	if err != nil {
		return fmt.Errorf("could not delete \"%s\": %w", path, f.parseError(err))
	}
	return nil
}

func (f filesHandler) parseError(err error) error {
	if errors.Is(err, printer.ErrNotFound) {
		return printer.ErrStorageNotFound
	}
	if errors.Is(err, printer.ErrConflict) {
		return printer.ErrConflict
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
