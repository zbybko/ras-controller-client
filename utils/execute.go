package utils

func ExecuteErr(command string, args ...string) error {
	_, err := Execute(command, args...)
	return err
}
