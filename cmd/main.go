// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"

// 	"time"

// 	"gopkg.in/telebot.v3"
// )

// // Структура для хранения данных из JSON
// type WordData struct {
// 	Synonyms []string `json:"синонимы"`
// 	Antonyms []string `json:"антонимы"`
// }

// // Глобальная переменная для словаря
// var dictionary map[string]WordData

// // Функция для загрузки JSON-файла
// func loadJSON(filename string) error {
// 	file, err := os.ReadFile(filename)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(file, &dictionary)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func main() {
// 	// Загружаем JSON-файл
// 	err := loadJSON("dictionary.json")
// 	if err != nil {
// 		log.Fatal("Ошибка загрузки JSON:", err)
// 	}

// 	botToken := "7489291057:AAFJ-tYlKMt75IwBWfGdxDkJ-hAGM8P8Cpk"

// 	// Настройки бота
// 	pref := telebot.Settings{
// 		Token:  botToken,
// 		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
// 	}

// 	bot, err := telebot.NewBot(pref)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Обработчик команды /start
// 	bot.Handle("/start", func(c telebot.Context) error {
// 		return c.Send("Привет! Отправь мне слово, и я найду его синонимы и антонимы.")
// 	})

// 	// Обработчик входящих сообщений (поиск слова)
// 	bot.Handle(telebot.OnText, func(c telebot.Context) error {
// 		word := strings.ToLower(c.Text())

// 		data, exists := dictionary[word]
// 		if !exists {
// 			return c.Send("Извините, я не нашел это слово в базе")
// 		}

// 		response := fmt.Sprintf("🔹 *Синонимы:* %s\n🔸 *Антонимы:* %s",
// 			strings.Join(data.Synonyms, ", "),
// 			strings.Join(data.Antonyms, ", "))

// 		return c.Send(response, telebot.ModeMarkdown)
// 	})

// 	// Запуск бота
// 	fmt.Println("Бот запущен...")
// 	bot.Start()
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

// Структура для хранения данных из JSON
type WordData struct {
	Synonyms []string `json:"синонимы"`
	Antonyms []string `json:"антонимы"`
}

// Глобальная переменная для словаря
var dictionary map[string]WordData

// Функция для загрузки JSON-файла
func loadJSON(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &dictionary)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Загружаем JSON-файл
	err := loadJSON("dictionary.json")
	if err != nil {
		log.Fatal("Ошибка загрузки JSON:", err)
	}

	botToken := "7489291057:AAFJ-tYlKMt75IwBWfGdxDkJ-hAGM8P8Cpk"

	// Настройки бота
	pref := telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	// Обработчик команды /start
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Привет! Отправь мне слово, и я найду его синонимы и антонимы.")
	})

	// Обработчик входящих сообщений (поиск слова)
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		word := strings.ToLower(c.Text())

		data, exists := dictionary[word]
		if !exists {
			return c.Send("Извините, я не нашел это слово в базе")
		}

		response := fmt.Sprintf("🔹 *Синонимы:* %s\n🔸 *Антонимы:* %s",
			strings.Join(data.Synonyms, ", "),
			strings.Join(data.Antonyms, ", "))

		return c.Send(response, telebot.ModeMarkdown)
	})

	// Запускаем бота в фоне
	go func() {
		fmt.Println("Бот запущен...")
		bot.Start()
	}()

	// Фиктивный сервер для Render (чтобы сервис не падал)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Дефолтный порт, если переменная не задана
	}

	fmt.Println("Фиктивный HTTP-сервер запущен на порту", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Бот работает"))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
