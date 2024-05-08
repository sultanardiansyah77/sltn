let selectedButton = null;

function updateSelectedDiamondPrice(price) {
    // Fungsi untuk memformat angka dengan titik sebagai pemisah ribuan
    function formatNumber(number) {
        return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
    }

    // Mengambil semua elemen dengan class 'price'
    var priceElements = document.querySelectorAll('.price');

    // Memperbarui setiap elemen dengan harga yang dipilih
    priceElements.forEach(function (priceElement) {
        var priceValue = priceElement.querySelector('span');
        var formattedPrice = formatNumber(price);
        priceValue.innerText = formattedPrice;

        // Menampilkan elemen harga jika diamond dipilih
        priceElement.style.display = 'block';
    });
}

function selectDiamond(button) {
    // Reset tombol sebelumnya
    if (selectedButton) {
        selectedButton.classList.remove('selected');
    }

    // Pilih tombol yang diklik
    button.classList.add('selected');
    selectedButton = button;

    // Ambil harga dari data-price
    var selectedDiamondPrice = button.dataset.price;

    // Memperbarui harga di semua elemen dengan class 'price'
    updateSelectedDiamondPrice(selectedDiamondPrice);

    // Lakukan sesuatu dengan nilai diamond yang dipilih (gantilah dengan aksi sesuai kebutuhan)
    console.log(`Selected Diamond Value: ${button.dataset.value}`);
}


function toggleExpand(boxId) {
    var expandableBoxes = document.querySelectorAll('.expandable-box');
    var expandIcon = document.getElementById(boxId).querySelector('.expand-icon');

    // Menutup semua kotak yang dapat diperluas
    expandableBoxes.forEach(function (box) {
        if (box.id !== boxId && box.classList.contains('expanded')) {
            box.classList.remove('expanded');
            box.querySelector('.expand-icon').innerHTML = '▼';
        }
    });

    // Toggle 'expanded' class untuk kotak yang dipilih
    var expandableBox = document.getElementById(boxId);
    expandableBox.classList.toggle('expanded');

    // Toggle ikon antara ▼ dan ▲
    expandIcon.innerHTML = (expandIcon.innerHTML === '▼') ? '▲' : '▼';
}

function handlePaymentClick(paymentName) {
    // Handle klik pada opsi pembayaran
    var paymentOption = event.currentTarget;

    // Ambil harga dari data-price
    var price = paymentOption.dataset.price;

    // Pilih tombol diamond yang sedang dipilih
    var selectedDiamond = document.querySelector('.diamond-option.selected');

    // Jika tidak ada diamond yang dipilih, keluar dari fungsi
    if (!selectedDiamond) {
        console.log('Pilih diamond terlebih dahulu');
        return;
    }

    // Hapus kelas 'selected' dari semua opsi pembayaran
    var allPaymentOptions = document.querySelectorAll('.payment-option');
    allPaymentOptions.forEach(function (option) {
        option.classList.remove('selected');
    });

    // Tambahkan kelas 'selected' ke opsi pembayaran yang diklik
    paymentOption.classList.add('selected');

    if (priceDisplay && priceValue) {
        // Setel harga ke dalam elemen span
        priceValue.innerText = price;

        // Tampilkan elemen harga jika diamond dipilih
        priceDisplay.style.display = 'block';
    }
}