Fungsi **`twoSum`** akan menerima input berupa sebuah array **`nums`** yang berisi bilangan bulat dan sebuah bilangan bulat **`target`**. Tujuan dari fungsi ini adalah untuk mencari dua angka dalam array **`nums`** yang jika dijumlahkan akan menghasilkan nilai **`target`**. Fungsi ini akan mengembalikan indeks dari dua angka tersebut dalam bentuk array.
​<br><br>
Misalkan terdapat array **`nums`** sebagai berikut:
​<br><br>
nums = [2, 7, 11, 15]
​<br><br>
Dan target = 9
​<br><br>
Kita harus mencari dua angka dalam array **`nums`** yang jika dijumlahkan akan menghasilkan 9. Pada contoh ini, angka 2 (pada indeks 0) dan angka 7 (pada indeks 1) jika dijumlahkan akan menghasilkan 9.
​<br><br>
Sehingga output yang diharapkan adalah [0, 1], yang merupakan indeks dari dua angka tersebut.
​<br>
​<br>
**`Penjelasan`**:
​<br>
<br>
* Pertama, kita membuat sebuah map **`numMap`** yang akan digunakan untuk menyimpan angka dan indeksnya.
* Selanjutnya, kita melakukan iterasi terhadap array **`nums`** menggunakan **`range`**. Pada setiap iterasi, kita akan memeriksa apakah ada angka yang merupakan komplement dari angka saat ini terhadap **`target`**.
* Komplement dihitung dengan mengurangi angka saat ini dari **`target`**.
* Jika komplement tersebut sudah ada dalam map **`numMap`**, artinya kita telah menemukan pasangan angka yang dijumlahkan akan menghasilkan **`target`**. Kita mengembalikan array berisi indeks angka komplement dan indeks angka saat ini.
* Jika komplement tidak ditemukan dalam map **`numMap`**, kita menyimpan angka saat ini ke dalam map **`numMap`** bersama dengan indeksnya.
* Jika setelah iterasi selesai, tidak ditemukan pasangan angka yang memenuhi kondisi, maka kita mengembalikan **`nil`** karena tidak ada solusi yang valid.
​<br>
<br>
Fungsi ini memiliki kompleksitas waktu O(n), di mana n adalah jumlah elemen dalam array **`nums`**. Hal ini karena kita melakukan iterasi satu kali untuk mencari pasangan angka yang memenuhi kondisi.
