package vm

var (
	verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enables verbose output (may impact performance)")
}
