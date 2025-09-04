package controllers

import (
	"api/app/models"
	"api/config"
	"api/utils"
	"bytes"
	"encoding/json"
	"fmt"
	html_tmpl "html/template"
	"io"
	"log"
	"net/http"
	text_tmpl "text/template"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/go-playground/validator.v9"
)

// お問い合わせ一覧を取得
func getAllContactListHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	contactList := models.GetAllContactList()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(contactList); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// お問い合わせの新規作成
func createContactHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	var contact models.Contact
	if err := json.Unmarshal(reqBody, &contact); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// validationの確認
	validate := validator.New()
	if err := validate.Struct(contact); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
		return
	}

	// modelの呼び出し
	err = models.InsertContact(&contact)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	if config.Config.Env != "local" {
		// 通知処理
		var message *mail.SGMailV3

		wd, err := utils.GetWorkingDir()
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}

		message = mail.NewV3Mail()
		from := mail.NewEmail("とこりり", "no-reply@tocoriri.com")
		message.SetFrom(from)

		type MailTemplate struct {
			Title         string
			TopMessage    string
			Name          string
			Company       string
			PhoneNumber   string
			Email         string
			Content       string
			BottomMessage string
		}

		var buffer *bytes.Buffer
		var text *text_tmpl.Template
		var html *html_tmpl.Template

		// お問い合わせ元への通知
		p := mail.NewPersonalization()
		to := mail.NewEmail(contact.Name, contact.Email)
		p.AddTos(to)
		message.AddPersonalizations(p)

		replyMail := MailTemplate{
			Title:       "【とこりり】お問い合わせを受け付けました",
			Name:        contact.Name,
			Company:     *contact.Company,
			PhoneNumber: *contact.PhoneNumber,
			Email:       contact.Email,
			Content:     contact.Content,
		}

		message.Subject = replyMail.Title

		// テキストパートを設定
		text = text_tmpl.Must(text_tmpl.ParseFiles(wd + "/app/controllers/mail/contact/auto_reply.txt"))
		buffer = &bytes.Buffer{}
		if err = text.Execute(buffer, replyMail); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		}

		c := mail.NewContent("text/plain", buffer.String())
		message.AddContent(c)

		// HTMLパートを設定
		html = html_tmpl.Must(html_tmpl.ParseFiles(wd + "/app/controllers/mail/contact/auto_reply.html"))
		buffer = &bytes.Buffer{}
		if err = html.Execute(buffer, replyMail); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		}

		c = mail.NewContent("text/html", buffer.String())
		message.AddContent(c)

		client := sendgrid.NewSendClient(config.Config.SendGridApiKey)
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}

		// 管理者への通知
		message = mail.NewV3Mail()
		message.SetFrom(from)

		adminUsers := models.GetAllUsers()
		for _, u := range adminUsers {
			ap := mail.NewPersonalization()
			to := mail.NewEmail(u.Name, u.Email)
			ap.AddTos(to)
			message.AddPersonalizations(ap)
		}

		adminMail := MailTemplate{
			Title:       "【とこりり】お問い合わせが届きました",
			Name:        contact.Name,
			Company:     *contact.Company,
			PhoneNumber: *contact.PhoneNumber,
			Email:       contact.Email,
			Content:     contact.Content,
		}

		message.Subject = adminMail.Title

		// テキストパートを設定
		text = text_tmpl.Must(text_tmpl.ParseFiles(wd + "/app/controllers/mail/contact/admin.txt"))
		buffer = &bytes.Buffer{}
		if err = text.Execute(buffer, replyMail); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		}
		msg := buffer.String()

		ac := mail.NewContent("text/plain", msg)
		message.AddContent(ac)

		// HTMLパートを設定
		html = html_tmpl.Must(html_tmpl.ParseFiles(wd + "/app/controllers/mail/contact/admin.html"))
		buffer = &bytes.Buffer{}
		if err = html.Execute(buffer, adminMail); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		}

		ac = mail.NewContent("text/html", buffer.String())
		message.AddContent(ac)

		response, err = client.Send(message)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	}

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
