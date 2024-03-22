//nolint:lll
package v1

import "github.com/ysmilda/prusalink-go/pkg/types"

type Version struct {
	API          string        `json:"api"`
	Version      string        `json:"version"`
	Printer      string        `json:"printer"`
	Text         string        `json:"text"`
	Firmware     string        `json:"firmware"`
	SDK          *string       `json:"sdk,omitempty"`
	Capabilities *Capabilities `json:"capabilities,omitempty"` // Additional capabilities the printer has. The object is expected to be extended in the future with more capabilities. The absence of a capability in the object, or the complete absence of the object means the printer doesn't support such capability (probably doesn't even know such capability might exist).
}

type Capabilities struct {
	UploadByPUT *bool                  `json:"upload_by_put,omitempty"` // The printer supports uploading GCodes by the PUT method (as described in this document). It is capable of doing the PUT and HEAD to /api/v1/files/{storage}/{path} and it is capable of answering the /api/v1/storage endpoint. In absence of this capability, client MAY opt to try the legacy "octoprint" POST to /api/files/{storage}.
	Extras      map[string]interface{} `json:"-,omitempty"`             // Any extra capabilities that were not documented at the time of development.
}

type StorageInfo struct {
	StorageList []Storage `json:"storage_list"`
}

type Storage struct {
	Name        *string `json:"name,omitempty"`         // Name of the storage, based on selected language
	Type        string  `json:"type"`                   // Storage source
	Path        string  `json:"path"`                   // Path to storage (not display path)
	PrintFiles  *int    `json:"print_files,omitempty"`  // Size of all print files in bytes
	SystemFiles *int    `json:"system_files,omitempty"` // Size of all system files in bytes
	FreeSpace   *int    `json:"free_space,omitempty"`   // System free space in bytes
	TotalSpace  *int    `json:"total_space,omitempty"`  // System total space in bytes
	Available   bool    `json:"available"`              // Whether the storage is available or not
	ReadOnly    *bool   `json:"read_only,omitempty"`    // Whether the storage is read only
}

type Camera struct {
	CameraID  string               `json:"camera_id"`
	Config    CameraHardwareConfig `json:"config"`
	Connected bool                 `json:"connected"` // Camera is successfully connected to PrusaLink
	Detected  bool                 `json:"detected"`  // Camera is detected by PrusaLink, but not saved yet
	Stored    bool                 `json:"stored"`    // Camera configuration is saved in PrusaLink
	Linked    bool                 `json:"linked"`    // Camera is linked to PrusaConnect
}

type CameraHardwareConfig struct {
	Path       string `json:"path"`
	Name       string `json:"name"`
	Driver     string `json:"driver"`
	Resolution string `json:"resolution"`
}

// Camera configuration.
type CameraConfig struct {
	Name                 string       `json:"name"`
	TriggerScheme        string       `json:"trigger_scheme"` // When the snapshot is taken
	AvailableResolutions []Resolution `json:"available_resolutions"`
	Resolution           Resolution   `json:"resolution"`
	Focus                float64      `json:"focus"` // Focus of the camera (0.0 - 1.0)
	Capabilities         []string     `json:"capabilities"`
}

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Camera configuration to set.
type CameraConfigSet struct {
	Name          string     `json:"name"`           // Name of the camera
	TriggerScheme string     `json:"trigger_scheme"` // When the snapshot is taken
	Resolution    Resolution `json:"resolution"`
	Rotation      int        `json:"rotation"` // Current rotation of the output image
	Focus         float64    `json:"focus"`    // Focus of the camera (0.0 - 1.0)
	Exposure      float64    `json:"exposure"`
	SendToConnect bool       `json:"send_to_connect"`
}

// Basic file info object, common for all files.
type GenericFileInfo struct {
	Name        string     `json:"name"` // Short Filename
	ReadOnly    bool       `json:"read_only"`
	Size        *int       `json:"size,omitempty"`         // Available for files only, not for folders
	Type        string     `json:"type"`                   // File could be print file, firmware file, other (e.g. configuration) file, or folder
	MTimestamp  types.Time `json:"m_timestamp"`            // Timestamp in seconds
	DisplayName *string    `json:"display_name,omitempty"` // Long Filename
}

// Other, not specified files info.
type FileInfo struct {
	GenericFileInfo
	Refs Refs `json:"refs"`
}

type Refs struct {
	Download string `json:"download"`
}

// Simplified print file info within the folder's children.
type PrintFileInfoBasic struct {
	GenericFileInfo
	Refs PrintFileRefs `json:"refs"`
}

// Reference links for file thumbnail, icon and download.
type PrintFileRefs struct {
	Download  string `json:"download"`
	Icon      string `json:"icon"`
	Thumbnail string `json:"thumbnail"`
}

