# friendly-garbanzo

Aplikasi order tiket event

Aplikasi ini dibuat untuk menggambarkan proses order tiket dari checkout/booking hingga ke payment.

terdapan beberapa API endpoint antara lain

GET /events
untuk menampilkan daftar event

GET /events/:id
untuk menampilkan detail event

POST /events
untuk create event

GET /tickets/:id
untuk menampilkan detail ticket

POST /tickets/booking
untuk checkout atau booking ticket

GET /tickets/inquiry
untuk cek tagihan pembayaran dari checkout / booking ticket

POST /tickets/payment
untuk pembayaran ticket


Untuk config di simpan di file .env contohnya bisa di lihat di .env.example

Untuk struktur table ada di folder sql

Agar aplikasi dapat dijalankan/dicompile perlu dilakukan generate dokumentasi terlebih dahulu, langkah-langkahnya adalah sebagai berikut:
1. atur configurasi pada .env file, jika belum ada dapat copy dari .env.example
3. sesuaikan host/port pada docs/swagger.json dengan konfigurasi di .env
4. jalankan aplikasi dengan perintah go run main.go atau dibuild terlebih dahulu dengan go build
5. dokumentasi dapat di akses di http://[host]:[port]/swagger/index.html (host dan port sesuai dengan konfigurasi). contoh: http://localhost:11300/swagger/index.html
