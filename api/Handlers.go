package main

import (
	"encoding/json"
	"fmt"
    "log"
	"net/http"
	"net/smtp"
	"os"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/lpernett/godotenv"
)

type createPoolResponse struct{
    Uuid string `json:"uuid"`
}

type addMemberRequest struct{
    Uid string `json:"uid"`
    Person Person `json:"person"`
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}


type createPoolHandler struct {}

func (h * createPoolHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
    Uuid := uuid.New().String()
    responseObj := &createPoolResponse{
        Uuid: Uuid,
    }

    res, err := json.Marshal(responseObj)
    if err != nil {
        w.WriteHeader(500)
        w.Write([]byte("Something went wrong"))
        return
    }

    if err := CreatePool(Uuid); err != nil {
        w.WriteHeader(500)
        w.Write([]byte("Something went wrong"))
        return
    }
    w.WriteHeader(200)
    w.Write(res)
}

func GetParams(r *http.Request) map[string]string {
    params := strings.Split(r.URL.String(), "?")[1]
    var paramList []string
    vars := map[string]string{}
    if strings.Contains(params, "&"){
        paramList = strings.Split(params, "&")
    } else {
        paramList = []string{params}
    }
    for _, item := range paramList {
        param := strings.Split(item, "=")
        vars[param[0]] = param[1]
    }
    return vars
}

type getPoolHandler struct {}
func (h * getPoolHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }

    // "eb2f08a5-12de-4253-9e14-cad1bf8f762e"
    if len(strings.Split(r.URL.String(), "?")) == 1 {
        w.WriteHeader(400)
        w.Write([]byte("Bad Request"))
        return
    }
    vars := GetParams(r)
    pool, err := GetPool(vars["uid"])
    if err != nil{
        if err.Error() == "No such pool" {
            w.WriteHeader(404)
            w.Write([]byte("Pool does not exist"))
        } else if err.Error() == "Error" {
            w.WriteHeader(500)
            w.Write([]byte("Something went wrong"))
        }
        return 
    }
    
    w.WriteHeader(200)
    w.Write(pool)
}

type addPersonHandler struct {}
func (h *addPersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
    var req addMemberRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("Bad Request"))
        return
    }
    if r.Header.Get("Content-Type") != "application/json" {
        w.WriteHeader(400)
        w.Write([]byte("Bad Request"))
        return
    }
    newPerson, err := InsertMember(&req)
    if err != nil {
        w.WriteHeader(404)
        w.Write([]byte("Pool does not exist"))
        return
    }
    if err := SendEmailJoinedPool(newPerson); err != nil {
	fmt.Println("Could not send email")
        w.WriteHeader(500)
        w.Write([]byte("Internal server error"))
        return
    }
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}

func SendEmailJoinedPool(p *Person) error {
    err := godotenv.Load()
    if err != nil {
        return err
    }
    pw := os.Getenv("GMAIL_PASSWORD")
    auth := smtp.PlainAuth("", "normananton03@gmail.com", pw, "smtp.gmail.com")
    to := []string{p.Email}

    // Create the message
//   msg := []byte(fmt.Sprintf("To: %s\r\n", p.Email) +
//       fmt.Sprintf("Subject: Hello %s %s\r\n", p.FirstName, p.LastName) +
//       "\r\n" +
//       "You have been added to a secret santa lotto pool!\r\n\r\nThe host of this pool will perform the lottery when everybody is in the pool. "+
//       "You will then be sent another email, telling you who was pulled for you to buy for.\r\n\r\nHappy Holidays!")
    msg := []byte(fmt.Sprintf("To: %s\r\n", p.Email) + 
       fmt.Sprintf("Subject: Hej %s %s\r\n", p.FirstName, p.LastName) +
       "\r\n" +
       "Du har blivit tillagt i en secret santa pool!\r\n\r\nVärden kommer utföra lottningen när alla är inne i poolen. "+
       "Du kommer då få ytterligare ett mail som meddelar vem du har blivit tilldelad att köpa en present till.\r\n\r\nGod jul!")

    err = smtp.SendMail("smtp.gmail.com:587", auth, "api", to, msg)
    if err != nil {
        log.Fatal(err)
        return err
    }
    return nil
}


type performLotteryHandler struct {}
func (h *performLotteryHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
    if len(strings.Split(r.URL.String(), "?")) == 1 {
        w.WriteHeader(400)
        w.Write([]byte("Bad Request"))
        return
    }
    params := GetParams(r)
    poolBytes, err := GetPool(params["uid"])
    if err != nil{
        w.WriteHeader(500)
        w.Write([]byte("Internal server error"))
        return
    }
    var pool Pool
    json.Unmarshal(poolBytes, &pool)

    result := GenerateResult(pool.Members)
    err = SendEmailLottery(result)
    if err != nil {
        w.WriteHeader(500)
        w.Write([]byte("Internal server error"))
        return
    }
    err = DeletePool(params["uid"])
    if err != nil{
        w.WriteHeader(500)
        w.Write([]byte("Internal server error"))
        return
    }

    w.WriteHeader(200)
    w.Write([]byte("OK"))
}

func GenerateResult(list []Person) map[Person]Person {
    pickFrom := []Person{}
    for _, item := range list{
        pickFrom = append(pickFrom, item)
    }
    pickFrom = append(pickFrom[len(pickFrom) / 2:], pickFrom[:len(pickFrom) / 2]...)
    slices.Reverse(pickFrom)
    res := map[Person]Person{}
    for _, person := range list{
        for {
            if person.toString() == pickFrom[0].toString(){
                // fmt.Printf("%s - %s\n", person.toString(), pickFrom[0].toString())
                slices.Reverse(pickFrom)
                continue
            }
            res[person] = pickFrom[0]
            if len(pickFrom) > 1{
                pickFrom = pickFrom[1:]
            }
            break
        }
    }
    // for key, value := range res{
    //     fmt.Printf("%s   -   %s\n", key.toString(), value.toString())
    // }

    return res
}


func SendEmailLottery(lottery map[Person]Person) error {
    err := godotenv.Load()
    if err != nil {
        return err
    }
    pw := os.Getenv("GMAIL_PASSWORD")
    auth := smtp.PlainAuth("", "normananton03@gmail.com", pw, "smtp.gmail.com")
    
    for person, buyTo := range lottery{
        to := []string{person.Email}

        msg := []byte(fmt.Sprintf("To: %s\r\n", person.Email) +
//       fmt.Sprintf("Subject: The Lottery is done %s %s!!\r\n", person.FirstName, person.LastName) +
//       "\r\n" +
//       fmt.Sprintf("Hello %s!\r\n\r\n", person.FirstName) +
//       "The lottery has now been performed! You have drawn... " +
//       fmt.Sprintf("%s %s!\r\n\r\n", buyTo.FirstName, buyTo.LastName) + "Good luck, and happy holidays!")
        fmt.Sprintf("Subject: Lottningen är gjord %s %s!!\r\n", person.FirstName, person.LastName) +
        "\r\n" +
        fmt.Sprintf("Hej %s!\r\n\r\n", person.FirstName) +
        "Lottningen är nu utförd! Du har dragit... " +
        fmt.Sprintf("%s %s att köpa till!\r\n\r\n", buyTo.FirstName, buyTo.LastName) + "Lycka till, och god jul!")

        err = smtp.SendMail("smtp.gmail.com:587", auth, "api", to, msg)
        if err != nil {
            log.Fatal(err)
            return err
        }
    }
    return nil
}
