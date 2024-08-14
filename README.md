### **Golang Todo with HTMX**
I hate spending time with frontend things. I wanted to improve and broaden my knowledge of Go and HTMX by building this project without having to think about popular framework or library which specializes in frontend, like React, Vue, Svelte etc.
![demo](https://github.com/icodeerror/go-todo-htmx/blob/main/go-todo-htmx-demo.gif)

------------


This project depends on the following packages
- [pgx](https://github.com/jackc/pgx "PGX") - PostgreSQL Driver and Toolkit
- [godotenv](https://github.com/joho/godotenv "godotenv") - A Go port of Ruby's dotenv library (Loads environment variables from .env files)

##### How to run
1. Add the following keys to your `.env` file.
```
DB_USER
DB_PASSWORD
DB_HOST
DB_PORT
DB_NAME
```
2. In project directory, type the following command in the terminal to run it.
```
go run .
```


For the external CDN, it uses
- [Bulma](https://bulma.io/ "Bulma") - The Modern CSS Framework
- [FontAwesome](https://fontawesome.com/ "FontAwesome") - Icon library
- [HTMX](https://htmx.org "HTMX") - high power tools for HTML

Note that this is hobby project and doesn't include testing.
