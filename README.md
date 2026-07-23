# Morse Uploader

Веб-сервис на Go для загрузки текстовых файлов и перевода текста в азбуку Морзе.

## Запуск

```bash
git clone https://github.com/NutcrackerCom/go-morse-uploader.git
cd go-morse-uploader
go mod download
go run ./cmd
```

```text
http://localhost:8080
```

## Хэндлеры

* `GET /` — отображает страницу загрузки файла.
* `POST /upload` — принимает файл из поля `myFile`, переводит текст в азбуку Морзе и сохраняет результат в папку `uploads`.
