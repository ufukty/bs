package colors

import "fmt"

func Blue(s any) string    { return fmt.Sprintf("\033[34m%v\033[0m", s) }
func Cyan(s any) string    { return fmt.Sprintf("\033[36m%v\033[0m", s) }
func Green(s any) string   { return fmt.Sprintf("\033[32m%v\033[0m", s) }
func Magenta(s any) string { return fmt.Sprintf("\033[35m%v\033[0m", s) }
func Red(s any) string     { return fmt.Sprintf("\033[31m%v\033[0m", s) }
func Yellow(s any) string  { return fmt.Sprintf("\033[33m%v\033[0m", s) }

func BlueL(s any) string    { return fmt.Sprintf("\033[94m%v\033[0m", s) }
func CyanL(s any) string    { return fmt.Sprintf("\033[96m%v\033[0m", s) }
func GreenL(s any) string   { return fmt.Sprintf("\033[92m%v\033[0m", s) }
func MagentaL(s any) string { return fmt.Sprintf("\033[95m%v\033[0m", s) }
func RedL(s any) string     { return fmt.Sprintf("\033[91m%v\033[0m", s) }
func YellowL(s any) string  { return fmt.Sprintf("\033[93m%v\033[0m", s) }
