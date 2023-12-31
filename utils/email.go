package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"os"
	"oscar/musinterest/initializers"
	"oscar/musinterest/models"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *models.User, data *EmailData) {
    	godotenv.Load(".env.local")
	_, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
		return
	}

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
		return
	}

	template.ExecuteTemplate(&body, "verficationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", "musinterest@org.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
	    log.Fatal("Could not parse smtp port", err)
	    return
	}

	fmt.Println(port)
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email", err)
		return
	}
}
