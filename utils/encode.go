package utils

import (
    "crypto/md5"
    "fmt"
    "io"
    "math"
    "math/rand"
    "time"
)

func GenerateUrl(num int64) string{
    rand.Seed(int64(time.Now().UnixNano()) + num)
    rnd := rand.Intn(int(math.Pow(10, 16)))
    s := base36Encode(rnd)
    return s
}

func GeneratePassword() string{
    rand.Seed(time.Now().Unix())
    rnd := rand.Intn(int(math.Pow(10, 16)))
    s := base36Encode(rnd)
    return s
}

func GenFilename(foldername, filename string, username string) string{
    h := md5.New()
    io.WriteString(h, foldername + "_" + filename + "_" + username)
    return string(fmt.Sprintf("%x", h.Sum(nil)))
}

func base36Encode(num int) string {
    res := ""
    if num <= 0{
        return res
    }else{
        for num != 0{
            index := num%36
            num = int(num/36)
            s := "0123456789abcdefghijklmnopqrstuvwxyz"
            res += string(s[index])
        }
        return res
    }
}