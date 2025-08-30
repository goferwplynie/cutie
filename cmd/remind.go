package cmd

import (
	"fmt"
	"math/rand"

	"github.com/goferwplynie/cutie/logger"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
	"github.com/spf13/cobra"
)

var noColor bool

var projectReminders = []string{
	"Hey cutie~ your project %s hasn’t been touched in a while! Go give it some love~",
	"It’s been some time since you worked on %s~ maybe check up on it? (*^ω^*)",
	"Don’t forget %s today~ It's waiting for you!",
	"Your project %s is lonely~ a little attention will make it happy!",
	"Just work a little bit on %s cutie. It's really worth it :3",
}

var deadlineReminders = []string{
	"Haaiii cutie :33! %s deadline is in %d days!",
	"%s is due in %d days! Keep going, you’re doing great~",
	"Don’t forget! %s deadline is getting close (%d days left)~ cheer up!",
	"%s has %d days left~ let’s finish it cutie >w<",
}

func getRandomProjectReminder(projectName string) string {
	template := projectReminders[rand.Intn(len(projectReminders))]
	return fmt.Sprintf(template, projectName)
}

func getRandomDeadlineReminder(projectName string, daysLeft int) string {
	template := deadlineReminders[rand.Intn(len(deadlineReminders))]
	return fmt.Sprintf(template, projectName, daysLeft)
}

var remindCmd = &cobra.Command{
	Use:   "remind",
	Short: "show reminders for projects",
	Long:  `command will show reminders for projects that you have not worked on for some time or if the deadline is close`,
	Run: func(cmd *cobra.Command, args []string) {
		storage := projectstorage.New("")
		if err := storage.Setup(); err != nil {
			logger.Error(fmt.Sprintf("error at setup :c : %v", err))
			return
		}
		if err := storage.SyncReminders(false); err != nil {
			logger.Error(fmt.Sprintf("error while checking for reminders TwT: %v", err))
			return
		}

		reminders, err := storage.GetReminders()
		if err != nil {
			logger.Error(fmt.Sprintf("error getting reminders :c : %v", err))
			return
		}

		if len(reminders.Reminders) == 0 && len(reminders.Deadlines) == 0 {
			fmt.Println("Yay! No pending reminders. Keep up the great work!")
			return
		}

		if len(reminders.Reminders) > 0 {
			fmt.Println("Projects that need some love:")
			for _, proj := range reminders.Reminders {
				if noColor {
					fmt.Println("   " + getRandomProjectReminder(proj))
				} else {
					logger.RandomColor("   " + getRandomProjectReminder(proj))
				}
			}
		}

		if len(reminders.Deadlines) > 0 {
			fmt.Println("Upcoming Deadlines:")
			for _, deadline := range reminders.Deadlines {
				if noColor {
					fmt.Println("   " + getRandomDeadlineReminder(deadline.Name, deadline.DaysLeft))
				} else {
					logger.RandomColor("   " + getRandomDeadlineReminder(deadline.Name, deadline.DaysLeft))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(remindCmd)

	remindCmd.Flags().BoolVar(&noColor, "nc", false, "no color: not colored output :c")
}
