package v1

import "github.com/ysmilda/prusalink-go/pkg/printer"

type Info struct {
	NozzleDiameter   float64 `json:"nozzle_diameter"`
	MMU              bool    `json:"mmu"`
	Serial           string  `json:"serial"`
	Hostname         string  `json:"hostname"`
	MinExtrusionTemp int     `json:"min_extrusion_temp"`
}

type Status struct {
	Job      *StatusJob      `json:"job,omitempty"`
	Storage  StatusStorage   `json:"storage"`
	Transfer *StatusTransfer `json:"transfer,omitempty"`
	Printer  StatusPrinter   `json:"printer"`
}

type StatusJob struct {
	ID               int               `json:"id"`
	Progress         float64           `json:"progress"` // Percents
	TimeRemaining    *printer.Duration `json:"time_remaining,omitempty"`
	FilamentChangeIn *printer.Duration `json:"filament_change_in,omitempty"`
	TimePrinting     printer.Duration  `json:"time_printing"`
}

// Telemetry info about current storage status.
type StatusStorage struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	ReadOnly bool   `json:"read_only"`
}

// Telemetry info about current transfer status.
type StatusTransfer struct {
	ID               int     `json:"id"`
	TimeTransferring int     `json:"time_transferring"`
	Progress         float64 `json:"progress"` // Percents
	Transferred      *int    `json:"transferred"`
}

// Telemetry info about printer.
type StatusPrinter struct {
	State        string   `json:"state"`
	TempBed      float64  `json:"temp_bed"`
	TargetBed    float64  `json:"target_bed"`
	TempNozzle   float64  `json:"temp_nozzle"`
	TargetNozzle float64  `json:"target_nozzle"`
	AxisZ        float64  `json:"axis_z"`
	AxisX        *float64 `json:"axis_x,omitempty"` // Available only when printer is not moving
	AxisY        *float64 `json:"axis_y,omitempty"` // Available only when printer is not moving
	Flow         int      `json:"flow"`
	Speed        int      `json:"speed"`
	FanHotend    int      `json:"fan_hotend"`
	FanPrint     int      `json:"fan_print"`
}

type Storage struct {
	StorageList []StorageInfo `json:"storage_list"`
}

type StorageInfo struct {
	Name      string `json:"name"`      // Name of the storage, based on selected language
	Type      string `json:"type"`      // Storage source
	Path      string `json:"path"`      // Path to storage (not display path)
	Available bool   `json:"available"` // Whether the storage is available or not
	ReadOnly  bool   `json:"read_only"` // Whether the storage is read only
}

type Job struct {
	ID            int               `json:"id"`
	State         string            `json:"state"`
	Progress      float64           `json:"progress"` // Percents
	TimeRemaining *printer.Duration `json:"time_remaining,omitempty"`
	TimePrinting  printer.Duration  `json:"time_printing"`
	File          JobFile           `json:"file"`
}

type JobFile struct {
	Refs        FileRefs      `json:"refs"`
	Name        string        `json:"name"`         // Short Filename
	DisplayName string        `json:"display_name"` // Long Filename
	Path        string        `json:"path"`
	Size        *int          `json:"size,omitempty"` // Bytes
	Timestamp   *printer.Time `json:"m_timestamp,omitempty"`
}

type Transfer struct {
	Type             string           `json:"type"`
	DisplayName      string           `json:"display_name"` // Long Filename
	Path             string           `json:"path"`
	Size             string           `json:"size"`        // Bytes
	Progress         float64          `json:"progress"`    // Percents
	Transferred      int              `json:"transferred"` // Transfered data in bytes
	TimeRemaining    printer.Duration `json:"time_remaining,omitempty"`
	TimeTransferring printer.Duration `json:"time_transferring"`
	ToPrint          bool             `json:"to_print"` // Whether or not print after finishing transfer (upload)
}

// FileInfo represents a file or folder on the printer. If the file is a folder the Children field will be populated.
// The children struct only goes one level deep, the printer has no recursive listing.
type FileInfo struct {
	Name        string        `json:"name"`
	DisplayName string        `json:"display_name"`
	ReadOnly    bool          `json:"ro"`
	Type        string        `json:"type"`
	Timestamp   *printer.Time `json:"m_timestamp,omitempty"`
	Size        *int          `json:"size,omitempty"`
	Refs        *FileRefs     `json:"refs,omitempty"`
	Children    []FileInfo    `json:"children,omitempty"`
}

// IsDir returns true if the file is a folder.
func (f FileInfo) IsDir() bool {
	return f.Type == "FOLDER"
}

type FileRefs struct {
	Download  string `json:"download"`
	Icon      string `json:"icon"`
	Thumbnail string `json:"thumbnail"`
}
