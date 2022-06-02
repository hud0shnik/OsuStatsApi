package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Структура для хранения полной информации о пользователе
type UserInfo struct {
	Username                 string  `json:"username"`
	Names                    string  `json:"previous_usernames"`
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
	DefaultGroup             string  `json:"default_group"`
	IsOnline                 string  `json:"is_online"`
	IsActive                 string  `json:"is_active"`
	IsDeleted                string  `json:"is_deleted"`
	IsBot                    string  `json:"is_bot"`
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
	FollowerCount            string  `json:"follower_count"`
	BestBeatMap              beatMap `json:"best_beat_map"`
}

// Структура для проверки статуса пользователя
type OnlineInfo struct {
	Status string `json:"is_online"`
}

// Структура для хранения информации о мапе
type beatMap struct {
	DifficultyRating string `json:"difficulty_rating"`
	Id               string `json:"id"`
	Rank             string `json:"rank"`
	EndedAt          string `json:"ended_at"`
	Accuracy         string `json:"accuracy"`
	MaximumCombo     string `json:"maximum_combo"`
	PP               string `json:"pp"`
	Passed           string `json:"passed"`
	TotalScore       string `json:"total_score"`
	LegacyPerfect    string `json:"legacy_perfect"`
	Replay           string `json:"replay"`
	Mode             string `json:"mode"`
	Status           string `json:"status"`
	TotalLength      string `json:"total_length"`
	Ar               string `json:"ar"`
	Bpm              string `json:"bpm"`
	Convert          string `json:"convert"`
	CountCircles     string `json:"count_circles"`
	CountSliders     string `json:"count_sliders"`
	CountSpinners    string `json:"count_spinners"`
	Cs               string `json:"cs"`
	DeletedAt        string `json:"deleted_at"`
	Drain            string `json:"drain"`
	HitLength        string `json:"hit_length"`
	IsScoreable      string `json:"is_scoreable"`
	LastUpdated      string `json:"last_updated"`
	ModeInt          string `json:"mode_int"`
	PassCount        string `json:"pass_count"`
	PlayCount        string `json:"play_count"`
	Ranked           string `json:"ranked"`
	Url              string `json:"url"`
	Checksum         string `json:"checksum"`
	Creator          string `json:"creator"`
	FavoriteCount    string `json:"favorite_count"`
	Hype             string `json:"hype"`
	Nsfw             string `json:"nsfw"`
	Offset           string `json:"offset"`
	Spotlight        string `json:"spotlight"`
}

// Функция поиска. Возвращает искомое значение и индекс
func find(str string, subStr string, char byte) (string, int) {

	// Поиск индекса начала нужной строки
	left := strings.Index(str, subStr) + len(subStr)

	// Проверка на существование нужной строки
	if left > len(subStr)-1 {

		// Крайняя часть искомой строки
		right := left

		for ; str[right] != char; right++ {
			// Доводит str[right] до символа char
		}
		return str[left:right], right
	}

	return "", 0
}

