package mail

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"lmizania/config"
	"lmizania/database"
	"log"
	"net/smtp"
	"time"

	r "github.com/redis/go-redis/v9"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

type EmailVerifier struct {
	client *r.Client
}

func NewVerifier() *EmailVerifier {
	return &EmailVerifier{
		client: database.RedisClient,
	}
}

func (v *EmailVerifier) GenerateOTP() string {
	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (v *EmailVerifier) SendVerfication(userID string, to []string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	otp := v.GenerateOTP()
	message := []byte(fmt.Sprintf("Verification code is %s", otp))
	v.client.Set(context.Background(), "userOTP:"+userID, otp, time.Hour*1)
	auth := smtp.PlainAuth("", config.MAILER_EMAIL, config.MAILER_PASSWORD, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, config.MAILER_EMAIL, to, message)
	if err != nil {
		return err
	}
	return nil
}

func (v *EmailVerifier) Verify(userID string, otp string) error {
	userOTP := v.client.Get(context.Background(), "userOTP:"+userID).Val()
	log.Println(userOTP)
	if userOTP != otp {
		return errors.New("verification failed")
	}
	return nil
}
