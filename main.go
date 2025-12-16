package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Lütfen hedef url giriniz ")
		fmt.Println("Örnek olarak https://www.sibervatan.org/ veya www.sibervatan.org/ gibi ")
		return

	}

	var url string = os.Args[1]

	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	var response, error = http.Get(url)

	if error != nil {

		fmt.Printf("beklenmeyen bir sorun oluştu internetinizi veya girdiğiniz linki kontrol edin  %v", error)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {

		var context, cancel = chromedp.NewContext(context.Background())

		defer cancel() // fonskiyondan çıkmadan önce  hafızayı temizliyor http://github.com/chromedp/chromedp?tab=readme-ov-file readme kısmındada bu şekl kullanmışlar

		var html string
		var screenshot []byte

		fmt.Println("Korkama terminal donmadı biraz bekle :) ")
		time.Sleep(2 * time.Second)
		fmt.Println("İlginç bilgilerde bugün Fernerbahçe 11 yıldır şampiyon olamıyor ama her yıl şampiyon olacakmış gibi davranıyor...")
		time.Sleep(2 * time.Second)
		fmt.Println("Galatasaray 25 kez Türkiye Şampiyonu olarak süper ligdeki en çok şampiyon olan takımdır. ")

		error = chromedp.Run(
			context,
			chromedp.EmulateViewport(1920, 1080),
			chromedp.Navigate(url),
			chromedp.OuterHTML("html", &html),
			chromedp.FullScreenshot(&screenshot, 90), // 90 sayısı fotoğrafın kalitesini belirtiyor değiştrebilirsiniz
		)

		if error != nil {
			fmt.Printf("veri çekme sırasında bir sorun oluştu: %v", error)
			return
		}

		error = os.WriteFile("output.txt", []byte(html), 0644) // 0644 dosya izinleri için
		if error != nil {
			fmt.Printf("html kaydedilemedi: %v", error)
		} else {
			fmt.Println("html kaydedildi.")
		}

		error = os.WriteFile("screenshot.png", screenshot, 0644)
		if error != nil {
			fmt.Printf("ekran görüntüsü kaydedilemedi: %v", error)
		} else {
			fmt.Println("ekran görüntüsü alındı.")
		}

	} else if response.StatusCode == 404 {
		fmt.Println("aradığınız site bulunumamadı 404 hatası veriyor")

	} else {

		fmt.Println("ulaşmaya çalıştığınız site erişime engelli veya başaka bir hata veriyor ", response.StatusCode)

	}

}
