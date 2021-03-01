// Package email provides email sending functionality.
package email

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	debug     bool
	templates *template.Template
)

// Email struct holds email address and recipient name.
type Email struct {
	Name    string
	Address string
}

// NewEmail returns an email address.
func NewEmail(name string, address string) Email {
	return Email{
		Name:    name,
		Address: address,
	}
}

func parseTemplates() error {
	templates = template.New("").Funcs(fMap)
	return filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templates.ParseFiles(path)
			return err
		}
		return err
	})
}

var fMap = template.FuncMap{
	"formatAsDate":     formatAsDate,
	"formatAsDuration": formatAsDuration,
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d.%d.%d", day, month, year)
}

func formatAsDuration(t time.Time) string {
	dur := t.Sub(time.Now())
	hours := int(dur.Hours())
	mins := int(dur.Minutes())

	v := ""
	if hours != 0 {
		v += strconv.Itoa(hours) + " hours and "
	}
	v += strconv.Itoa(mins) + " minutes"
	return v
}