// Print file metadata parsed from G-code or SL1, all data are optional.
type PrintFileMetadata struct {
	BedTemperature          *int            `json:"bed_temperature,omitempty"`          // Degrees Celsius
	BedTemperaturePerTool   *[]int          `json:"bed_temperature_per_tool,omitempty"` // Degrees Celsius. unsure about json tag
	Temperature             *int            `json:"temperature,omitempty"`              // Nozzle temperature, degrees Celsius
	TemperaturePerTool      *[]int          `json:"temperature_per_tool,omitempty"`     // Nozzle temperature per tool, degrees Celsius. Unsure about json tag
	BrimWidth               *int            `json:"brim_width,omitempty"`               // Milimeters
	EstimatedPrintingTime   *string         `json:"estimated_printing_time,omitempty"`  // Unsure about json tag
	EstimatedPrintTime      *types.Duration `json:"estimated_print_time,omitempty"`
	FadedLayers             *int            `json:"faded_layers,omitempty"`
	FilamentCost            *float64        `json:"filament_cost,omitempty"`              // Unsure about json tag
	FilamentCostPerTool     *[]float64      `json:"filament_cost_per_tool,omitempty"`     // Unsure about json tag
	FilamentUsedCm3         *float64        `json:"filament_used_cm3,omitempty"`          // Unsure about json tag
	FilamentUsedCm3PerTool  *[]float64      `json:"filament_used_cm3_per_tool,omitempty"` // Unsure about json tag
	FilamentUsedGram        *float64        `json:"filament_used_g,omitempty"`            // Unsure about json tag
	FilamentUsedGramPerTool *[]float64      `json:"filament_used_g_per_tool,omitempty"`   // Unsure about json tag
	FilamentUsedMm          *float64        `json:"filament_used_mm,omitempty"`           // Unsure about json tag
	FilamentUsedMmPerTool   *[]float64      `json:"filament_used_mm_per_tool,omitempty"`  // Unsure about json tag
	FilamentType            *string         `json:"filament_type,omitempty"`
	FilamentTypePerTool     *[]string       `json:"filament_type_per_tool,omitempty"` // Unsure about json tag
	FillDensity             *string         `json:"fill_density,omitempty"`           // Percents
	InitialExposureTime     *types.Duration `json:"initial_exposure_time,omitempty"`
	LayerHeight             *float64        `json:"layer_height,omitempty"`
	MaterialName            *string         `json:"material_name,omitempty"`
	ExposureTime            *types.Duration `json:"exposure_time,omitempty"`
	MaxExposureTime         *types.Duration `json:"max_exposure_time,omitempty"`
	MaxInitialExposureTime  *types.Duration `json:"max_initial_exposure_time,omitempty"`
	MinExposureTime         *types.Duration `json:"min_exposure_time,omitempty"`
	MinInitialExposureTime  *types.Duration `json:"min_initial_exposure_time,omitempty"`
	NozzleDiameter          *float64        `json:"nozzle_diameter,omitempty"`          // Milimeters
	NozzleDiameterPerTool   *[]float64      `json:"nozzle_diameter_per_tool,omitempty"` // Milimeters. Unsure about json tag
	NormalPercentPresent    *bool           `json:"normal_percent_present,omitempty"`
	NormalLeftPresent       *bool           `json:"normal_left_present,omitempty"`
	QuietPercentPresent     *bool           `json:"quiet_percent_present,omitempty"`
	QuietLeftPresent        *bool           `json:"quiet_left_present,omitempty"`
	LayerInfoPresent        *bool           `json:"layer_info_present,omitempty"`
	MaxLayerZ               *float64        `json:"max_layer_z,omitempty"`
	PrintTime               *types.Duration `json:"print_time,omitempty"`
	PrinterModel            *string         `json:"printer_model,omitempty"`
	SupportMaterial         *string         `json:"support_material,omitempty"`
	Ironing                 *int            `json:"ironing,omitempty"`
	RequiredResinMl         *float64        `json:"required_resin_ml,omitempty"`
	Profile                 *string         `json:"profile,omitempty"`
}

// Full print file info from the file's detail.
type PrintFileInfo struct {
	PrintFileInfoBasic
	Meta PrintFileMetadata `json:"meta"`
}

// Simplified firmware file info within the folder's children.
type FirmwareFileInfoBasic struct {
	GenericFileInfo
	Refs Refs `json:"refs"`
}

// Full firmware file info from the file's detail.
type FirmwareFileInfo struct {
	FirmwareFileInfoBasic
	Meta FirmwareMeta `json:"meta"`
}

type FirmwareMeta struct {
	Version        string `json:"version"` // Firmware version in text format
	PrinterType    int    `json:"printer_type"`
	PrinterVersion int    `json:"printer_version"`
}

// Info about the folder and its content, except nested children.
type FolderInfo struct {
	GenericFileInfo
	Children []interface{} `json:"children"` // TODO: Implement OneOf (FileInfo, PrintFileInfoBasic, FirmwareFileInfoBasic, FolderInfo)
}

