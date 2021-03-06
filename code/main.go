package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Структура значка профиля
type Badge struct {
	AwardedAt   string `json:"awarded_at"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

// Структура для хранения полной информации о пользователе
type UserInfo struct {
	Username                 string  `json:"username"`
	Names                    string  `json:"previous_usernames"`
	Badges                   []Badge `json:"badges"`
	AvatarUrl                string  `json:"avatar_url"`
	UserID                   string  `json:"id"`
	CountryCode              string  `json:"country_code"`
	GlobalRank               string  `json:"global_rank"`
	CountryRank              string  `json:"country_rank"`
	PP                       string  `json:"pp"`
	PlayTime                 string  `json:"play_time"`
	PlayTimeSeconds          string  `json:"play_time_seconds"`
	SSH                      string  `json:"ssh"`
	SS                       string  `json:"ss"`
	SH                       string  `json:"sh"`
	S                        string  `json:"s"`
	A                        string  `json:"a"`
	RankedScore              string  `json:"ranked_score"`
	Accuracy                 string  `json:"accuracy"`
	PlayCount                string  `json:"play_count"`
	TotalScore               string  `json:"total_score"`
	TotalHits                string  `json:"total_hits"`
	MaximumCombo             string  `json:"maximum_combo"`
	Replays                  string  `json:"replays"`
	Level                    string  `json:"level"`
	SupportLvl               string  `json:"support_level"`
	FollowerCount            string  `json:"follower_count"`
	DefaultGroup             string  `json:"default_group"`
	IsOnline                 string  `json:"is_online"`
	IsActive                 string  `json:"is_active"`
	IsAdmin                  string  `json:"is_admin"`
	IsModerator              string  `json:"is_moderator"`
	IsNat                    string  `json:"is_nat"`
	IsGmt                    string  `json:"is_gmt"`
	IsBng                    string  `json:"is_bng"`
	IsBot                    string  `json:"is_bot"`
	IsSilenced               string  `json:"is_silenced"`
	IsDeleted                string  `json:"is_deleted"`
	IsRestricted             string  `json:"is_restricted"`
	IsLimitedBan             string  `json:"is_limited_bn"`
	IsFullBan                string  `json:"is_full_bn"`
	IsSupporter              string  `json:"is_supporter"`
	LastVisit                string  `json:"last_visit"`
	ProfileColor             string  `json:"profile_color"`
	RankedBeatmapsetCount    string  `json:"ranked_beatmapset_count"`
	PendingBeatmapsetCount   string  `json:"pending_beatmapset_count"`
	PmFriendsOnly            string  `json:"pm_friends_only"`
	GraveyardBeatmapsetCount string  `json:"graveyard_beatmapset_count"`
	BeatmapPlaycountsCount   string  `json:"beatmap_playcounts_count"`
	CommentsCount            string  `json:"comments_count"`
	FavoriteBeatmapsetCount  string  `json:"favorite_beatmapset_count"`
	GuestBeatmapsetCount     string  `json:"guest_beatmapset_count"`
	BestBeatMap              beatMap `json:"best_beat_map"`
}

// Структура для хранения информации о мапе
type beatMap struct {
	Title            string   `json:"title"`
	DifficultyRating string   `json:"difficulty_rating"`
	Id               string   `json:"id"`
	BuildId          string   `json:"build_id"`
	Statistics       string   `json:"statistics"`
	Rank             string   `json:"rank"`
	Mods             []string `json:"mods"`
	EndedAt          string   `json:"ended_at"`
	StartedAt        string   `json:"started_at"`
	Accuracy         string   `json:"accuracy"`
	MaximumCombo     string   `json:"maximum_combo"`
	PP               string   `json:"pp"`
	Passed           string   `json:"passed"`
	TotalScore       string   `json:"total_score"`
	LegacyPerfect    string   `json:"legacy_perfect"`
	Replay           string   `json:"replay"`
	Mode             string   `json:"mode"`
	Status           string   `json:"status"`
	TotalLength      string   `json:"total_length"`
	Ar               string   `json:"ar"`
	Bpm              string   `json:"bpm"`
	Convert          string   `json:"convert"`
	CountCircles     string   `json:"count_circles"`
	CountSliders     string   `json:"count_sliders"`
	CountSpinners    string   `json:"count_spinners"`
	Cs               string   `json:"cs"`
	DeletedAt        string   `json:"deleted_at"`
	Drain            string   `json:"drain"`
	HitLength        string   `json:"hit_length"`
	IsScoreable      string   `json:"is_scoreable"`
	LastUpdated      string   `json:"last_updated"`
	ModeInt          string   `json:"mode_int"`
	PassCount        string   `json:"pass_count"`
	PlayCount        string   `json:"play_count"`
	Ranked           string   `json:"ranked"`
	Url              string   `json:"url"`
	Checksum         string   `json:"checksum"`
	Creator          string   `json:"creator"`
	FavoriteCount    string   `json:"favorite_count"`
	Hype             string   `json:"hype"`
	Nsfw             string   `json:"nsfw"`
	Offset           string   `json:"offset"`
	Spotlight        string   `json:"spotlight"`
	RulesetId        string   `json:"ruleset_id"`
}

// Структура для проверки статуса пользователя
type OnlineInfo struct {
	Status string `json:"is_online"`
}

// Функция поиска. Возвращает искомое значение и индекс последнего символа
func findWithIndex(str, subStr, stopChar string, start int) (string, int) {

	// Обрезка левой границы поиска
	str = str[start:]

	// Проверка на существование нужной строки
	if strings.Contains(str, subStr) {

		// Поиск индекса начала нужной строки
		left := strings.Index(str, subStr) + len(subStr)

		// Поиск правой границы
		right := left + strings.Index(str[left:], stopChar)

		// Обрезка и вывод результата
		return str[left:right], right + start
	}

	return "", 0
}

// Облегчённая функция поиска. Возвращает только искомое значение
func find(str, subStr, stopChar string) string {

	// Проверка на существование нужной строки
	if strings.Contains(str, subStr) {

		// Обрезка левой части
		str = str[strings.Index(str, subStr)+len(subStr):]

		// Обрезка правой части и вывод результата
		return str[:strings.Index(str, stopChar)]
	}

	return ""
}

// Функция получения информации о пользователе
func getUserInfo(id, mode string) UserInfo {

	// Если пользователь не ввёл id, по умолчанию ставит мой id
	if id == "" {
		id = "29829158"
	}

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id + "/" + mode)
	if err != nil {
		return UserInfo{}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// HTML полученной страницы в формате string
	pageStr := string(body)

	// Проверка на страницу пользователя
	if !strings.Contains(pageStr, "js-react--profile") {
		return UserInfo{}
	}

	// Обрезка юзелесс части html"ки
	pageStr = strings.ReplaceAll(pageStr[strings.Index(pageStr, "current_mode"):], "&quot;", " ")

	// Сохранение html"ки в файл sample.html (для тестов)
	/*
		if err := os.WriteFile("sample.html", []byte(pageStr), 0666); err != nil {
			log.Fatal(err)
		}
	*/

	// Структура, которую будет возвращать функция
	result := UserInfo{
		UserID: id,
	}

	left := 0

	/* -----------------------------------------------------------
	# Далее происходит заполнение полей функцией find			 #
	# после каждого поиска тело сайта обрезается для оптимизации #
	------------------------------------------------------------*/

	//--------------------------- Лучшая мапа ------------------------------

	result.BestBeatMap.Accuracy = find(pageStr, "accuracy :", ",")
	result.BestBeatMap.Id = find(pageStr, "beatmap_id :", ",")
	result.BestBeatMap.BuildId = find(pageStr, "build_id :", ",")
	result.BestBeatMap.EndedAt = find(pageStr, "ended_at : ", " ")
	result.BestBeatMap.MaximumCombo, left = findWithIndex(pageStr, "max_combo :", ",", 0)
	pageStr = pageStr[left:]

	// Цикл для обработки модов
	for c := 0; pageStr[c] != ']'; c++ {
		if pageStr[c:c+10] == "acronym : " {
			result.BestBeatMap.Mods = append(result.BestBeatMap.Mods, pageStr[c+10:c+12])
		}
	}

	result.BestBeatMap.Passed = find(pageStr, "passed :", ",")
	result.BestBeatMap.StartedAt = find(pageStr, "started_at :", ",")
	result.BestBeatMap.Statistics = find(pageStr, "statistics :{ ", "}")
	result.BestBeatMap.Rank = find(pageStr, "rank : ", " ")
	result.BestBeatMap.RulesetId = find(pageStr, "ruleset_id :", ",")
	result.BestBeatMap.TotalScore = find(pageStr, "total_score :", ",")
	result.BestBeatMap.LegacyPerfect = find(pageStr, "legacy_perfect :", ",")
	result.BestBeatMap.PP = find(pageStr, "pp :", ",")
	result.BestBeatMap.Replay = find(pageStr, "replay :", ",")
	result.BestBeatMap.DifficultyRating = find(pageStr, "difficulty_rating :", ",")
	result.BestBeatMap.Mode = find(pageStr, "mode : ", " ")
	result.BestBeatMap.Status = find(pageStr, "status : ", " ")
	result.BestBeatMap.TotalLength = find(pageStr, "total_length :", ",")
	result.BestBeatMap.Ar = find(pageStr, "ar :", ",")
	result.BestBeatMap.Bpm = find(pageStr, "bpm :", ",")
	result.BestBeatMap.Convert = find(pageStr, "convert :", ",")
	result.BestBeatMap.CountCircles = find(pageStr, "count_circles :", ",")
	result.BestBeatMap.CountSliders = find(pageStr, "count_sliders :", ",")
	result.BestBeatMap.CountSpinners = find(pageStr, "count_spinners :", ",")
	result.BestBeatMap.Cs = find(pageStr, " cs :", ",")
	result.BestBeatMap.DeletedAt = find(pageStr, "deleted_at :", ",")
	result.BestBeatMap.Drain = find(pageStr, "drain :", ",")
	result.BestBeatMap.HitLength = find(pageStr, "hit_length :", ",")
	result.BestBeatMap.IsScoreable = find(pageStr, "is_scoreable :", ",")
	result.BestBeatMap.LastUpdated = find(pageStr, "last_updated : ", " ")
	result.BestBeatMap.ModeInt = find(pageStr, "mode_int :", ",")
	result.BestBeatMap.PassCount = find(pageStr, "passcount :", ",")
	result.BestBeatMap.PlayCount = find(pageStr, "playcount :", ",")
	result.BestBeatMap.Ranked = find(pageStr, "ranked :", ",")
	result.BestBeatMap.Url = find(pageStr, "url : ", " ")
	result.BestBeatMap.Url = strings.ReplaceAll(result.BestBeatMap.Url, "\\", "")
	result.BestBeatMap.Checksum = find(pageStr, "checksum : ", " ")
	result.BestBeatMap.Creator, left = findWithIndex(pageStr, "creator : ", " ", 0)
	pageStr = pageStr[left:]

	result.BestBeatMap.FavoriteCount = find(pageStr, "favourite_count :", ",")
	result.BestBeatMap.Hype = find(pageStr, "hype :", ",")
	result.BestBeatMap.Nsfw = find(pageStr, "nsfw :", ",")
	result.BestBeatMap.Offset = find(pageStr, "offset :", ",")
	result.BestBeatMap.Spotlight = find(pageStr, "spotlight :", ",")
	result.BestBeatMap.Title = find(pageStr, "title : ", ",")

	//--------------------------- Статистика игрока ------------------------------

	// В последний раз был в сети
	result.LastVisit, left = findWithIndex(pageStr, "last_visit : ", " ", 0)

	// Сообщения только от друзей
	result.PmFriendsOnly, left = findWithIndex(pageStr, "pm_friends_only :", ",", left)

	// Ссылка на аватар
	result.AvatarUrl, left = findWithIndex(pageStr, "avatar_url : ", " ", left)
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")

	// Код страны
	result.CountryCode, left = findWithIndex(pageStr, "country_code : ", " ", left)

	// Группа
	result.DefaultGroup, left = findWithIndex(pageStr, "default_group : ", " ", left)

	// Активность
	result.IsActive, left = findWithIndex(pageStr, "is_active :", ",", left)

	// Бот
	result.IsBot, left = findWithIndex(pageStr, "is_bot :", ",", left)

	// Удалённый профиль
	result.IsDeleted, left = findWithIndex(pageStr, "is_deleted :", ",", left)

	// Статус в сети
	result.IsOnline, left = findWithIndex(pageStr, "is_online :", ",", left)

	// Подписка
	result.IsSupporter, left = findWithIndex(pageStr, "is_supporter :", ",", left)

	// Цвет профиля
	result.ProfileColor, left = findWithIndex(pageStr, "profile_colour :", ",", left)

	// Юзернейм
	result.Username, left = findWithIndex(pageStr, "username : ", " ", left)

	// Администрация
	result.IsAdmin, left = findWithIndex(pageStr, "is_admin :", ",", left)

	// Номинатор
	result.IsBng, left = findWithIndex(pageStr, "is_bng :", ",", left)

	// Вечный бан
	result.IsFullBan, left = findWithIndex(pageStr, "is_full_bn :", ",", left)

	// Команда глобальной модерации
	result.IsGmt, left = findWithIndex(pageStr, "is_gmt :", ",", left)

	// Временный бан
	result.IsLimitedBan, left = findWithIndex(pageStr, "is_limited_bn :", ",", left)

	// Модератор
	result.IsModerator, left = findWithIndex(pageStr, "is_moderator :", ",", left)

	// Команда оценки номинаций
	result.IsNat, left = findWithIndex(pageStr, "is_nat :", ",", left)

	// Ограничение
	result.IsRestricted, left = findWithIndex(pageStr, "is_restricted :", ",", left)

	// Немота
	result.IsSilenced, left = findWithIndex(pageStr, "is_silenced :", ",", left)

	// Значки
	for c := strings.Index(pageStr, "badges :["); pageStr[c] != ']'; c++ {
		if pageStr[c:c+13] == "awarded_at : " {
			result.Badges = append(result.Badges, Badge{
				AwardedAt:   find(pageStr[c:], "awarded_at : ", " "),
				Description: find(pageStr[c:], "description : ", ","),
				ImageUrl:    strings.ReplaceAll(find(pageStr[c:], "image_url : ", " "), "\\", ""),
			})
		}
	}

	// Количество игр карт
	result.BeatmapPlaycountsCount, left = findWithIndex(pageStr, "beatmap_playcounts_count :", ",", left)

	// Количество комментариев
	result.CommentsCount, left = findWithIndex(pageStr, "comments_count :", ",", left)

	// Количество любимых карт
	result.FavoriteBeatmapsetCount, left = findWithIndex(pageStr, "favourite_beatmapset_count :", ",", left)

	// Подписчики
	result.FollowerCount, left = findWithIndex(pageStr, "follower_count :", ",", left)

	// Заброшенные карты
	result.GraveyardBeatmapsetCount, left = findWithIndex(pageStr, "graveyard_beatmapset_count :", ",", left)

	// Карты с гостевым участием
	result.GuestBeatmapsetCount, left = findWithIndex(pageStr, "guest_beatmapset_count :", ",", left)

	// Карты на рассмотрении
	result.PendingBeatmapsetCount, left = findWithIndex(pageStr, "pending_beatmapset_count :", ",", left)

	// Юзернеймы
	result.Names, left = findWithIndex(pageStr, "previous_usernames :[ ", " ],", left)

	// Рейтинговые и одобренные карты
	result.RankedBeatmapsetCount, left = findWithIndex(pageStr, "ranked_beatmapset_count :", ",", left)

	// Уровень
	result.Level, left = findWithIndex(pageStr, "level :{ current :", ",", left)

	// Глобальный рейтинг
	result.GlobalRank, left = findWithIndex(pageStr, "global_rank :", ",", left)

	// PP-хи
	result.PP, left = findWithIndex(pageStr, "pp :", ",", left)

	// Всего очков
	result.RankedScore, left = findWithIndex(pageStr, "ranked_score :", ",", left)

	// Точность попаданий
	result.Accuracy, left = findWithIndex(pageStr, "hit_accuracy :", ",", left)

	// Количество игр
	result.PlayCount, left = findWithIndex(pageStr, "play_count :", ",", left)

	// Время в игре в секундах
	result.PlayTimeSeconds, left = findWithIndex(pageStr, "play_time :", ",", left)

	// Время в игре в часах
	duration, _ := time.ParseDuration(result.PlayTimeSeconds + "s")
	result.PlayTime = duration.String()

	// Рейтинговые очки
	result.TotalScore, left = findWithIndex(pageStr, "total_score :", ",", left)

	// Всего попаданий
	result.TotalHits, left = findWithIndex(pageStr, "total_hits :", ",", left)

	// Максимальное комбо
	result.MaximumCombo, left = findWithIndex(pageStr, "maximum_combo :", ",", left)

	// Реплеев просмотрено другими
	result.Replays, left = findWithIndex(pageStr, "replays_watched_by_others :", ",", left)

	// SS-ки
	result.SS, left = findWithIndex(pageStr, "grade_counts :{ ss :", ",", left)

	// SSH-ки
	result.SSH, left = findWithIndex(pageStr, "ssh :", ",", left)

	// S-ки
	result.S, left = findWithIndex(pageStr, "s :", ",", left)

	// SH-ки
	result.SH, left = findWithIndex(pageStr, "sh :", ",", left)

	// A-хи
	result.A, left = findWithIndex(pageStr, "a :", "}", left)

	// Рейтинг в стране
	result.CountryRank, _ = findWithIndex(pageStr, "country_rank :", ",", left)

	// Уровень подписки
	result.SupportLvl = find(pageStr, "support_level :", ",")

	return result
}

// Функция получения информации о пользователе
func getOnlineInfo(id string) OnlineInfo {

	// Формирование и исполнение запроса
	resp, err := http.Get("https://osu.ppy.sh/users/" + id)
	if err != nil {
		return OnlineInfo{}
	}

	// Запись респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Проверка на страницу пользователя
	if !strings.Contains(string(body), "js-react--profile") {
		return OnlineInfo{}
	}

	// Структура, которую будет возвращать функция
	result := OnlineInfo{}

	// Статус в сети
	result.Status = find(string(body), "is_online&quot;:", ",")

	return result
}

// Функция отправки информации о пользователе
func sendUserInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(getUserInfo(mux.Vars(request)["id"], mux.Vars(request)["mode"]))
}

// Функция отправки информации о статусе пользователя
func sendOnlineInfo(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(getOnlineInfo(mux.Vars(request)["id"]))
}

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))

	/*	Сетап для тестов
		var sd int64
		for i := 0; i < 100; i++ {
			t := time.Now()
			getUserInfo("29829158", "")
			sd += time.Since(t).Milliseconds()
			fmt.Println("{", i, "}cur: \t", sd/(int64(i)+1))
		}
		println("fin:\t", sd/100)
	*/

	// Роутер
	router := mux.NewRouter()

	// Маршруты

	router.HandleFunc("/user", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/", sendUserInfo).Methods("GET")

	router.HandleFunc("/user/{id}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}/", sendUserInfo).Methods("GET")

	router.HandleFunc("/online/{id}", sendOnlineInfo).Methods("GET")
	router.HandleFunc("/online/{id}/", sendOnlineInfo).Methods("GET")

	// Запуск API

	// Для Heroku
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

	// Для локалхоста (127.0.0.1:8080/)
	log.Fatal(http.ListenAndServe(":8080", router))
}
