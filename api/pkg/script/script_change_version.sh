#!/bin/bash

# Путь к файлу, который нужно изменить
pwd=$(pwd)
file_path="generated/api.gen.go"

# Проверяем, существует ли файл
if [ ! -f "$file_path" ]; then
    echo "Файл $file_path не найден."
    exit 1
fi

# Читаем содержимое файла
content=$(cat "$file_path")

# Изменяем первую строку
new_content=$(echo "$content" | sed '1s/\/\/go:build go1.22/\/\/go:build go1.21/')

# Записываем изменённое содержимое в файл
echo "$new_content" > "$file_path"

# echo "Файл $file_path успешно изменён."