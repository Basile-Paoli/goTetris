package main

const (
	defaultTickDuration    = 10
	defaultTicksBeforeLock = 10
	defaultDAS             = 8
	defaultARR             = 3
	defaultSoftDropRate    = 0
)

type Config struct {
	tickDuration    int
	ticksBeforeLock int
	das             int
	arr             int
	softDropRate    int
}

func defaultConfig() Config {
	return Config{
		tickDuration:    defaultTickDuration,
		ticksBeforeLock: defaultTicksBeforeLock,
		das:             defaultDAS,
		arr:             defaultARR,
		softDropRate:    defaultSoftDropRate,
	}
}
