-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Anamakine: db
-- Üretim Zamanı: 15 Ara 2021, 22:53:15
-- Sunucu sürümü: 10.6.5-MariaDB-1:10.6.5+maria~focal
-- PHP Sürümü: 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Veritabanı: `golang`
--
CREATE DATABASE IF NOT EXISTS `golang` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `golang`;

-- --------------------------------------------------------

--
-- Tablo için tablo yapısı `interests`
--

CREATE TABLE `interests` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Tablo döküm verisi `interests`
--

INSERT INTO `interests` (`id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'C#', 1, '2021-12-15 22:21:57', NULL, NULL),
(2, 'GO', 1, '2021-12-15 22:23:15', NULL, NULL),
(3, 'PHP', 1, '2021-12-15 22:22:23', NULL, NULL),
(4, 'Laravel', 1, '2021-12-15 22:22:29', NULL, NULL),
(5, 'Lumen', 1, '2021-12-15 22:22:33', NULL, NULL),
(6, 'Microservice Architecture', 1, '2021-12-15 22:22:47', NULL, NULL);

-- --------------------------------------------------------

--
-- Tablo için tablo yapısı `users`
--

CREATE TABLE `users` (
  `id` bigint(11) NOT NULL,
  `username` varchar(150) NOT NULL,
  `first_name` varchar(150) NOT NULL,
  `last_name` varchar(150) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password` varchar(255) NOT NULL,
  `profile` varchar(255) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Tablo döküm verisi `users`
--

INSERT INTO `users` (`id`, `username`, `first_name`, `last_name`, `email`, `password`, `profile`, `status`, `created_at`, `update_at`, `deleted_at`) VALUES
(1, 'gokhanozgur', 'gökhan', 'özgür', 'g@g.com', '$2a$04$/PlH3Yk0I.wCNClLe3UcNu01Mxfl27rSgymyG55Kae1896WIpGNTG', 'bla bla', 1, '2021-12-15 22:48:46', NULL, NULL),
(2, 'testuser', 'test', 'user', 't@u.com', '$2a$04$T7SiF0YX8cEOqqR/elqR6.GN.xJE2la7uCt5TDt3QQtldzVlvxd3W', 'bla bla', 0, '2021-12-15 22:49:05', NULL, NULL);

-- --------------------------------------------------------

--
-- Tablo için tablo yapısı `user_interests`
--

CREATE TABLE `user_interests` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `interest_id` int(20) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Tablo döküm verisi `user_interests`
--

INSERT INTO `user_interests` (`id`, `user_id`, `interest_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 1, 1, '2021-12-15 22:35:00', NULL, NULL),
(2, 1, 2, 1, '2021-12-15 22:35:09', NULL, NULL),
(3, 1, 3, 1, '2021-12-15 22:35:12', NULL, NULL),
(4, 1, 4, 1, '2021-12-15 22:35:16', NULL, NULL),
(5, 1, 5, 1, '2021-12-15 22:35:19', NULL, NULL),
(6, 1, 6, 1, '2021-12-15 22:35:23', NULL, NULL),
(8, 2, 4, 1, '2021-12-15 22:51:25', NULL, NULL);

--
-- Dökümü yapılmış tablolar için indeksler
--

--
-- Tablo için indeksler `interests`
--
ALTER TABLE `interests`
  ADD PRIMARY KEY (`id`);

--
-- Tablo için indeksler `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Tablo için indeksler `user_interests`
--
ALTER TABLE `user_interests`
  ADD PRIMARY KEY (`id`);

--
-- Dökümü yapılmış tablolar için AUTO_INCREMENT değeri
--

--
-- Tablo için AUTO_INCREMENT değeri `interests`
--
ALTER TABLE `interests`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- Tablo için AUTO_INCREMENT değeri `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Tablo için AUTO_INCREMENT değeri `user_interests`
--
ALTER TABLE `user_interests`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
