package v1

import (
	"github.com/spf13/cobra"
	"github.com/ysmilda/prusalink-go/cli/cli"
)

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Retrieves job information about the printer.",

	RunE: func(_ *cobra.Command, _ []string) error {
		job, err := conn.Job().Get()
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

	jobStopCmd.Flags().IntP("id", "i", 0, "ID of the job to stop")
	jobPauseCmd.Flags().IntP("id", "i", 0, "ID of the job to pause")
	jobResumeCmd.Flags().IntP("id", "i", 0, "ID of the job to resume")
	_ = jobStopCmd.MarkFlagRequired("id")
	_ = jobPauseCmd.MarkFlagRequired("id")
	_ = jobResumeCmd.MarkFlagRequired("id")
}

var jobStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the job with the given ID",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return conn.Job().Stop(id)
	},
}

var jobPauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses the job with the given ID",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return conn.Job().Pause(id)
	},
}

var jobResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumes the job with the given ID",
	RunE: func(cmd *cobra.Command, _ []string) error {
		id, _ := cmd.Flags().GetInt("id")
		return conn.Job().Resume(id)
	},
}
