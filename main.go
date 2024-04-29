package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/swaggo/http-swagger"
)

// User структура для пользователя
type User struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// authenticate функция для аутентификации пользователя по электронной почте
func authenticate(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Здесь вы можете добавить логику для проверки аутентификации пользователя по электронной почте
    // Например, проверка пароля и существование пользователя в базе данных

    // В качестве примера я просто верну успех и данные пользователя
    json.NewEncoder(w).Encode(user)
}

func main() {
    router := mux.NewRouter()

    // Маршрут для аутентификации пользователя
    router.HandleFunc("/authenticate", authenticate).Methods("POST")

    // Маршрут для Swagger UI
    router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

    log.Fatal(http.ListenAndServe(":8080", router))
}