type Info struct {
	MMU               *bool    `json:"mmu,omitempty"`
	Name              *string  `json:"name,omitempty"`
	Location          *string  `json:"location,omitempty"`
	FarmMode          *bool    `json:"farm_mode,omitempty"`
	NozzleDiameter    *float64 `json:"nozzle_diameter,omitempty"`
	MinExtrusionTemp  *int     `json:"min_extrusion_temp,omitempty"`
	Serial            *string  `json:"serial,omitempty"`
	SDReady           *bool    `json:"sd_ready,omitempty"`
	ActiveCamera      *bool    `json:"active_camera,omitempty"`
	Hostname          *string  `json:"hostname,omitempty"`
	Port              *string  `json:"port,omitempty"`
	NetworkErrorChime *bool    `json:"network_error_chime,omitempty"`
}

type Status struct {
	Job      *StatusJob      `json:"job,omitempty"`
	Printer  StatusPrinter   `json:"printer"`
	Transfer *StatusTransfer `json:"transfer,omitempty"`
	Storage  *StatusStorage  `json:"storage,omitempty"`
	Camera   *StatusCamera   `json:"camera,omitempty"`
}

// Telemetry info about current job, all values are optional.
type StatusJob struct {
	ID            *int            `json:"id,omitempty"`
	Progress      *float64        `json:"progress,omitempty"` // Percents
	TimeRemaining *types.Duration `json:"time_remaining,omitempty"`
	TimePrinting  *types.Duration `json:"time_printing,omitempty"`
}

// Telemetry info about printer.
type StatusPrinter struct {
	State         string         `json:"state"`
	TempNozzle    *float64       `json:"temp_nozzle,omitempty"`
	TargetNozzle  *float64       `json:"target_nozzle,omitempty"`
	TempBed       *float64       `json:"temp_bed,omitempty"`
	TargetBed     *float64       `json:"target_bed,omitempty"`
	AxisX         *float64       `json:"axis_x,omitempty"` // Available only when printer is not moving
	AxisY         *float64       `json:"axis_y,omitempty"` // Available only when printer is not moving
	AxisZ         *float64       `json:"axis_z,omitempty"`
	Flow          *int           `json:"flow,omitempty"`
	Speed         *int           `json:"speed,omitempty"`
	FanHotend     *int           `json:"fan_hotend,omitempty"`
	FanPrint      *int           `json:"fan_print,omitempty"`
	StatusPrinter *StatusMessage `json:"status_printer,omitempty"`
	StatusConnect *StatusMessage `json:"status_connect,omitempty"`
}

type StatusMessage struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// Telemetry info about current transfer status.
type StatusTransfer struct {
	ID               int      `json:"id"`
	TimeTransferring int      `json:"time_transferring"`
	Progress         *float64 `json:"progress,omitempty"` // Percents
	DataTransferred  *int     `json:"data_transferred,omitempty"`
}

// Telemetry info about current storage status.
type StatusStorage struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	ReadOnly  bool   `json:"read_only"`
	FreeSpace *int   `json:"free_space,omitempty"`
}

// Telemetry info about default working camera, if available.
type StatusCamera struct {
	ID string `json:"id"`
}

// PrusaLink package version available to update.
type PrusaLinkPackage struct {
	NewVersion string `json:"new_version"` // Package version available for update
}

type Transfer struct {
	Type             string          `json:"type"`
	DisplayName      string          `json:"display_name"` // Long Filename
	Path             string          `json:"path"`
	URL              *string         `json:"url,omitempty"`
	Size             *string         `json:"size,omitempty"` // Bytes
	Progress         float64         `json:"progress"`       // Percents
	Transferred      int             `json:"transferred"`    // Transfered data in bytes
	TimeRemaining    *types.Duration `json:"time_remaining,omitempty"`
	TimeTransferring types.Duration  `json:"time_transferring"`
	ToPrint          bool            `json:"to_print"` // Whether or not print after finishing transfer (upload)
}

type JobSerialPrint struct {
	SerialPrint bool `json:"serial_print"` // Whether the printer is printing from the serial line
}

type JobFilePrint struct {
	File File `json:"file"`
}

type File struct {
	Name        string             `json:"name"`                   // Short Filename
	DisplayName *string            `json:"display_name,omitempty"` // Long Filename
	Path        string             `json:"path"`
	DisplayPath *string            `json:"display_path,omitempty"`
	Size        *int               `json:"size,omitempty"` // Bytes
	MTimestamp  types.Time         `json:"m_timestamp"`
	Meta        *PrintFileMetadata `json:"meta,omitempty"`
	Refs        *PrintFileRefs     `json:"refs,omitempty"`
}

type Job struct {
	ID                  int             `json:"id"`
	State               string          `json:"state"`
	Progress            float64         `json:"progress"` // Percents
	TimeRemaining       *types.Duration `json:"time_remaining,omitempty"`
	TimePrinting        types.Duration  `json:"time_printing"`
	InaccurateEstimates *bool           `json:"inaccurate_estimates,omitempty"` //  Whether the time estimates are accurate or inaccurate
	*JobSerialPrint     `json:",omitempty"`
	*JobFilePrint       `json:",omitempty"`
}

type Error struct {
	Code  string `json:"code"`  //  Prusa error code
	Title string `json:"title"` // Prusa error title
	Text  string `json:"text"`  // Prusa error text
	URL   string `json:"url"`   // Link to the Prusa help page
}
