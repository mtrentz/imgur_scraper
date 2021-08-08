package main

import "fmt"

func main() {

	baseUrl := "https://i.imgur.com/"

	for i := 0; i < 50; i++ {
		// Gera o codigo do imgur
		code := ImgurCodeGenerator(6)
		// Coloca um .png pra garantir que o site vai abrir só a imagem caso existir
		imageName := code + ".png"
		// Gera o url do imgur para a imagem
		url := baseUrl + imageName
		fmt.Printf("%d - %s \n", i+1, url)
		// Salva o arquivo caso seja mesmo uma img. Mesmo se for um jpeg vai salvar
		filePath := "imgs/" + imageName
		DownloadImage(filePath, url)
	}
}
