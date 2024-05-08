document.addEventListener('DOMContentLoaded', function () {
    // active nav link
    var ul = document.querySelector('nav ul');
    var li = document.querySelectorAll('nav ul li');

    li.forEach(el => {
        el.addEventListener('click', function () {
            ul.querySelector('.active').classList.remove('active');
            el.classList.add('active');
        });
    });
    // image slider
    const swiper = new Swiper('.swiper-container', {
        autoplay: {
            delay: 3000,
            disableOnInteraction: false,
        },
        loop: true,
        pagination: {
            el: '.swiper-pagination',
            clickable: true,
        },
        preloadImages: false,
        lazy: true,
    });
});

function toggleContent(contentId, button) {
    // Menyembunyikan semua konten
    $(".content").hide();

    // Menampilkan atau menyembunyikan konten yang dipilih
    $("#" + contentId).show();

    // Menghapus kelas "active" dari semua tombol dengan kelas "dert"
    $(".btn.dert").removeClass("active");

    // Menambahkan kelas "active" pada tombol yang diklik
    $(button).addClass("active");
}


