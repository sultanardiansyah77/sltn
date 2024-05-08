var timeInSeconds = 2 * 60 * 60;

// Fungsi untuk memperbarui countdown
function updateCountdown() {
    var hours = Math.floor(timeInSeconds / 3600);
    var minutes = Math.floor((timeInSeconds % 3600) / 60);
    var seconds = timeInSeconds % 60;

    // Tampilkan waktu countdown dalam format jam:menit:detik
    document.getElementById('countdown').textContent = hours + ' jam ' + minutes + ' menit ' + seconds + ' detik';

    // Kurangi waktu
    timeInSeconds--;

    // Jeda selama 1 detik sebelum memperbarui kembali
    setTimeout(updateCountdown, 1000);

    // Cek apakah waktu telah habis
    if (timeInSeconds < 0) {
        document.getElementById('countdown').textContent = 'Waktu telah habis!';
        // Tambahkan tindakan setelah waktu habis jika diperlukan
    }
}

// Panggil fungsi updateCountdown untuk memulai countdown
updateCountdown();