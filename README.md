# simplesmtp
Just a simple SMTP server, implementation of @corpix smtpd library

# Install 
go get -u github.com/corpix/smtpd
git clone https://github.com/hackerscrolls/simplesmtp
go run simplesmtp.go -save -i 0.0.0.0 -p 25

# Flags 
-save   Save all incoming emails to files. Creates log/ directory and saves all emails there
-i    Specify IP to listen
-p    Specify TCP Port to listen
