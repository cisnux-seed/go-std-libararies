package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"name", "age", "email"})
	_ = writer.Write([]string{"fajra", "20", "fajra@gmail.com"})
	_ = writer.Write([]string{"cisnux", "18", "cisnux@gmail.com"})
	_ = writer.Write([]string{"risqulla", "21", "risqulla@gmail.com"})
	writer.Flush() // Flush writes any buffered data to the underlying io.Writer.
}
