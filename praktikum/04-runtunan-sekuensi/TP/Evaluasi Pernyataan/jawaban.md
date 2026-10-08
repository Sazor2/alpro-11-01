## 1
Nilai Akhir Variabel result: 25

## 2
![Output program pernyataan](TP/pernyataan/output.png)

## 3
### kondisi 1
Ya, benar (10 > 5 adalah true).Yang terjadi selanjutnya: Program masuk ke dalam blok if lalu mengevaluasi kondisi bersarang y < 10 (5 < 10 [true]). Karena benar, program mengeksekusi result = x + y ($10 + 5$), sehingga nilai result berubah dari 0 menjadi 15

### kondisi 2
Ya, benar (15 > 10 [true] DAN 10 == 10 [true] $\rightarrow$ true).Pengaruhnya ke result: Program mengeksekusi result += z ($15 + 15$), sehingga nilai result bertambah menjadi 30

### kondisi 3
Ya, benar (10 == 10 [true] ATAU 5 > 10 [false] $\rightarrow$ true).Yang terjadi: Program mengeksekusi baris result += 5 ($30 + 5$), sehingga nilai result bertambah menjadi 35. (Bagian else if dan else dilewati).