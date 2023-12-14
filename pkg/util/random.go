package util

import (
	"math/rand"
	"time"
)

// 字母 a-zA-Z
const stringUpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const stringLowerLetters = "abcdefghijklmnopqrstuvwxyz"

// 数字
const numberLetters = "0123456789"

// 特殊字符
const specialLetters = "!@#$%^&*"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

/**
 * 基础随机方法
 */
func baseRandStr(n int, value string) string {
	b := make([]byte, n)
	len := len(value)
	for i := range b {
		b[i] = value[rand.Intn(len)]
	}
	return string(b)
}

/**
 * 随机字符串 大写
 */
func StringUpperRandom(len int) string {
	return baseRandStr(len, stringUpperLetters)
}

/**
 * 随机字符串 小写
 */
func StringLowerRandom(len int) string {
	return baseRandStr(len, stringLowerLetters)
}

/**
 * 随机字符串
 */
func StringRandom(len int) string {
	return baseRandStr(len, stringLowerLetters+stringUpperLetters)
}

/**
 * 随机数字
 */
func NumberRandom(len int) string {
	return baseRandStr(len, numberLetters)
}

func Random(len int) string {
	return baseRandStr(len, stringLowerLetters+stringUpperLetters+numberLetters+specialLetters)
}
