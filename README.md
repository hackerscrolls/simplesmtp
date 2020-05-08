# simplesmtp
Just a simple SMTP server, implementation of @corpix smtpd library

Original library:

https://github.com/corpix/smtpd



Original example:

https://github.com/corpix/smtpd/blob/master/examples/logging-smtpd/main.go

# Install and Run
```
go get -u github.com/corpix/smtpd
git clone https://github.com/hackerscrolls/simplesmtp
cd simplesmtp
go run simplesmtp.go -save -i 0.0.0.0 -p 25
```

# Flags 
```
-save   Save all incoming emails to files. Creates log/ directory and saves all emails there
-i      Specify IP to listen
-p      Specify TCP Port to listen
```
