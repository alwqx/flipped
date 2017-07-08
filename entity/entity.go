package entity

// HeartBeat model for heartbeat
type HeartBeat struct {
	ID        int64 `json:"id"`
	Timestamp int64 `json:"timestamp"`
	HeartRate int16 `json:"heart_rate"`
}

type HeartStrength struct {
	ID            int64 `json:"id"`
	Timestamp     int64 `json:"timestamp"`
	HeartStrength int16 `json:"heart_strength"`
}

type KeyHeartStrength struct {
	Key   string        `json:"key"`
	Value HeartStrength `json:"value"`
}

type HSM map[string]HeartStrength

// config
type Mysql struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type Flipped struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type FlippedConfig struct {
	Mysql   Mysql   `json:"mysql"`
	Flipped Flipped `json:"flipped"`
}
