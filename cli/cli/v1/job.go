package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Retrieves job information about the printer.",

	RunE: func(_ *cobra.Command, _ []string) error {
		job, err := printer.Job().GetCurrent()
		if err != nil {
			return err
		}
		if job == nil {
			println("No job is currently running")
			return nil
		}
		return cli.Print(job)
	},
}

func init() {
	v1Cmd.AddCommand(jobCmd)

	jobCmd.AddCommand(jobStopCmd)
	jobCmd.AddCommand(jobPauseCmd)
	jobCmd.AddCommand(jobResumeCmd)
	jobCmd.AddCommand(jobContinueCmd)

	jobStopCmd.Flags().IntP("id", "i", 0, "ID of the job to stop")
	jobPauseCmd.Flags().IntP("id", "i", 0, "ID of the job to pause")
	jobResumeCmd.Flags().IntP("id", "i", 0, "ID of the job to resume")
	jobContinueCmd.Flags().IntP("id", "i", 0, "ID of the job to continue")
	_ = jobStopCmd.MarkFlagRequired("id")
	_ = jobPauseCmd.MarkFlagRequired("id")
	_ = jobResumeCmd.MarkFlagRequired("id")
	_ = jobContinueCmd.MarkFlagRequired("id")
}

var jobStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the job with the given ID",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return printer.Job().Stop(id)
	},
}

var jobPauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses the job with the given ID",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return printer.Job().Pause(id)
	},
}

var jobResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumes the job with the given ID",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return printer.Job().Resume(id)
	},
}

var jobContinueCmd = &cobra.Command{
	Use:   "continue",
	Short: "Continues the job with the given ID",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return printer.Job().Continue(id)
	},
}
