package pasiencontroller

import (
	"html/template"
	"net/http"

	"github.com/aalbaling/go-crud/entities"
	"github.com/aalbaling/go-crud/models"
)

var PasienModel = models.NewPasienModel()

func Index(w http.ResponseWriter, r *http.Request) {

	pasien, _ := PasienModel.FindAll()

	data := map[string]interface{}{
		"pasien": pasien,
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

		PasienModel.Create(pasien)
		data := map[string]interface{}{
			"pesan": "Data pasien berhasil disimpan",
		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(w, data)
	}

}
func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
