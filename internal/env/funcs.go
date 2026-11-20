package env

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"os"
	"path"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "NO-HOSTNAME"
	}
	return hostname
}

func GenStrBytes(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func GenStrLength(length int) string {
	pickFrom := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	maxNum := big.NewInt(int64(len(pickFrom)))
	rStr := make([]rune, length)
	for i := range rStr {
		rNum, _ := rand.Int(rand.Reader, maxNum)
		rStr[i] = pickFrom[rNum.Int64()]
	}
	return string(rStr)
}

func GetDataPath() string {
	sDataPath := os.Getenv("OPENSTUB_DATA_PATH")
	if sDataPath == "" {
		sDataPath = "./data"
	}
	return path.Clean(sDataPath)
}

func GetSysConfigPath() string {
	dataPath := GetDataPath()
	return path.Join(dataPath, "config.json")
}

func GetSysDBPath() string {
	dataPath := GetDataPath()
	return path.Join(dataPath, "system.db")
}
