package util

import (
	"math/rand"
	"time"
)

// 字母 a-zA-Z
const stringLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 数字
const numberLetters = "0123456789"

// 特殊字符
const specialLetters = "!@#$%^&*"

/**
 * 基础随机方法
 */
func baseRandStr(n int, value string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	len := len(value)
	for i := range b {
		b[i] = value[rand.Intn(len)]
	}
	return string(b)
}

/**
 * 随机字符串
 */
func StringRandom(len int) string {
	return baseRandStr(len, stringLetters)
}

/**
 * 随机数字
 */
func NumberRandom(len int) string {
	return baseRandStr(len, numberLetters)
}

func Random(len int) string {
	return baseRandStr(len, stringLetters+numberLetters+specialLetters)
}
