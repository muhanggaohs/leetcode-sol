Penjelasan:
<br><br>
1. Pertama, kita menggunakan dua stack (tumpukan) untuk menyimpan digit-digit dari kedua linked list. Stack digunakan untuk membalikkan urutan digit dari linked list.
2. Kita melakukan iterasi pada linked list pertama (l1) dan memasukkan setiap node ke dalam stack1.
3. Selanjutnya, kita melakukan iterasi pada linked list kedua (l2) dan memasukkan setiap node ke dalam stack2.
4. Setelah itu, kita inisialisasi variabel carry (pembawa) dengan 0, yang akan kita gunakan saat melakukan penjumlahan.
5. Selama stack1 atau stack2 masih memiliki elemen, atau carry bernilai lebih dari 0, kita melakukan penjumlahan.
6. Pada setiap iterasi, kita mengambil digit pertama dari stack1 dan stack2 (jika masih ada), atau menggunakan 0 jika stack sudah habis.
7. Kemudian, kita menjumlahkan digit dari stack1, stack2, dan carry.
8. Jika hasil penjumlahan lebih dari 9, kita perbarui carry dengan 1 dan ambil digit terakhir dari hasil penjumlahan.
9. Selanjutnya, kita membuat node baru dengan digit terakhir hasil penjumlahan dan menghubungkannya dengan node sebelumnya.
10. Terakhir, kita mengembalikan linked list hasil penjumlahan dengan mengembalikan node terakhir yang telah tersimpan dalam variabel prev.
<br><br>
Dengan menggunakan pendekatan ini, kita dapat menambahkan dua angka yang direpresentasikan sebagai linked list tanpa harus membalik urutan linked list asli.