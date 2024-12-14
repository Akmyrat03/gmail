package service

import (
	"context"
	"mail-service/models"
	"mail-service/repository"
	"net/smtp"
)

type ContactService struct {
	Repo *repository.MessageRepository
}

func NewContactService(repo *repository.MessageRepository) *ContactService {
	return &ContactService{Repo: repo}
}

// Gmail bilgileri (Uygulama Şifresi gereklidir)
const (
	smtpServer = "smtp.gmail.com"
	smtpPort   = "587"
	smtpEmail  = "akmobile.tm@gmail.com" // Kendi Gmail adresiniz
	smtpPass   = "whclvwobghfdrmqm"      // Gmail uygulama şifresi
)

func (s *ContactService) SendMessage(ctx context.Context, message models.ContactMessage) error {
	// 1. Save the message to the database
	if err := s.Repo.SaveMessage(ctx, message); err != nil {
		return err
	}

	// Gmail SMTP setup
	auth := smtp.PlainAuth("", smtpEmail, smtpPass, smtpServer)

	// Email to admin (yourself)
	adminTo := []string{smtpEmail} // Your Gmail
	adminSubject := "Yeni İletişim Mesajı"
	adminBody := "Ad: " + message.Name + "\nEmail: " + message.Email + "\nMesaj:\n" + message.Message
	adminMessage := "From: " + smtpEmail + "\n" +
		"To: " + smtpEmail + "\n" +
		"Subject: " + adminSubject + "\n\n" +
		adminBody

	// Email to the user (sender)
	userTo := []string{message.Email} // User's email
	userSubject := "Mesajınız Alındı"
	userBody := "Merhaba " + message.Name + ",\n\n" +
		"Mesajınızı aldık! İşte göndermiş olduğunuz mesaj:\n\n" +
		"Mesaj:\n" + message.Message + "\n\n" +
		"Teşekkürler,\nAkMobile Destek Ekibi"
	userMessage := "From: " + smtpEmail + "\n" +
		"To: " + message.Email + "\n" +
		"Subject: " + userSubject + "\n\n" +
		userBody

	// 2. Send the emails
	// Send email to admin
	if err := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpEmail, adminTo, []byte(adminMessage)); err != nil {
		return err
	}

	// Send email to the user
	if err := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpEmail, userTo, []byte(userMessage)); err != nil {
		return err
	}

	return nil
}
