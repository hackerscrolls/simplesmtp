package main

import (
	"io/ioutil"
	"net"
	"time"
	"flag"
	"strings"
	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"os"
	"math/rand"
	"github.com/corpix/smtpd"
	"strconv"
)

type smtpServer struct{}

var(
	save = false
	pathSave = "log/"
)

func (s *smtpServer) ServeSMTP(c net.Conn, e *smtpd.Envelope) {
	logrus.Infof(
		"Received message from %s those envelope: %s\n",
		c.RemoteAddr(),
		spew.Sdump(e),
	)

	msg, err := e.Message()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(msg.Body)
	if err != nil {
		panic(err)
	}
	logrus.Infof("Message body: %s\n", body)

	if save {
		SaveEmail(body, e.From, e.To[0])
	}
}

func SaveEmail(body []byte, fromAddr string, toAddr string ){
	rnd := rand.Intn(9999)
	t := time.Now()
	timeStamp := t.Format("2006-01-02-15_04_05")
	fromAddr = strings.ReplaceAll(fromAddr,"@","_")
	toAddr = strings.ReplaceAll(toAddr,"@","_")

	filename := fromAddr+"_to_"+toAddr+"_"+timeStamp+"_"+strconv.Itoa(rnd)+".txt"

	err := ioutil.WriteFile(pathSave+filename, body, 0644)

	if err != nil {
		logrus.Error(err)
		logrus.Println("Trying save without addrs")
		filename = "_"+timeStamp+"_"+strconv.Itoa(rnd)+".txt"
		err = ioutil.WriteFile(pathSave+filename, body, 0644)
		if err != nil {
			logrus.Error(err)
		}
	}

}

func main() {

	ipPtr := flag.String("i","0.0.0.0","IP interface")
	portPtr := flag.String("p","25","Port")
	tempFilesPtr := flag.Bool("save",false,"Save emails to file")

	flag.Parse()

	ip := *ipPtr
	port := *portPtr
	save = *tempFilesPtr

	if _, err := os.Stat(pathSave); os.IsNotExist(err) && save {
		os.Mkdir(pathSave, 774)
	}

	var err error

	c, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		panic(err)
	}

	logrus.Println("Starting SMTP server on "+port+"..")
	for {
		err = smtpd.Serve(c, &smtpServer{})
		if err != nil {
			logrus.Error(err)
			time.Sleep(1 * time.Second)
		}
	}
}
