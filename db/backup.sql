-- MariaDB dump 10.19-11.2.2-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: wikipedia_api
-- ------------------------------------------------------
-- Server version	11.2.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema_migrations` (
  `version` varchar(128) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES
('20240102154804');
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wikipedia`
--

DROP TABLE IF EXISTS `wikipedia`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `wikipedia` (
  `wiki_id` varchar(35) NOT NULL,
  `wiki_word` varchar(50) NOT NULL,
  `wiki_description` text NOT NULL,
  PRIMARY KEY (`wiki_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wikipedia`
--

LOCK TABLES `wikipedia` WRITE;
/*!40000 ALTER TABLE `wikipedia` DISABLE KEYS */;
INSERT INTO `wikipedia` VALUES
('03ff46d7c1814c929799233c75f63fde','Nocturnal','Active during the night'),
('0cb227be2b664a25ba2f7b83117bcd93','Sonorous','(of a person\'s voice or other sound) Imposingly deep and full'),
('1614a3a4c3454b1eba821872e1c80cb5','Efflorescence','The process of flowering or blooming'),
('2bbc942df1b842ab97bd2f506e3e119f','Ephemeral','Lasting for a very short time transient'),
('2e32aad909394868abad04b7a5b77f12','Pineapple','A tropical fruit with a sweet, juicy flavor'),
('2f862ac6abce404bafea8675b07695ab','Serene','Calm, peaceful, and untroubled'),
('34a882229f4a4a82995cc9ded435c508','Ineffable','Too great or extreme to be expressed or described in words indescribable'),
('3ca16d2724a34b9c9771c8ac5ef07c9d','Cascade','A small waterfall or a sequence of stages in a process where each stage initiates the next'),
('3ea6a56b5641402b99f56a6aa240fab1','Quixotic','Extremely idealistic; unrealistic and impractical'),
('43cf1295a6664d78abc60b60bce5f168','Harmony','The combination of different elements in a way that produces a pleasing result'),
('44bfb6179d9c40979bd9d42d1cc65863','Ethereal','Existing in a higher spiritual plane, intangible and delicate'),
('4683853035d647848d37765dd71d8520','Serendipity','The occurrence and development of events by chance in a happy or beneficial way'),
('481183f407d747008241327d41d768c4','Obfuscate','Render obscure unclear or unintelligible'),
('48f92d0582474ebbb23ea2d0dc7432c0','Resilient','Able to withstand or recover quickly from difficult conditions'),
('4b9945b7786f4fdebdd3710eb409093d','Quagmire','A soft and boggy area of land that gives way underfoot'),
('4e5c404632d048a8b851a0523b9deeef','Effulgent','Shining brightly; radiant'),
('561cfac6bfe2412c9bbaf77bb9fe0cfc','Ethereal','Extremely delicate, light, and beautiful otherworldly'),
('57507b8ed49d4194bde6006134a5604c','Nebulous','In the form of a cloud or haze hazy'),
('596618b804e146a6934cc8546b633cbc','Mellifluous','Pleasingly smooth and musical in sound'),
('59e0dc6075534100b8a2375d77e764dd','Harmonious','Forming a harmonious or consistent whole balanced and orderly'),
('612094c1d5f34d0589555fa8c03715a4','Serendipity','The occurrence and development of events by chance in a happy or beneficial way'),
('6236739f158645ce943b7469697760db','Ejen Ali',' sebuah serial animasi Malaysia yang diproduksi oleh WAU Animation, yang menceritakan tentang seorang bocah yang tak sengaja menjadi agen MATA setelah mengaktifkan Infinity Retinal Intelligence System (IRIS), sebuah perangkat prototipe yang dibuat oleh MATA. Setelah insiden tersebut, Ali dan pamannya Bakar bekerja sama dalam misi MATA. '),
('683d9b6f78ae4fc1a531bf1d80d281ba','Nebulous','In the form of a cloud or haze hazy'),
('68a815f26c354b82a83c49246f1613fa','Inscrutable','Impossible to understand or interpret'),
('6c8d1b1c2af34a04b7d190cf0c46fcd5','Panache','Distinctive and stylish elegance'),
('71639710b5294eb68aabb0e4158b0cb3','Evanescent','Quickly fading or disappearing'),
('7363e4be129848278cbfa35db3916bfe','Incandescent','Emitting light as a result of being heated; extremely bright or radiant'),
('7d29ddba6dce472cbd639083a44b840b','Innovation','The introduction of something new a change in methods, ideas, or products'),
('80009d24e0bc4ef58a0a08afd1b36d58','Cascade','A large number of things happening in quick succession'),
('869fc125f76e4e73b8df3eead4424a91','Ineffable','Too great or extreme to be expressed or described in words indescribable'),
('8a1318b102634fe4b3ec797015092ab8','Resplendent','Shining brilliantly characterized by a glowing splendor'),
('8d22627ae52c4e30aa2627b46a87abe4','Adventure','An exciting or remarkable experience'),
('8d99d9779de14f0a96d1bbe7f12fe297','Zenith','The time at which something is most powerful or successful'),
('97b0d3d2dbce464383c689d2fa77a748','Clandestine','Kept secret or done secretly often for illicit purposes'),
('9d041d9042ff48d599860d56f294f609','Farda Ayu','Farda Ayu Nurfatika'),
('a3b00aaab7764625a844690bf6c22eeb','Idyllic','Extremely happy, peaceful, or picturesque especially in a simple or rustic way'),
('a496789847f34054ae5257f4461fbe37','Quagmire','A soft and boggy area of land that gives way underfoot'),
('a62dd212124f45e1b92c0e7799bfda7b','Panacea','A solution or remedy for all difficulties or diseases'),
('a77d9875d0964fafb9fe5ad8ad6ea228','Mesmerize','To captivate and entrance someone with a charm or spell'),
('a7e8f18698fa4169a492ba5ea2edd551','Jubilation','A feeling of great joy and triumph'),
('ac230bc5b6524d4596de076c5b6613d0','Mellifluous','Pleasingly smooth and musical in sound'),
('aea85839ca09411682ea29154361d62e','Ethereal','Extremely delicate light and beautiful; otherworldly'),
('aebb6bacfb384b2aa8f7851570bc6e56','Quandary','A state of uncertainty or perplexity especially as requiring a choice between equally unfavorable options'),
('b15f3ebf681741d1926869e86a92ff64','Labyrinth','A complex structure with winding passages and blind alleys'),
('b583d18ebd994136889f24b786b11cd9','Whimsic','An impulsive or fanciful notion'),
('b844a85109cf4b149ae2febeb85e46a8','Obfuscate','Render obscure unclear or unintelligible'),
('bd5f7c17a59b4582a588a9958a96fe6a','Serendipity','The occurrence and development of events by chance in a happy or beneficial way'),
('bed17af3fc724b93a7c1d9002635de1c','Tranquil','A state of peacefulness and serenity'),
('bfca256e25f249ab806c00b0470fd77c','Mellifluous','Pleasingly smooth and musical in sound'),
('c9f36b3837bc4d19807fb1cd3ff2f57a','Luminous','Emitting or reflecting light; full of light; bright or shining especially in the dark'),
('cc633116ad8843ccb1ab1d551b0273d4','Serenity','The state of being calm, peaceful, and untroubled'),
('cd10eb5a031647da821b918458e284d2','Zest','Enthusiasm and eagerness to do something'),
('d58468cdfcfe4487bc960072000786a5','Whimsical','Playfully quaint or fanciful especially in an appealing and amusing way'),
('d83c93c6273c4a00b13d668fa60d14fe','Labyrinthine','Complex and intricate in structure or appearance'),
('d87c8e31bf134e60aab7e2e9f64bac6a','Effervescent','Possessing a lively energy and enthusiasm'),
('d8d3214e4c4e490baff75ad1172b2c22','Palimpsest','A manuscript or piece of writing material on which later writing has been superimposed on effaced earlier writing'),
('da18480a2aa94f86bc895f4bbf7a371b','Luminous','Emitting or reflecting light; full of light; bright or shining especially in the dark'),
('de25b5754ac54232ace6e044073b9338','Resonate','Produce or be filled with a deep full reverberating sound'),
('deb24966438c4ef0935c9f3b6036b2c5','Ephemeral','Lasting for a very short time transient'),
('e2525807020141fc8d7eb8f6d8ce6839','Ubiquitous','Present appearing or found everywhere'),
('e67df82645cb4ac8b18c5ad94b447f95','Ineffable','Too great or extreme to be expressed or described in words indescribable'),
('ec64a58f91764709a376c51eb4a9807c','Happiness','A state of contentment satisfaction and joy'),
('ec8d559521144571926e695c942d41d8','Eccentric','Deviating from the norm or exhibiting unusual behavior'),
('eeaf22b8957f4175a3d3ff3dea3a5014','Zenith','The time at which something is most powerful or successful'),
('f0c16ffd14fa4abfa49c65f2b557c619','Serendipity','The occurrence and development of events by chance in a happy or beneficial way'),
('f12472b01dd64beb9bfcfbea6d82beae','Quell','Suppress (a feeling especially an unpleasant one)'),
('f203fb73b8914d9b89fcba1dbb7a9e2f','Bring Me The Horizon','grup musik rock Inggris yang dibentuk di Sheffield pada tahun 2004. Saat ini grup ini digawangi vokalis Oliver Sykes, gitaris Lee Malia, bassis Matt Kean, dan drummer Matt Nicholls. Saat ini mereka berada di bawah kontrak dengan RCA Records secara global dan Columbia Records secara eksklusif di Amerika Serikat. '),
('f5cea6535a53433db1be6067dd80e275','Mellifluous','Pleasingly smooth and musical in sound');
/*!40000 ALTER TABLE `wikipedia` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-04 21:04:50
