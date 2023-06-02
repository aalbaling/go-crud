package pasiencontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/aalbaling/go-crud/entities"
	"github.com/aalbaling/go-crud/libraries"
	"github.com/aalbaling/go-crud/models"
)

var validation = libraries.NewValidation()
var PasienModel = models.NewPasienModel()

func Index(w http.ResponseWriter, r *http.Request) {

	Pasien, _ := PasienModel.FindAll()

	data := map[string]interface{}{
		"pasien": Pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var pasien entities.Pasien
		pasien.NamaLengkap = r.Form.Get("nama_lengkap")
		pasien.NIK = r.Form.Get("nik")
		pasien.JenisKelamin = r.Form.Get("jenis_kelamin")
		pasien.TempatLahir = r.Form.Get("tempat_lahir")
		pasien.TanggalLahir = r.Form.Get("tanggal_lahir")
		pasien.Alamat = r.Form.Get("alamat")
		pasien.NoHp = r.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)
		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors

		} else {
			data["pesan"] = "Data pasien berhasil disimpan"
			PasienModel.Create(pasien)

		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(w, data)
	}

}
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		queryString := r.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var pasien entities.Pasien
		PasienModel.Find(id, &pasien)

		data := map[string]interface{}{
			"pasien": pasien,
		}

		temp, err := template.ParseFiles("views/pasien/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var pasien entities.Pasien
		pasien.Id, _ = strconv.ParseInt(r.Form.Get("id"), 10, 64)
		pasien.NamaLengkap = r.Form.Get("nama_lengkap")
		pasien.NIK = r.Form.Get("nik")
		pasien.JenisKelamin = r.Form.Get("jenis_kelamin")
		pasien.TempatLahir = r.Form.Get("tempat_lahir")
		pasien.TanggalLahir = r.Form.Get("tanggal_lahir")
		pasien.Alamat = r.Form.Get("alamat")
		pasien.NoHp = r.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)
		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors

		} else {
			data["pesan"] = "Data pasien berhasil diperbaharui"
			PasienModel.Update(pasien)

		}

		temp, _ := template.ParseFiles("views/pasien/edit.html")
		temp.Execute(w, data)
	}

}
func Delete(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	PasienModel.Delete(id)

	http.Redirect(w, r, "/pasien", http.StatusSeeOther)

}