// Функция получения информации о пользователе
func getUserInfo(id, mode string) UserInfo {

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

	// Обрезка юзелесс части html'ки
	pageStr = pageStr[strings.Index(pageStr, "current_mode"):]
	pageStr = strings.ReplaceAll(pageStr, "&quot;", " ")

	// Сохранение html'ки в файл sample.html

	if err := os.WriteFile("sampleN.html", []byte(pageStr), 0666); err != nil {
		log.Fatal(err)
	}

	// Структура, которую будет возвращать функция
	result := UserInfo{
		UserID: id,
	}

	i, i2 := 0, 0

	/* -----------------------------------------------------------
	# Далее происходит заполнение полей функцией find			 #
	# после каждого поиска тело сайта обрезается для оптимизации #
	------------------------------------------------------------ */

	// Лучшая мапа
	result.BestBeatMap.Accuracy, _ = find(pageStr, "accuracy :", ',')
	result.BestBeatMap.Id, _ = find(pageStr, "beatmap_id :", ',')
	result.BestBeatMap.EndedAt, _ = find(pageStr, "ended_at : ", ' ')
	result.BestBeatMap.MaximumCombo, _ = find(pageStr, "max_combo :", ',')
	result.BestBeatMap.Passed, _ = find(pageStr, "passed :", ',')
	result.BestBeatMap.Rank, _ = find(pageStr, "rank : ", ' ')
	result.BestBeatMap.TotalScore, _ = find(pageStr, "total_score :", ',')
	result.BestBeatMap.LegacyPerfect, _ = find(pageStr, "legacy_perfect :", ',')
	result.BestBeatMap.PP, _ = find(pageStr, "pp :", ',')
	result.BestBeatMap.Replay, _ = find(pageStr, "replay :", ',')
	result.BestBeatMap.DifficultyRating, _ = find(pageStr, "difficulty_rating :", ',')
	result.BestBeatMap.Mode, _ = find(pageStr, "mode : ", ' ')
	result.BestBeatMap.Status, _ = find(pageStr, "status : ", ' ')
	result.BestBeatMap.TotalLength, _ = find(pageStr, "total_length :", ',')
	result.BestBeatMap.Ar, _ = find(pageStr, "ar :", ',')
	result.BestBeatMap.Bpm, _ = find(pageStr, "bpm :", ',')
	result.BestBeatMap.Convert, _ = find(pageStr, "convert :", ',')
	result.BestBeatMap.CountCircles, _ = find(pageStr, "count_circles :", ',')
	result.BestBeatMap.CountSliders, _ = find(pageStr, "count_sliders :", ',')
	result.BestBeatMap.CountSpinners, _ = find(pageStr, "count_spinners :", ',')
	result.BestBeatMap.Cs, _ = find(pageStr, " cs :", ',')
	result.BestBeatMap.DeletedAt, _ = find(pageStr, "deleted_at :", ',')
	result.BestBeatMap.Drain, _ = find(pageStr, "drain :", ',')
	result.BestBeatMap.HitLength, _ = find(pageStr, "hit_length :", ',')
	result.BestBeatMap.IsScoreable, _ = find(pageStr, "is_scoreable :", ',')
	result.BestBeatMap.LastUpdated, _ = find(pageStr, "last_updated : ", ' ')
	result.BestBeatMap.ModeInt, _ = find(pageStr, "mode_int :", ',')
	result.BestBeatMap.PassCount, _ = find(pageStr, "passcount :", ',')
	result.BestBeatMap.PlayCount, _ = find(pageStr, "playcount :", ',')
	result.BestBeatMap.Ranked, _ = find(pageStr, "ranked :", ',')
	result.BestBeatMap.Url, _ = find(pageStr, "url : ", ' ')
	result.BestBeatMap.Url = strings.ReplaceAll(result.BestBeatMap.Url, "\\", "")
	result.BestBeatMap.Checksum, i = find(pageStr, "checksum : ", ' ')
	pageStr = pageStr[i:]
	//
	result.BestBeatMap.Creator, _ = find(pageStr, "creator : ", ' ')
	result.BestBeatMap.FavoriteCount, _ = find(pageStr, "favourite_count :", ',')
	result.BestBeatMap.Hype, _ = find(pageStr, "hype :", ',')
	result.BestBeatMap.Nsfw, _ = find(pageStr, "nsfw :", ',')
	result.BestBeatMap.Offset, _ = find(pageStr, "offset :", ',')
	result.BestBeatMap.Spotlight, _ = find(pageStr, "spotlight :", ',')

	pageStr = pageStr[i:]

	// В последний раз был в сети
	result.LastVisit, i = find(pageStr, "last_visit :", ',')
	i2 += i

	// Сообщения только от друзей
	result.PmFriendsOnly, i = find(pageStr[i2:], "pm_friends_only :", ',')
	i2 += i

	// Ссылка на аватар
	result.AvatarUrl, i = find(pageStr[i2:], "avatar_url : ", ' ')
	result.AvatarUrl = strings.ReplaceAll(result.AvatarUrl, "\\", "")
	i2 += i

	// Код страны
	result.CountryCode, i = find(pageStr[i2:], "country_code : ", ' ')
	i2 += i

	// Группа
	result.DefaultGroup, i = find(pageStr[i2:], "default_group : ", ' ')
	i2 += i

	// Активность
	result.IsActive, i = find(pageStr[i2:], "is_active :", ',')
	i2 += i

	// Бот
	result.IsBot, i = find(pageStr[i2:], "is_bot :", ',')
	i2 += i

	// Удалённый профиль
	result.IsDeleted, i = find(pageStr[i2:], "is_deleted :", ',')
	i2 += i

	// Статус в сети
	result.IsOnline, i = find(pageStr[i2:], "is_online :", ',')
	i2 += i

	// Подписка
	result.IsSupporter, i = find(pageStr[i2:], "is_supporter :", ',')
	i2 += i

	// Цвет профиля
	result.ProfileColor, i = find(pageStr[i2:], "profile_colour :", ',')
	i2 += i

	// Юзернейм
	result.Username, i = find(pageStr[i2:], "username : ", ' ')
	i2 += i

	// Количество игр карт
	result.BeatmapPlaycountsCount, i = find(pageStr[i2:], "beatmap_playcounts_count :", ',')
	i2 += i

	// Количество комментариев
	result.CommentsCount, i = find(pageStr[i2:], "comments_count :", ',')
	i2 += i

	// Количество любимых карт
	result.FavoriteBeatmapsetCount, i = find(pageStr[i2:], "favourite_beatmapset_count :", ',')
	i2 += i

	// Подписчики
	result.FollowerCount, i = find(pageStr[i2:], "follower_count :", ',')
	i2 += i

	// Заброшенные карты
	result.GraveyardBeatmapsetCount, i = find(pageStr[i2:], "graveyard_beatmapset_count :", ',')
	i2 += i

	// Карты с гостевым участием
	result.GuestBeatmapsetCount, i = find(pageStr[i2:], "guest_beatmapset_count :", ',')
	i2 += i

	// Карты на рассмотрении
	result.PendingBeatmapsetCount, i = find(pageStr[i2:], "pending_beatmapset_count :", ',')
	i2 += i

	// Юзернеймы
	result.Names, i = find(pageStr[i2:], "previous_usernames :[ ", ']')
	i2 += i

	// Рейтинговые и одобренные карты
	result.RankedBeatmapsetCount, i = find(pageStr[i2:], "ranked_beatmapset_count :", ',')
	i2 += i

	// Уровень
	result.Level, i = find(pageStr[i2:], "level :{ current :", ',')
	i2 += i

	// Глобальный рейтинг
	result.GlobalRank, i = find(pageStr[i2:], "global_rank :", ',')
	i2 += i

	// PP-хи
	result.PP, i = find(pageStr[i2:], "pp :", ',')
	i2 += i

	// Всего очков
	result.RankedScore, i = find(pageStr[i2:], "ranked_score :", ',')
	i2 += i

	// Точность попаданий
	result.Accuracy, i = find(pageStr[i2:], "hit_accuracy :", ',')
	i2 += i

	// Количество игр
	result.PlayCount, i = find(pageStr[i2:], "play_count :", ',')
	i2 += i

	// Время в игре в секундах
	result.PlayTimeSeconds, i = find(pageStr[i2:], "play_time :", ',')
	i2 += i

	// Время в игре в часах
	duration, _ := time.ParseDuration(result.PlayTimeSeconds + "s")
	result.PlayTime = duration.String()

	// Рейтинговые очки
	result.TotalScore, i = find(pageStr[i2:], "total_score :", ',')
	i2 += i

	// Всего попаданий
	result.TotalHits, i = find(pageStr[i2:], "total_hits :", ',')
	i2 += i

	// Максимальное комбо
	result.MaximumCombo, i = find(pageStr[i2:], "maximum_combo :", ',')
	i2 += i

	// Реплеев просмотрено другими
	result.Replays, i = find(pageStr[i2:], "replays_watched_by_others :", ',')
	i2 += i

	// SS-ки
	result.SS, i = find(pageStr[i2:], "grade_counts :{ ss :", ',')
	i2 += i

	// SSH-ки
	result.SSH, i = find(pageStr[i2:], "ssh :", ',')
	i2 += i

	// S-ки
	result.S, i = find(pageStr[i2:], "s :", ',')
	i2 += i

	// SH-ки
	result.SH, i = find(pageStr[i2:], "sh :", ',')
	i2 += i

	// A-хи
	result.A, i = find(pageStr[i2:], "a :", '}')
	i2 += i

	// Рейтинг в стране
	result.CountryRank, i = find(pageStr[i2:], "country_rank :", ',')
	i2 += i

	// Уровень подписки
	result.SupportLvl, _ = find(pageStr[i2:], "support_level :", ',')

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
	result.Status, _ = find(string(body), "is_online&quot;:", ',')

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

	router.HandleFunc("/user/{id}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}", sendUserInfo).Methods("GET")
	router.HandleFunc("/user/{id}/{mode}/", sendUserInfo).Methods("GET")

	router.HandleFunc("/online/{id}", sendOnlineInfo).Methods("GET")
	router.HandleFunc("/online/{id}/", sendOnlineInfo).Methods("GET")

	// Запуск API

	// Для Heroku
	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

	// Для локалхоста (127.0.0.1:8080/)
	log.Fatal(http.ListenAndServe(":8080", router))
}
