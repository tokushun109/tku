package controllers

import (
	"api/app/controllers/utils"
	"api/app/models"
	"api/config"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/go-playground/validator.v9"
)

// お問い合わせの新規作成
func createContactHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
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
	wd, err := utils.GetWorkingDir()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// メール通知
	var message *mail.SGMailV3

	message = mail.NewV3Mail()
	from := mail.NewEmail("tocoriri", "no-reply@tocoriri.com")
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
	var tpl *template.Template

	// お問い合わせ元への通知
	p := mail.NewPersonalization()
	to := mail.NewEmail(contact.Name, contact.Email)
	p.AddTos(to)
	message.AddPersonalizations(p)

	replyMail := MailTemplate{
		Title:       "【tocoriri】お問い合わせを受け付けました",
		Name:        contact.Name,
		Company:     *contact.Company,
		PhoneNumber: *contact.PhoneNumber,
		Email:       contact.Email,
		Content:     contact.Content,
	}

	message.Subject = replyMail.Title

	// テキストパートを設定
	c := mail.NewContent("text/plain", "テストのテキストメールです。")
	message.AddContent(c)
	// HTMLパートを設定
	buffer = &bytes.Buffer{}

	tpl = template.Must(template.ParseFiles(wd + "/app/controllers/mail/contact/auto_reply.html"))
	if err = tpl.Execute(buffer, replyMail); err != nil {
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
		Title:       "【tocoriri】お問い合わせが届きました",
		Name:        contact.Name,
		Company:     *contact.Company,
		PhoneNumber: *contact.PhoneNumber,
		Email:       contact.Email,
		Content:     contact.Content,
	}

	message.Subject = adminMail.Title
	// テキストパートを設定
	ac := mail.NewContent("text/plain", "テストのテキストメールです。")
	message.AddContent(ac)
	// HTMLパートを設定
	buffer = &bytes.Buffer{}

	tpl = template.Must(template.ParseFiles(wd + "/app/controllers/mail/contact/admin.html"))
	if err = tpl.Execute(buffer, adminMail); err != nil {
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

	// lineに通知する

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
