Pada soal ini, diberikan sebuah array bilangan bulat `**arr**` dan sebuah selisih bilangan `**difference**`. Kita diminta untuk mengembalikan panjang subsequence terpanjang dalam `**arr**` yang membentuk deret aritmatika dengan selisih `**difference**`.
<br><br>
Sebuah subsequence adalah sekumpulan bilangan dalam `**arr**` yang dapat diperoleh dengan menghapus beberapa atau tidak ada elemen tanpa mengubah urutan elemen yang tersisa.
<br><br>
**Contoh 1**:
**Input**: arr = [1,2,3,4], difference = 1
**Output**: 4
**Penjelasan**: Subsequence aritmatika terpanjang adalah [1,2,3,4].
<br><br>
**Contoh 2**:
**Input**: arr = [1,3,5,7], difference = 1
**Output**: 1
**Penjelasan**: Subsequence aritmatika terpanjang adalah satu elemen saja.
<br><br>
**Batasan**:
* 1 <= arr.length <= 105
* -104 <= arr[i], difference <= 104
<br><br>
Pada solusi di atas, menggunakan pendekatan dynamic programming dengan memanfaatkan map **`dp`** untuk menyimpan panjang subsequence terpanjang untuk setiap bilangan dalam `**arr**`. Iterasi dilakukan pada setiap elemen `**num**` dalam `**arr**`. Jika ada elemen sebelumnya dengan selisih difference dari `**num**`, maka panjang subsequence num dapat ditambahkan dengan 1 berdasarkan subsequence yang sudah ada untuk elemen sebelumnya. Jika tidak ada, maka panjang subsequence `**num**` diatur menjadi 1. Selama iterasi, kami juga memperbarui `**maxLength**` untuk melacak panjang subsequence terpanjang yang ditemukan sejauh ini.
<br><br>
Pada akhirnya, fungsi mengembalikan `**maxLength**` sebagai hasil.