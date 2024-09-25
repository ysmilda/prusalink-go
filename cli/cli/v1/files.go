package v1

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var (
	storage string
	path    string
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Retrieves information about the files on the printer.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		files, err := conn.Files().List(storage, path)
		if err != nil {
			return err
		}
		return cli.Print(files)
	},
}

func init() {
	v1Cmd.AddCommand(filesCmd)

	filesCmd.AddCommand(filesCreateFolderCmd)
	filesCmd.AddCommand(filesUploadCmd)
	filesCmd.AddCommand(filesStartPrintCmd)
	filesCmd.AddCommand(filesDeleteCmd)

	filesCmd.PersistentFlags().StringVar(&storage, "storage", "", "Storage medium to access")
	filesCmd.PersistentFlags().StringVar(&path, "path", "", "Path to access. Depending on the command it can link to a folder or a file") //nolint: lll
	_ = filesCmd.MarkPersistentFlagRequired("storage")
	_ = filesCmd.MarkPersistentFlagRequired("path")

	filesUploadCmd.Flags().StringP("file", "f", "", "File to transfer to the printer")
	filesUploadCmd.Flags().Bool("overwrite", false, "Overwrite the file if it already exists")
	filesUploadCmd.Flags().Bool("print", false, "Print the file after uploading")
	_ = filesUploadCmd.MarkFlagRequired("file")

	filesStartPrintCmd.Flags().Bool("deleteNonEmptyFolder", false, "Delete the folder even if it is not empty.")
}

var filesCreateFolderCmd = &cobra.Command{
	Use:   "create-folder",
	Short: "Creates a new folder on the printer.",
	Long:  "Creates a new folder on the printer. The path must point to the folder that needs to be created.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return conn.Files().CreateFolder(storage, path)
	},
}

var filesUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a new file on the printer.",
	Long: `Uploads a new file on the printer. 
The file must have a .gcode extension and the path must point to a folder.
If the file already exists on the printer, it will not be overwritten unless the --overwrite flag is set.
If the --print flag is set, the file will be printed after uploading.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		file, _ := cmd.Flags().GetString("file")
		overwrite, _ := cmd.Flags().GetBool("overwrite")
		printAfterUpload, _ := cmd.Flags().GetBool("print")
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		return conn.Files().Upload(storage, path, file, content, overwrite, printAfterUpload)
	},
}

var filesStartPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Starts a print on the printer.",
	Long:  "Starts a print on the printer. The path must point to a .gcode file.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return conn.Files().StartPrint(storage, path)
	},
}

var filesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a file or folder on the printer.",
	Long:  "Deletes a file or folder on the printer. The path must point to the file or folder that needs to be deleted.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		deleteNonEmptyFolder, _ := cmd.Flags().GetBool("deleteNonEmptyFolder")
		return conn.Files().Delete(storage, path, deleteNonEmptyFolder)
	},
}
