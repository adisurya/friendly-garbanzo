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

Untuk dokumentasi lebih detail dapat diakses di http://localhost:11300/swagger/index.html setelah aplikasi dibuild dan di jalankan

Untuk config di simpan di file .env contohnya bisa di lihat di .env.example
