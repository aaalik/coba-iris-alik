package controller

import (
	"../model"
	"encoding/json"
	"github.com/kataras/iris/context"
	"strconv"
	"time"
)

func Tambah(ctx context.Context){
	userID := 1+1
    ctx.Writef("Hasil penjumlahan: %d", userID)
}

func Kurang(ctx context.Context){
	userID := 1-1
    ctx.Writef("Hasil pengurangan: %d", userID)
}

func Kali(ctx context.Context){
	userID := 1*1
    ctx.Writef("Hasil perkalian: %d", userID)
}

func Bagi(ctx context.Context){
	userID := 1/1
    ctx.Writef("Hasil pembagian: %d", userID)
}

func Create(ctx context.Context){
	nama := ctx.PostValue("nama")
	lahir := ctx.PostValue("lahir")
	t, _ := time.Parse("2006-01-02", lahir)
	model.Buat(nama, t)
	ctx.Writef(t.String())
}

func Read(ctx context.Context){
	x := model.Baca()
	i, _ := json.Marshal(x)
	ctx.Writef(string(i))
}

func Update(ctx context.Context){
	id := ctx.PostValue("id")
	i, _ := strconv.Atoi(id)
	nama := ctx.PostValue("nama")
	lahir := ctx.PostValue("lahir")
	t, _ := time.Parse("2006-01-02", lahir)
	model.Ubah(nama, t, i)
	ctx.Writef(t.String())
}

func Delete(ctx context.Context){
	id := ctx.PostValue("id")
	i, _ := strconv.Atoi(id)
	model.Hapus(i)
}