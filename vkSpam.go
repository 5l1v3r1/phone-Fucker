package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func vkSpam(phone string) {
	token, err := ioutil.ReadFile("assets/token.txt")
	if err != nil {
		fmt.Println("Введите в файл assets/token.txt токен от VK")
		time.Sleep(1 * time.Second)
		return
	}

	tfile, err := ioutil.ReadFile("assets/spammers.txt")
	if err != nil {
		fmt.Println("Введите в файл assets/spammer.txt ID аккаунтов")
		time.Sleep(1 * time.Second)
		return
	}
	targets := strings.Split(string(tfile), "\r\n")

	for i, target := range targets {

		green := color.New(color.FgGreen)
		boldGreen := green.Add(color.Bold)
		tlen := strconv.Itoa(len(targets))
		var url = "https://api.vk.com/method/messages.send?user_id=" + target + "&message=Здравствуйте,+я+по+объявлению+Номер+для+связи:+" + phone + "&random_id=0&v=5.95&access_token=" + string(token)
		http.Get(url)
		boldGreen.Println("[ " + strconv.Itoa(i+1) + "/" + tlen + " ]" + " Отправлено ")
		time.Sleep(1 * time.Second)

	}

}
