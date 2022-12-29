-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 23 Des 2022 pada 10.52
-- Versi server: 10.1.38-MariaDB
-- Versi PHP: 7.3.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_gamestore`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `coupon`
--

CREATE TABLE `coupon` (
  `coupon_id` varchar(255) NOT NULL,
  `customer_id` varchar(255) NOT NULL,
  `types` varchar(255) NOT NULL,
  `status` tinyint(2) NOT NULL DEFAULT '0',
  `expired_date` date NOT NULL,
  `created_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `criteria`
--

CREATE TABLE `criteria` (
  `criteria_id` int(11) NOT NULL,
  `criteria_name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `criteria`
--

INSERT INTO `criteria` (`criteria_id`, `criteria_name`) VALUES
(1, 'Service Console'),
(2, 'New Console'),
(3, 'New Game'),
(4, 'Second Game'),
(5, 'Accessories Console');

-- --------------------------------------------------------

--
-- Struktur dari tabel `customer`
--

CREATE TABLE `customer` (
  `customer_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `created_time` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `customer`
--

INSERT INTO `customer` (`customer_id`, `name`, `alamat`, `phone_number`, `created_time`) VALUES
('CUST12226895', 'Rida', 'Bonang', '081818181818', '2022-12-11'),
('CUST15399180', 'Yoke', 'Tangerang', '081111111111', '2022-12-12'),
('CUST16081110', 'Cahyati', 'Bekasi', '082365957600', '2022-12-23'),
('CUST22456861', 'Rahmat', 'Depok', '081282329488', '2022-12-16'),
('CUST25113358', 'Cahyati', 'Bekasi', '082365957600', '2022-12-23'),
('CUST29055003', 'Budi', 'Bekasi', '085600912383', '2022-12-16'),
('CUST29347724', 'Shinta', 'Jakarta', '083265732190', '2022-12-12');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaction`
--

CREATE TABLE `transaction` (
  `transaction_id` varchar(255) NOT NULL,
  `customer_id` varchar(255) DEFAULT NULL,
  `revenue` int(11) DEFAULT '0',
  `coupon_id` varchar(255) DEFAULT NULL,
  `discount_price` int(11) DEFAULT '0',
  `purchase_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `transaction`
--

INSERT INTO `transaction` (`transaction_id`, `customer_id`, `revenue`, `coupon_id`, `discount_price`, `purchase_date`) VALUES
('TRAX40818546', 'CUST15399180', 5000000, '', 0, '2022-11-05'),
('TRAX49773756', 'CUST12226895', 1500000, '', 0, '2022-12-13'),
('TRAX62331089', 'CUST15399180', 1850000, '', 0, '2022-12-08'),
('TRAX67201794', 'CUST15399180', 5600000, '', 0, '2022-11-05'),
('TRAX74160559', 'CUST15399180', 2250000, '', 0, '2022-12-05'),
('TRAX78567088', 'CUST12226895', 1500000, '', 0, '2022-12-13');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaction_items`
--

CREATE TABLE `transaction_items` (
  `item_id` int(11) NOT NULL,
  `transaction_id` varchar(255) NOT NULL,
  `criteria_id` int(11) NOT NULL,
  `revenue_item` int(11) NOT NULL,
  `date_created` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `transaction_items`
--

INSERT INTO `transaction_items` (`item_id`, `transaction_id`, `criteria_id`, `revenue_item`, `date_created`) VALUES
(1, 'TRAX49773756', 2, 500000, '2022-12-13 20:43:03'),
(2, 'TRAX49773756', 3, 1000000, '2022-12-13 20:43:03'),
(3, 'TRAX78567088', 2, 500000, '2022-12-16 19:46:31'),
(4, 'TRAX78567088', 3, 1000000, '2022-12-16 19:46:31'),
(5, 'TRAX74160559', 5, 500000, '2022-12-22 10:50:24'),
(6, 'TRAX74160559', 5, 1000000, '2022-12-22 10:50:24'),
(7, 'TRAX74160559', 1, 750000, '2022-12-22 10:50:24'),
(8, 'TRAX62331089', 3, 850000, '2022-12-22 11:17:42'),
(9, 'TRAX62331089', 4, 250000, '2022-12-22 11:17:42'),
(10, 'TRAX62331089', 5, 750000, '2022-12-22 11:17:42'),
(23, 'TRAX36674280', 3, 850000, '2022-12-23 11:53:37'),
(24, 'TRAX36674280', 4, 250000, '2022-12-23 11:53:37'),
(25, 'TRAX36674280', 5, 750000, '2022-12-23 11:53:37'),
(27, 'TRAX40818546', 2, 5000000, '2022-12-23 15:44:29'),
(28, 'TRAX67201794', 2, 3000000, '2022-12-23 15:47:35'),
(29, 'TRAX67201794', 3, 750000, '2022-12-23 15:47:35'),
(30, 'TRAX67201794', 3, 850000, '2022-12-23 15:47:35'),
(31, 'TRAX67201794', 1, 1000000, '2022-12-23 15:47:35'),
(32, 'TRAX45843184', 2, 3000000, '2022-12-23 16:49:22'),
(33, 'TRAX45843184', 1, 500000, '2022-12-23 16:49:22'),
(34, 'TRAX37672880', 2, 3000000, '2022-12-23 16:50:43'),
(35, 'TRAX37672880', 1, 500000, '2022-12-23 16:50:43');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `coupon`
--
ALTER TABLE `coupon`
  ADD PRIMARY KEY (`coupon_id`);

--
-- Indeks untuk tabel `criteria`
--
ALTER TABLE `criteria`
  ADD PRIMARY KEY (`criteria_id`);

--
-- Indeks untuk tabel `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`customer_id`);

--
-- Indeks untuk tabel `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`transaction_id`);

--
-- Indeks untuk tabel `transaction_items`
--
ALTER TABLE `transaction_items`
  ADD PRIMARY KEY (`item_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `criteria`
--
ALTER TABLE `criteria`
  MODIFY `criteria_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `transaction_items`
--
ALTER TABLE `transaction_items`
  MODIFY `item_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
