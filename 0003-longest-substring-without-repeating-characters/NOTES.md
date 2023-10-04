```
Penjelasan:
```
​
* Fungsi ini menggunakan 2 index `left` dan `right` untuk membentuk sliding window yang akan melacak substring tanpa karakter yang berulang. Nilai awal kedua index ini = 0.
* Map `lastSeen` digunakan untuk menyimpan index terakhir dari setiap karakter dalam string `s`. Map ini digunakan untuk melacak karakter terakhir kali ditemukan dalam substring.
* Selama iterasi melalui string `s` dengan index `right`:
1. Jika karakter saat ini (`s[right]`) sudah pernah ditemukan (`found` adalah `true`) dan indeks karakter tersebut lebih besar atau sama dengan `left`, maka kita memperbarui `left` menjadi index setelah karakter terakhir kali munculnya. Hal ini menghindari karakter yang berulang di dalam jendela geser.
2. Memeriksa apakah panjang substring saat ini (`right - left + 1`) lebih panjang dari panjang maksimum yang telah ditemukan sebelumnya (`maxLength`). Jika ya, kita perbarui `maxLength` dengan panjang substring saat ini.
3. Kami kemudian menyimpan indeks karakter saat ini (`right`) dalam map `lastSeen` untuk mengingat di mana karakter ini terakhir kali muncul dalam substring saat ini.
* Setelah iterasi selesai, `maxLength` akan berisi panjang dari substring terpanjang tanpa karakter yang berulang dalam string s, dan kita mengembalikan nilai ini sebagai hasil.
<br><br>
Dengan menggunakan pendekatan sliding window ini, kita dapat secara efisien mencari panjang substring terpanjang tanpa karakter yang berulang dalam string input `s` tanpa perlu memeriksa ulang karakter yang sudah diperiksa sebelumnya.