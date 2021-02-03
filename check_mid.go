package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/labstack/echo"

  "crypto/subtle"

  "github.com/labstack/echo/middleware"
  mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func main() {
  mailjetClient := mailjet.NewMailjetClient("420ba9aa6229abbd55b0b67e46238aec", "e2cecdede1c800924ffa9b20d93857c1")
  e := echo.New()

  e.Use(middleware.BasicAuth(func(user, pass string, c echo.Context) (bool, error) {
    // Be careful to use constant time comparison to prevent timing attacks
    if subtle.ConstantTimeCompare([]byte(user), []byte("Kushagra")) == 1 &&
      subtle.ConstantTimeCompare([]byte(pass), []byte("123456")) == 1 {
      messagesInfo := []mailjet.InfoMessagesV31{
        mailjet.InfoMessagesV31{
          From: &mailjet.RecipientV31{
            Email: "kushagrap24@gmail.com",
            Name:  "Kushagra",
          },
          To: &mailjet.RecipientsV31{
            mailjet.RecipientV31{
              Email: "kushagra.cavisson@gmail.com",
              Name:  "Kushagra parashar",
            },
          },

          Cc: &mailjet.RecipientsV31{
            mailjet.RecipientV31{
              Email: "shubhamparasar00@gmail.com",
              Name:  "Kushagra parashar",
            },
          },
          Bcc: &mailjet.RecipientsV31{
            mailjet.RecipientV31{
              Email: "shubhamparasar00@gmail.com",
              Name:  "Kushagra parashar",
            },
          },

          Subject:  "Greetings from Mailjet.",
          TextPart: "My first Mailjet email",
          HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
          CustomID: "AppGettingStartedTest",
        },
      }
      messages := mailjet.MessagesV31{Info: messagesInfo}
      res, err := mailjetClient.SendMailV31(&messages)
      if err != nil {
        log.Fatal(err)
      }
      fmt.Printf("Data: %+v\n", res)

      return true, nil
    }
    return false, nil
  }))
  // e.Use(k(a))

  e.GET("/", func(c echo.Context) error {
    fmt.Println("hey")
    return c.String(http.StatusOK, "My name is kushagra Parashar, Hope this code is working fine!!!")
  })

  e.Logger.Fatal(e.Start(":3000"))
}
